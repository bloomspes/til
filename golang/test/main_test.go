package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Test functions",
	func() {
		Context("sample", func() {
			It("one should return 1", func() {
				Expect(1).To(Equal(one()))
			})
			It("name should return name", func() {
				Expect("name").To(Equal(name()))
			})
		})
	})

func name() interface{} {
	return "name"
}

func one() interface{} {
	return 1
}
