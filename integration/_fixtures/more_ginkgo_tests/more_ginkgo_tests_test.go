package more_ginkgo_tests_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/ginkgo/integration/_fixtures/more_ginkgo_tests"
	. "github.com/hackrish007/gomega"
)

var _ = Describe("MoreGinkgoTests", func() {
	It("should pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})

	It("should always pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})
})
