package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product model", func() {
	var category models.Category
	var image models.Image
	var option models.Option
	var optionValue models.OptionValue
	var productOption models.ProductOption
	var customer models.Customer
	var product models.Product

	BeforeEach(func() {
		category = CreateCategory()
		image = CreateImage()
		optionValue = CreateOptionValue()
		option = CreateOption(&[]models.OptionValue{
			optionValue,
		})
		customer = CreateCustomer()
		product = CreateProductWithoutOrder(&[]models.Category{
			category,
		}, &[]models.Image{
			image,
		}, &[]models.Customer{
			customer,
		})
		productOption = CreateProductOption(&product, &[]models.OptionValue{
			optionValue,
		})
	})
	AfterEach(func() {
		productOption.Delete()
		product.Delete()
		customer.Delete()
		optionValue.Delete()
		category.Delete()
		image.Delete()
		option.Delete()
	})

	Describe("Basic crud testing", func() {
		Context("Product without order", func() {
			It("Should exist in db", func() {
				var productObtained models.Product
				err := productObtained.SetID(product.GetID())
				Expect(err).To(BeNil())
				err = productObtained.FetchByID()
				Expect(err).To(BeNil())
				Expect(productObtained.ProductName).To(Equal(product.ProductName))
			})
		})

	})

})
