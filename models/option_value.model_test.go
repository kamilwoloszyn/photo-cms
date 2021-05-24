package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptionValue model", func() {
	var option models.Option
	var optionValue models.OptionValue

	BeforeEach(func() {
		optionValue = CreateOptionValue()
		option = CreateOption(&[]models.OptionValue{
			optionValue,
		})

	})

	AfterEach(func() {
		option.Delete()
		optionValue.Delete()
	})

	Describe("Basic crud testing", func() {
		It("Should be in db", func() {
			var obtainedOptionValue models.OptionValue
			err := obtainedOptionValue.SetID(option.GetID())
			Expect(err).To(BeNil())
			err = obtainedOptionValue.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedOptionValue.Value).To(Equal(optionValue.Value))
		})
	})

	Describe("Relationship test", func() {

	})
})
