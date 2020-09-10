/*
Ginkgo accepts a number of configuration options.

These are documented [here](http://onsi.github.io/ginkgo/#the-ginkgo-cli)

You can also learn more via

	ginkgo help

or (I kid you not):

	go test -asdf
*/
package config

import (
	"flag"
	"net/http"
	"time"

	"github.com/onsi/ginkgo/types"
)

const VERSION = "2.0.0-alpha"

type GinkgoConfigType struct {
	RandomSeed         int64
	RandomizeAllSpecs  bool
	RegexScansFilePath bool
	FocusStrings       []string
	SkipStrings        []string
	FailOnPending      bool
	FailFast           bool
	FlakeAttempts      int
	EmitSpecProgress   bool
	DryRun             bool
	DebugParallel      bool

	ParallelNode  int
	ParallelTotal int
	ParallelHost  string
}

type DefaultReporterConfigType struct {
	NoColor           bool
	SlowSpecThreshold float64
	Succinct          bool
	Verbose           bool
	FullTrace         bool
	ReportPassed      bool
	JUnitReportFile   string
}

type deprecatedConfigsType struct {
	NoisySkippings bool
	NoisyPendings  bool
}

func NewDefaultGinkgoConfig() GinkgoConfigType {
	return GinkgoConfigType{
		RandomSeed:    time.Now().Unix(),
		ParallelNode:  1,
		ParallelTotal: 1,
	}
}

func NewDefaultReporterConfig() DefaultReporterConfigType {
	return DefaultReporterConfigType{
		SlowSpecThreshold: 5.0,
	}
}

var GinkgoConfig = NewDefaultGinkgoConfig()
var DefaultReporterConfig = NewDefaultReporterConfig()

var GinkgoConfigFlags = GinkgoFlags{
	{KeyPath: "G.RandomSeed", Name: "seed", SectionKey: "order", UsageDefaultValue: "randomly generated by Ginkgo",
		Usage: "The seed used to randomize the spec suite."},
	{KeyPath: "G.RandomizeAllSpecs", Name: "randomize-all", SectionKey: "order", DeprecatedName: "randomizeAllSpecs", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "If set, ginkgo will randomize all specs together.  By default, ginkgo only randomizes the top level Describe, Context and When containers."},

	{KeyPath: "G.FailOnPending", Name: "fail-on-pending", SectionKey: "failure", DeprecatedName: "failOnPending", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "If set, ginkgo will mark the test suite as failed if any specs are pending."},
	{KeyPath: "G.FailFast", Name: "fail-fast", SectionKey: "failure", DeprecatedName: "failFast", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "If set, ginkgo will stop running a test suite after a failure occurs."},
	{KeyPath: "G.FlakeAttempts", Name: "flake-attempts", SectionKey: "failure", UsageDefaultValue: "0 - failed tests are not retried", DeprecatedName: "flakeAttempts", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "Make up to this many attempts to run each spec. If any of the attempts succeed, the suite will not be failed."},

	{KeyPath: "G.DryRun", Name: "dry-run", SectionKey: "debug", DeprecatedName: "dryRun", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "If set, ginkgo will walk the test hierarchy without actually running anything.  Best paired with -v."},
	{KeyPath: "G.EmitSpecProgress", Name: "progress", SectionKey: "debug",
		Usage: "If set, ginkgo will emit progress information as each spec runs to the GinkgoWriter."},

	{KeyPath: "G.FocusStrings", Name: "focus", SectionKey: "filter",
		Usage: "If set, ginkgo will only run specs that match this regular expression. Can be specified multiple times, values are ORed."},
	{KeyPath: "G.SkipStrings", Name: "skip", SectionKey: "filter",
		Usage: "If set, ginkgo will only run specs that do not match this regular expression. Can be specified multiple times, values are ORed."},
	{KeyPath: "G.RegexScansFilePath", Name: "regex-scans-filepath", SectionKey: "filter", DeprecatedName: "regexScansFilePath", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "If set, ginkgo regex matching also will look at the file path (code location)."},
	{KeyPath: "G.DebugParallel", Name: "debug-parallel", SectionKey: "debug", DeprecatedName: "debug",
		Usage: "If set, ginkgo will emit node output to files when running in parallel."},
}

var GinkgoParallelConfigFlags = GinkgoFlags{
	{KeyPath: "G.ParallelNode", Name: "parallel.node", SectionKey: "low-level-parallel", UsageDefaultValue: "1",
		Usage: "This worker node's (one-indexed) node number.  For running specs in parallel."},
	{KeyPath: "G.ParallelTotal", Name: "parallel.total", SectionKey: "low-level-parallel", UsageDefaultValue: "1",
		Usage: "The total number of worker nodes.  For running specs in parallel."},
	{KeyPath: "G.ParallelHost", Name: "parallel.host", SectionKey: "low-level-parallel", UsageDefaultValue: "set by Ginkgo CLI",
		Usage: "The address for the server that will synchronize the running nodes."},
}

