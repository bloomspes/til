package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSimpleArraySum(ar []int32) int32 {
 return 6
}

var _ = Describe("Test functions",
	func() {
		Context("TestSimpleArraySum", func() {
			It("should be return array sum", func() {
				Expect(TestSimpleArraySum([1, 2, 3]))
			})
		})

	},
)
