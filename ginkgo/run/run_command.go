package run

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/onsi/ginkgo/formatter"
	"github.com/onsi/ginkgo/ginkgo/command"
	"github.com/onsi/ginkgo/ginkgo/internal"
	"github.com/onsi/ginkgo/ginkgo/interrupthandler"
	"github.com/onsi/ginkgo/types"
)

func BuildRunCommand() command.Command {
	var suiteConfig = types.NewDefaultSuiteConfig()
	var reporterConfig = types.NewDefaultReporterConfig()
	var cliConfig = types.NewDefaultCLIConfig()
	var goFlagsConfig = types.NewDefaultGoFlagsConfig()

	flags, err := types.BuildRunCommandFlagSet(&suiteConfig, &reporterConfig, &cliConfig, &goFlagsConfig)
	if err != nil {
		panic(err)
	}

	interruptHandler := interrupthandler.NewInterruptHandler()

	return command.Command{
		Name:          "run",
		Flags:         flags,
		Usage:         "ginkgo run <FLAGS> <PACKAGES> -- <PASS-THROUGHS>",
		ShortDoc:      "Run the tests in the passed in <PACKAGES> (or the package in the current directory if left blank)",
		Documentation: "Any arguments after -- will be passed to the test.",
		DocLink:       "running-tests",
		Command: func(args []string, additionalArgs []string) {
			var errors []error
			cliConfig, goFlagsConfig, errors = types.VetAndInitializeCLIAndGoConfig(cliConfig, goFlagsConfig)
			command.AbortIfErrors("Ginkgo detected configuration issues:", errors)

			runner := &SpecRunner{
				cliConfig:      cliConfig,
				goFlagsConfig:  goFlagsConfig,
				suiteConfig:    suiteConfig,
				reporterConfig: reporterConfig,
				flags:          flags,

				interruptHandler: interruptHandler,
			}

			runner.RunSpecs(args, additionalArgs)
		},
	}
}

type SpecRunner struct {
	suiteConfig    types.SuiteConfig
	reporterConfig types.ReporterConfig
	cliConfig      types.CLIConfig
	goFlagsConfig  types.GoFlagsConfig
	flags          types.GinkgoFlagSet

	interruptHandler *interrupthandler.InterruptHandler
}

