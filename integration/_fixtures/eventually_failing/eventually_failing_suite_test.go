package eventually_failing_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestEventuallyFailing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EventuallyFailing Suite")
}
