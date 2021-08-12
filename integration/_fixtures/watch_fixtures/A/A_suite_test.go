package A_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Suite")
}
