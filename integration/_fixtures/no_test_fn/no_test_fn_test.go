package no_test_fn_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/ginkgo/integration/_fixtures/no_test_fn"
	. "github.com/hackrish007/gomega"
)

var _ = Describe("NoTestFn", func() {
	It("should proxy strings", func() {
		Ω(StringIdentity("foo")).Should(Equal("foo"))
	})
})
