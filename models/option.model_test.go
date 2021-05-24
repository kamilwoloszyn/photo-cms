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
		optionVal = CreateOptionValue()
		option = CreateOption(&[]models.OptionValue{
			optionVal,
		})
	})
	AfterEach(func() {
		option.Delete()
		optionVal.Delete()
	})

	Describe("Basic crud testing", func() {
		var obtainedOption models.Option
		err := obtainedOption.SetID(option.GetID())
		Expect(err).To(BeNil())
		err = obtainedOption.FetchById()
		Expect(err).To(BeNil())
		Expect(obtainedOption.Name).To(Equal(option.Name))
	})
})
