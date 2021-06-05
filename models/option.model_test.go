package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Option model test", func() {

	var optionVal models.OptionValue
	var anotherOptionVal models.OptionValue
	var option models.Option

	BeforeEach(func() {
		option = CreateOption()
	})
	AfterEach(func() {
		option.Delete()

	})

	Describe("Basic crud testing", func() {
		var obtainedOption models.Option
		It("Should be stored into db", func() {
			err := obtainedOption.SetID(option.GetID())
			Expect(err).To(BeNil())
			err = obtainedOption.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedOption.Name).To(Equal(option.Name))
		})

	})
	Describe("Relationship test", func() {

		BeforeEach(func() {
			optionVal = CreateOptionValue(option)
			anotherOptionVal = models.OptionValue{
				Value:      "Pink",
				ExtraPrice: 13,
			}
			CreateCustomOptionValue(option, &anotherOptionVal)
		})

		AfterEach(func() {
			optionVal.Delete()
			anotherOptionVal.Delete()
		})

		Context("One value", func() {
			It("Should be into db", func() {
				err := option.GetOptionValues()
				Expect(err).To(BeNil())
				Expect(option.OptionValue[0].Value).To(Equal(optionVal.Value))
			})

		})

		Context("More values", func() {
			It("Should be in db", func() {
				err := option.GetOptionValues()
				Expect(err).To(BeNil())
				Expect(len(option.OptionValue)).To(Equal(4))
			})

		})
	})
})
