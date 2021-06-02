package checkers_test

import (
	"github.com/google/uuid"
	"github.com/kamilwoloszyn/photo-cms/pkg/checkers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers test", func() {

	var (
		emptyId   checkers.IdChecker
		goodId    checkers.IdChecker
		goodIdStr checkers.IdChecker
		badIdStr  checkers.IdChecker
	)

	BeforeEach(func() {
		emptyId = checkers.UuidString("")
		goodId = checkers.UuidGeneric(uuid.New())
		goodIdStr = checkers.UuidString(uuid.New().String())
		badIdStr = checkers.UuidString("123553-dfwefwf-r23232")
	})

	Describe("UUID test", func() {
		Context("Correct UUID generic ", func() {
			It("Should be valid", func() {
				result := goodId.IsValid()
				Expect(result).To(BeTrue())
			})

			It("Should not be empty", func() {
				result := goodId.IsEmpty()
				Expect(result).To(BeFalse())
			})
		})
		Context("Correct Uuid string", func() {
			It("Should be valid", func() {
				result := goodIdStr.IsValid()
				Expect(result).To(BeTrue())
			})
			It("Should not be empty", func() {
				result := goodIdStr.IsEmpty()
				Expect(result).To(BeFalse())
			})
		})
		Context("Empty id string", func() {
			It("Should be empty", func() {
				result := emptyId.IsEmpty()
				Expect(result).To(BeTrue())
			})
			It("Should not be valid", func() {
				result := emptyId.IsValid()
				Expect(result).To(BeFalse())
			})

		})

		Context("Wrong uuid", func() {
			It("Should not be valid", func() {
				result := badIdStr.IsValid()
				Expect(result).To(BeFalse())
			})
			It("Should not be empty", func() {
				result := badIdStr.IsEmpty()
				Expect(result).To(BeFalse())
			})
		})
	})
})
