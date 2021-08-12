package suite_command_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestSuiteCommand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite Command Suite")
}
