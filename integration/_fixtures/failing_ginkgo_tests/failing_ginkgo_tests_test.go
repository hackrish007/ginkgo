package failing_ginkgo_tests_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/ginkgo/integration/_fixtures/failing_ginkgo_tests"
	. "github.com/hackrish007/gomega"
)

var _ = Describe("FailingGinkgoTests", func() {
	It("should fail", func() {
		Ω(AlwaysFalse()).Should(BeTrue())
	})

	It("should pass", func() {
		Ω(AlwaysFalse()).Should(BeFalse())
	})
})
