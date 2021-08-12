package tmp

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestTmp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tmp Suite")
}
