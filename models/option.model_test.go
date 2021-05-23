package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Option model test", func() {

	var option models.Option

	BeforeEach(func() {
		option = CreateOption()
	})
	AfterEach(func() {
		option.Delete()
	})

	Context("Basic crud testing", func() {
		var obtainedOption models.Option
		err := obtainedOption.SetID(option.GetID())
		Expect(err).To(BeNil())
		err = obtainedOption.FetchById()
		Expect(err).To(BeNil())
		Expect(obtainedOption.Name).To(Equal(option.Name))
	})
})
