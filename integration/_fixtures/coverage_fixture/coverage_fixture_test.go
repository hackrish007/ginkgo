package coverage_fixture_test

import (
	. "github.com/hackrish007/ginkgo/integration/_fixtures/coverage_fixture"
	. "github.com/hackrish007/ginkgo/integration/_fixtures/coverage_fixture/external_coverage_fixture"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

var _ = Describe("CoverageFixture", func() {
	It("should test A", func() {
		Ω(A()).Should(Equal("A"))
	})

	It("should test B", func() {
		Ω(B()).Should(Equal("B"))
	})

	It("should test C", func() {
		Ω(C()).Should(Equal("C"))
	})

	It("should test D", func() {
		Ω(D()).Should(Equal("D"))
	})

	It("should test external package", func() {
		Ω(Tested()).Should(Equal("tested"))
	})
})
