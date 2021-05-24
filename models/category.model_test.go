package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Category Model", func() {
	var category models.Category
	BeforeEach(func() {
		category = CreateCategory()
	})
	AfterEach(func() {
		category.Delete()
	})
	Describe("Crud operations", func() {
		It("Should be stored into db", func() {
			var obtainedCategory models.Category
			err := obtainedCategory.SetID(category.GetID())
			Expect(err).To(BeNil())
			err = obtainedCategory.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedCategory.CategoryName).To(Equal(category.CategoryName))
		})
	})
})
