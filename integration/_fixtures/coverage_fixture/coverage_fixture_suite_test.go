package coverage_fixture_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestCoverageFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CoverageFixture Suite")
}
