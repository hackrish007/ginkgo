package spec_iterator_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestSpecIterator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SpecIterator Suite")
}
