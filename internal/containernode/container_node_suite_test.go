package containernode_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestContainernode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containernode Suite")
}
