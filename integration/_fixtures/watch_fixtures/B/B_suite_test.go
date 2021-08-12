package B_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "B Suite")
}