var ReporterConfigFlags = GinkgoFlags{
	{KeyPath: "R.NoColor", Name: "no-color", SectionKey: "output", DeprecatedName: "noColor", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "If set, suppress color output in default reporter."},
	{KeyPath: "R.SlowSpecThreshold", Name: "slow-spec-threshold", SectionKey: "output", UsageArgument: "float, in seconds", UsageDefaultValue: "5.0", DeprecatedName: "slowSpecThreshold", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "Specs that take longer to run than this threshold are flagged as slow by the default reporter."},
	{KeyPath: "D.NoisyPendings", DeprecatedName: "noisyPendings", DeprecatedDocLink: "removed--noisypendings-and--noisyskippings"},
	{KeyPath: "D.NoisySkippings", DeprecatedName: "noisySkippings", DeprecatedDocLink: "removed--noisypendings-and--noisyskippings"},
	{KeyPath: "R.Verbose", Name: "v", SectionKey: "output",
		Usage: "If set, default reporter print out all specs as they begin."},
	{KeyPath: "R.Succinct", Name: "succinct", SectionKey: "output",
		Usage: "If set, default reporter prints out a very succinct report"},
	{KeyPath: "R.FullTrace", Name: "trace", SectionKey: "output",
		Usage: "If set, default reporter prints out the full stack trace when a failure occurs"},
	{KeyPath: "R.ReportPassed", Name: "report-passed", SectionKey: "output", DeprecatedName: "reportPassed", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "If set, default reporter prints out captured output of passed tests."},
	{KeyPath: "R.JUnitReportFile", Name: "junit-report", SectionKey: "output", DeprecatedName: "reportFile", DeprecatedDocLink: "changed-command-line-flags",
		Usage: "If set, Ginkgo will generate a junit test report at the specified destination."},
}

var FlagSections = GinkgoFlagSections{
	{Key: "multiple-suites", Style: "{{dark-green}}", Heading: "Running Multiple Test Suites"},
	{Key: "order", Style: "{{green}}", Heading: "Controlling Test Order"},
	{Key: "parallel", Style: "{{yellow}}", Heading: "Controlling Test Parallelism"},
	{Key: "low-level-parallel", Style: "{{yellow}}", Heading: "Controlling Test Parallelism",
		Description: "These are set by the Ginkgo CLI, {{red}}{{bold}}do not set them manually{{/}} via go test.\nUse ginkgo -p or ginkgo -nodes=N instead."},
	{Key: "filter", Style: "{{cyan}}", Heading: "Filtering Tests"},
	{Key: "failure", Style: "{{red}}", Heading: "Failure Handling"},
	{Key: "output", Style: "{{magenta}}", Heading: "Controlling Output Formatting"},
	{Key: "code-and-coverage-analysis", Style: "{{orange}}", Heading: "Code and Coverage Analysis"},
	{Key: "performance-analysis", Style: "{{coral}}", Heading: "Performance Analysis"},
	{Key: "debug", Style: "{{blue}}", Heading: "Debugging Tests"},
	{Key: "watch", Style: "{{light-yellow}}", Heading: "Controlling Ginkgo Watch"},
	{Key: "misc", Style: "{{light-gray}}", Heading: "Miscellaneous"},
	{Key: "go-build", Style: "{{light-gray}}", Heading: "Go Build Flags", Succinct: true,
		Description: "These flags are inherited from go build.  Run {{bold}}ginkgo help build{{/}} for more detailed flag documentation."},
}

// The FlagSet for the user's test sute itself.
func BuildTestSuiteFlagSet() (GinkgoFlagSet, error) {
	flags := GinkgoConfigFlags.CopyAppend(GinkgoParallelConfigFlags...).CopyAppend(ReporterConfigFlags...)
	flags = flags.WithPrefix("ginkgo")
	bindings := map[string]interface{}{
		"G": &GinkgoConfig,
		"R": &DefaultReporterConfig,
		"D": &deprecatedConfigsType{},
	}
	extraGoFlagsSection := GinkgoFlagSection{Style: "{{gray}}", Heading: "Go test flags"}

	return NewAttachedGinkgoFlagSet(flag.CommandLine, flags, bindings, FlagSections, extraGoFlagsSection)
}

func VetConfig(flagSet GinkgoFlagSet, config GinkgoConfigType, reporterConfig DefaultReporterConfigType) []error {
	errors := []error{}

	if flagSet.WasSet("count") || flagSet.WasSet("test.count") {
		errors = append(errors, types.GinkgoErrors.InvalidGoFlagCount())
	}

	if flagSet.WasSet("parallel") || flagSet.WasSet("test.parallel") {
		errors = append(errors, types.GinkgoErrors.InvalidGoFlagParallel())
	}

	if config.ParallelTotal < 1 {
		errors = append(errors, types.GinkgoErrors.InvalidParallelTotalConfiguration())
	}

	if config.ParallelNode > config.ParallelTotal || config.ParallelNode < 1 {
		errors = append(errors, types.GinkgoErrors.InvalidParallelNodeConfiguration())
	}

	if config.ParallelTotal > 1 {
		if config.ParallelHost == "" {
			errors = append(errors, types.GinkgoErrors.MissingParallelHostConfiguration())
		} else {
			resp, err := http.Get(config.ParallelHost + "/up")
			if err != nil || resp.StatusCode != http.StatusOK {
				errors = append(errors, types.GinkgoErrors.UnreachableParallelHost(config.ParallelHost))
			}
			if err != nil {
				resp.Body.Close()
			}
		}
	}

	if config.DryRun && config.ParallelTotal > 1 {
		errors = append(errors, types.GinkgoErrors.DryRunInParallelConfiguration())
	}

	if reporterConfig.Succinct && reporterConfig.Verbose {
		errors = append(errors, types.GinkgoErrors.ConflictingVerboseSuccinctConfiguration())
	}

	return errors
}
