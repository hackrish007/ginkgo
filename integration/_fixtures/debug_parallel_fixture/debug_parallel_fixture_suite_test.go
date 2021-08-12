package debug_parallel_fixture_test

import (
	"testing"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

func TestDebugParallelFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DebugParallelFixture Suite")
}
