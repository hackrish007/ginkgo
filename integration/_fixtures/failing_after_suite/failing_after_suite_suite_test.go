package failing_before_suite_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestFailingAfterSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FailingAfterSuite Suite")
}

var _ = BeforeSuite(func() {
	println("BEFORE SUITE")
})

var _ = AfterSuite(func() {
	println("AFTER SUITE")
	panic("BAM!")
})
