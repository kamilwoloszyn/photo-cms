package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Option model test", func() {

	var optionVal models.OptionValue
	var option models.Option

	BeforeEach(func() {
		option = CreateOption()
		optionVal = CreateOptionValue(&option)
	})
	AfterEach(func() {
		optionVal.Delete()
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
})
