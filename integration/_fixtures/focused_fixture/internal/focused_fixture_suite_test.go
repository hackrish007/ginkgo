package focused_fixture_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestFocused_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Focused_fixture Suite")
}
