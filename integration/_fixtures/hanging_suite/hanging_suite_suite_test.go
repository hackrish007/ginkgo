package hanging_suite_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestHangingSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HangingSuite Suite")
}
