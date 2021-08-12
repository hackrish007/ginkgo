package B_test

import (
	. "$ROOT_PATH$/B"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

var _ = Describe("B", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
