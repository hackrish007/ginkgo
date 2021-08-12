package codelocation_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestCodelocation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CodeLocation Suite")
}
