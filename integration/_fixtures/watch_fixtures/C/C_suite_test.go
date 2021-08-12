package C_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "C Suite")
}
