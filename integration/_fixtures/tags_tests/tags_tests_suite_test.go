package tags_tests_test

import (
	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"

	"testing"
)

func TestTagsTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TagsTests Suite")
}
