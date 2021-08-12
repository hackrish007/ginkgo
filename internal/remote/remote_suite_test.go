package remote_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestRemote(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Remote Spec Forwarding Suite")
}
