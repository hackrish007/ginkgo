package progress_fixture_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestProgressFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProgressFixture Suite")
}
