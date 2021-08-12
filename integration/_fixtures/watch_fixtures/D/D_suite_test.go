package D_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "D Suite")
}
