package A_test

import (
	. "$ROOT_PATH$/A"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

var _ = Describe("A", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