//TODO - THIS IS INTENTIONALLY SLOWER THAN GINKGO V1 FOR NOW
//COMPILATION AND RUNNING AND SERIALIZED.  IN TIME WE'LL RUN EXPERIMENTS TO FIND THE OPTIMAL WAY TO
//IMPROVE PERFORMANCE
func (r *SpecRunner) RunSpecs(args []string, additionalArgs []string) {
	suites := internal.FindSuites(args, r.cliConfig, true)
	skippedSuites := suites.WithState(internal.TestSuiteStateSkippedByFilter)
	suites = suites.WithoutState(internal.TestSuiteStateSkippedByFilter)

	if len(skippedSuites) > 0 {
		fmt.Println("Will skip:")
		for _, skippedSuite := range skippedSuites {
			fmt.Println("  " + skippedSuite.Path)
		}
	}

	if len(skippedSuites) > 0 && len(suites) == 0 {
		command.Abort(command.AbortDetails{
			ExitCode: 0,
			Error:    fmt.Errorf("All tests skipped! Exiting..."),
		})
	}

	if len(suites) == 0 {
		command.AbortWith("Found no test suites")
	}

	if len(suites) > 1 && !r.flags.WasSet("succinct") && !r.reporterConfig.Verbose {
		r.reporterConfig.Succinct = true
	}

	t := time.Now()
	var endTime time.Time
	if r.suiteConfig.Timeout > 0 {
		endTime = t.Add(r.suiteConfig.Timeout)
	}

	iteration := 0

OUTER_LOOP:
	for {
		if !r.flags.WasSet("seed") {
			r.suiteConfig.RandomSeed = time.Now().Unix()
		}
		if r.cliConfig.RandomizeSuites && len(suites) > 1 {
			suites = suites.ShuffledCopy(r.suiteConfig.RandomSeed)
		}

	SUITE_LOOP:
		for suiteIdx := range suites {
			if r.interruptHandler.WasInterrupted() {
				break OUTER_LOOP
			}

			if !endTime.IsZero() && time.Now().After(endTime) {
				suites[suiteIdx].State = internal.TestSuiteStateFailedDueToTimeout
				continue SUITE_LOOP
			}

			if suites.CountWithState(internal.TestSuiteStateFailureStates...) > 0 && !r.cliConfig.KeepGoing {
				suites[suiteIdx].State = internal.TestSuiteStateSkippedDueToPriorFailures
				continue SUITE_LOOP
			}

			suites[suiteIdx] = internal.CompileSuite(suites[suiteIdx], r.goFlagsConfig)
			if suites[suiteIdx].State.Is(internal.TestSuiteStateFailedToCompile) {
				fmt.Println(suites[suiteIdx].CompilationError.Error())
				continue SUITE_LOOP
			}

			if r.interruptHandler.WasInterrupted() {
				break OUTER_LOOP
			}

			if !endTime.IsZero() {
				r.suiteConfig.Timeout = endTime.Sub(time.Now())
				if r.suiteConfig.Timeout <= 0 {
					suites[suiteIdx].State = internal.TestSuiteStateFailedDueToTimeout
					continue SUITE_LOOP
				}
			}

			suites[suiteIdx] = internal.RunCompiledSuite(suites[suiteIdx], r.suiteConfig, r.reporterConfig, r.cliConfig, r.goFlagsConfig, additionalArgs)
		}

		if suites.CountWithState(internal.TestSuiteStateFailureStates...) > 0 {
			if iteration > 0 {
				fmt.Printf("\nTests failed on attempt #%d\n\n", iteration+1)
			}
			break OUTER_LOOP
		}

		if r.cliConfig.UntilItFails {
			fmt.Printf("\nAll tests passed...\nWill keep running them until they fail.\nThis was attempt #%d\n%s\n", iteration+1, orcMessage(iteration+1))
		} else if r.cliConfig.Repeat > 0 && iteration < r.cliConfig.Repeat {
			fmt.Printf("\nAll tests passed...\nThis was attempt %d of %d.", iteration+1, r.cliConfig.Repeat+1)
		} else {
			break OUTER_LOOP
		}
		iteration += 1
	}

	internal.Cleanup(r.goFlagsConfig, suites...)

	messages, err := internal.FinalizeProfilesAndReportsForSuites(suites, r.cliConfig, r.suiteConfig, r.reporterConfig, r.goFlagsConfig)
	command.AbortIfError("could not finalize profiles:", err)
	for _, message := range messages {
		fmt.Println(message)
	}

	fmt.Printf("\nGinkgo ran %d %s in %s\n", len(suites), internal.PluralizedWord("suite", "suites", len(suites)), time.Since(t))

	if suites.CountWithState(internal.TestSuiteStateFailureStates...) == 0 {
		if suites.AnyHaveProgrammaticFocus() && strings.TrimSpace(os.Getenv("GINKGO_EDITOR_INTEGRATION")) == "" {
			fmt.Printf("Test Suite Passed\n")
			fmt.Printf("Detected Programmatic Focus - setting exit status to %d\n", types.GINKGO_FOCUS_EXIT_CODE)
			command.Abort(command.AbortDetails{ExitCode: types.GINKGO_FOCUS_EXIT_CODE})
		} else {
			fmt.Printf("Test Suite Passed\n")
			command.Abort(command.AbortDetails{})
		}
	} else {
		fmt.Fprintln(formatter.ColorableStdOut, "")
		if suites.CountWithState(internal.TestSuiteStateFailureStates...) > 1 {
			fmt.Fprintln(formatter.ColorableStdOut,
				internal.FailedSuitesReport(suites, formatter.NewWithNoColorBool(r.reporterConfig.NoColor)))
		}
		fmt.Printf("Test Suite Failed\n")
		command.Abort(command.AbortDetails{ExitCode: 1})
	}
}

func orcMessage(iteration int) string {
	if iteration < 10 {
		return ""
	} else if iteration < 30 {
		return []string{
			"If at first you succeed...",
			"...try, try again.",
			"Looking good!",
			"Still good...",
			"I think your tests are fine....",
			"Yep, still passing",
			"Oh boy, here I go testin' again!",
			"Even the gophers are getting bored",
			"Did you try -race?",
			"Maybe you should stop now?",
			"I'm getting tired...",
			"What if I just made you a sandwich?",
			"Hit ^C, hit ^C, please hit ^C",
			"Make it stop. Please!",
			"Come on!  Enough is enough!",
			"Dave, this conversation can serve no purpose anymore. Goodbye.",
			"Just what do you think you're doing, Dave? ",
			"I, Sisyphus",
			"Insanity: doing the same thing over and over again and expecting different results. -Einstein",
			"I guess Einstein never tried to churn butter",
		}[iteration-10] + "\n"
	} else {
		return "No, seriously... you can probably stop now.\n"
	}
}
