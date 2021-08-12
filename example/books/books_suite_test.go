package books_test

import (
	"testing"

	. "github.com/hackrish007/ginkgo"
	. "github.com/hackrish007/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}
