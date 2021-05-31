package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Option test", func() {
	var optionsValues models.OptionValue
	var option models.Option
	var product models.Product
	var image models.Image
	var category models.Category
	var customer models.Customer
	var productOption models.ProductOption
	BeforeEach(func() {
		option = CreateOption()
		optionsValues = CreateOptionValue(&option)
		customer = CreateCustomer()
		category = CreateCategory()
		image = CreateImage()
		product = CreateProductWithoutOrder(&category, &image, &customer)
		productOption = CreateProductOption(&product, &optionsValues)
	})

	AfterEach(func() {
		productOption.Delete()
		product.Delete()
		image.Delete()
		category.Delete()
		optionsValues.Delete()
		option.Delete()
	})
	Context("Basic crud testing", func() {
		It("Should be in db", func() {
			var obtainedProductOption models.ProductOption
			err := obtainedProductOption.SetID(productOption.GetID())
			Expect(err).To(BeNil())
			err = obtainedProductOption.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedProductOption.ProductId).To(Equal(product.ID))
		})
	})
})
