package integration_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
	"github.com/hackrish007/gomega/gbytes"
	"github.com/hackrish007/gomega/gexec"
)

var _ = Describe("TestDescription", func() {
	var pathToTest string

	BeforeEach(func() {
		pathToTest = tmpPath("test_description")
		copyIn(fixturePath("test_description"), pathToTest, false)
	})

	It("should capture and emit information about the current test", func() {
		session := startGinkgo(pathToTest, "--noColor")
		Eventually(session).Should(gexec.Exit(1))

		Ω(session).Should(gbytes.Say("TestDescription should pass:false"))
		Ω(session).Should(gbytes.Say("TestDescription should fail:true"))
	})
})
