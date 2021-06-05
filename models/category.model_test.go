package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Category Model", func() {
	var (
		obtainedCategory models.Category
		category         models.Category
		image            models.Image
		customers        models.Customer
	)
	BeforeEach(func() {
		category = CreateCategory()
	})
	Describe("Crud operations", func() {

		Context("Create or update", func() {
			AfterEach(func() {
				category.Delete()
			})
			It("Should create a new category", func() {
				err := obtainedCategory.SetID(category.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedCategory.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedCategory.CategoryName).To(Equal(category.CategoryName))
			})

			It("Should update an existing category", func() {
				category.CategoryName = "Changed"
				err := category.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				obtainedCategory.SetID(category.GetID())
				err = obtainedCategory.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedCategory.CategoryName).To(Equal("Changed"))
			})
		})

		Context("Delete", func() {
			It("Should delete an existing category", func() {
				var testCategory models.Category
				testCategory.SetID(category.GetID())
				err := category.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = testCategory.FetchById()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
				Expect(testCategory.CategoryName).To(Equal(""))
			})
		})

	})
	Describe("Relationship test", func() {
		var product models.Product
		Context("Category with product", func() {
			BeforeEach(func() {
				image = CreateImage()
				customers = CreateCustomer()
				product = CreateProductWithoutOrder(category, image, customers)
			})
			AfterEach(func() {
				product.Delete()
				image.Delete()
				customers.Delete()
				category.Delete()
			})

			It("Should have a product", func() {
				err := category.FetchProducts()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(category.Product)).To(Equal(1))
				Expect(product.CategoryId).To(Equal(category.GetID()))
			})
		})

	})
})
