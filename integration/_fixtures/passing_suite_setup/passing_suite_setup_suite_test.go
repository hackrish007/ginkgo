package passing_before_suite_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestPassingSuiteSetup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PassingSuiteSetup Suite")
}

var a string
var b string

var _ = BeforeSuite(func() {
	a = "ran before suite"
	println("BEFORE SUITE")
})

var _ = AfterSuite(func() {
	b = "ran after suite"
	println("AFTER SUITE")
})
