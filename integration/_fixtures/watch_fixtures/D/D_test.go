package D_test

import (
	. "$ROOT_PATH$/D"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

var _ = Describe("D", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
