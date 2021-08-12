package outline_test

import (
	"testing"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

func TestOutline(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Outline Suite")
}
