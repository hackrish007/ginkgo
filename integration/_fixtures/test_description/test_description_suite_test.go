package test_description_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestTestDescription(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestDescription Suite")
}
