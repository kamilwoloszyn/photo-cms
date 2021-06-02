package checkers_test

import (
	"github.com/google/uuid"
	"github.com/kamilwoloszyn/photo-cms/pkg/checkers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers test", func() {

	var (
		e              uuid.UUID
		emptyStringId  checkers.IdChecker
		emptyGenericId checkers.IdChecker
		goodId         checkers.IdChecker
		goodIdStr      checkers.IdChecker
		badIdStr       checkers.IdChecker
	)

	BeforeEach(func() {
		emptyStringId = checkers.UuidString("")
		emptyGenericId = checkers.UuidGeneric(e)
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
		Context("UUUID generic empty", func() {
			It("Should be empty", func() {
				result := emptyGenericId.IsEmpty()
				Expect(result).To(BeTrue())
			})
			It("Should be not valid", func() {
				result := emptyGenericId.IsValid()
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
				result := emptyStringId.IsEmpty()
				Expect(result).To(BeTrue())
			})
			It("Should not be valid", func() {
				result := emptyStringId.IsValid()
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
