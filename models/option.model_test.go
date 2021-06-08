package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
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
		BeforeEach(func() {
			obtainedOption = models.Option{}
		})
		Context("Create or update methods", func() {
			It("Should create a new option model", func() {
				err := obtainedOption.SetID(option.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedOption.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedOption.Name).To(Equal(option.Name))
			})

			It("Should update an existing option", func() {
				err := obtainedOption.SetID(option.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				option.Name = "changed"
				err = option.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedOption.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedOption.Name).To(Equal(option.Name))
			})
		})

		Context("Delete method", func() {
			It("Should delete an option", func() {
				err := obtainedOption.SetID(option.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = option.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedOption.FetchById()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
				Expect(obtainedOption.Name).To(BeEmpty())
			})
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

		Context("Two values", func() {
			It("Should be in db", func() {
				err := option.GetOptionValues()
				Expect(err).To(BeNil())
				Expect(len(option.OptionValue)).To(Equal(2))
			})

		})
	})
})
