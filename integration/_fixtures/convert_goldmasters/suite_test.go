package tmp_test

import (
	"testing"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

func TestConvertFixtures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConvertFixtures Suite")
}
