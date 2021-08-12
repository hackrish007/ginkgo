package testingtproxy_test

import (
	"testing"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

func TestTestingtproxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testingtproxy Suite")
}
