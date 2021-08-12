package fail_fixture_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestFail_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Skip_fixture Suite")
}
