package C_test

import (
	. "$ROOT_PATH$/C"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

var _ = Describe("C", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
