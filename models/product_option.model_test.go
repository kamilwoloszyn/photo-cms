package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Product Option test", func() {
	var (
		optionsValues models.OptionValue
		option        models.Option
		product       models.Product
		image         models.Image
		category      models.Category
		customer      models.Customer
		productOption models.ProductOption
		image2        models.Image
		product2      models.Product
	)
	BeforeEach(func() {
		image2 = models.Image{
			Name:     "samplee_image",
			FullPath: "/tmp/samplee_image",
			Size:     1234,
		}
		product2 = models.Product{
			UnitPrice:   130,
			ProductName: "samplee_product",
			Quantity:    4,
		}
	})

	JustBeforeEach(func() {
		option = CreateOption()
		optionsValues = CreateOptionValue(option)
		customer = CreateCustomer()
		category = CreateCategory()
		image = CreateImage()
		product = CreateProductWithoutOrder(category, image, customer)
		productOption = CreateProductOption(product, optionsValues)
		CreateCustomImage(&image2)
		CreateCustomProductWithoutOrder(category, image2, customer, &product2)
	})
	AfterEach(func() {
		productOption.Delete()
		product.Delete()
		product2.Delete()
		image.Delete()
		image2.Delete()
		category.Delete()
		optionsValues.Delete()
		option.Delete()
		customer.Delete()
	})
	Context("Basic crud testing", func() {
		var (
			obtainedProductOption models.ProductOption
		)
		BeforeEach(func() {
			obtainedProductOption = models.ProductOption{}
		})
		Context("Update or delete operations", func() {
			It("Should create a new product option", func() {
				err := obtainedProductOption.SetID(productOption.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedProductOption.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedProductOption.GetID()).To(Equal(productOption.GetID()))
				Expect(obtainedProductOption.OptionValueId).To(Equal(productOption.OptionValueId))
				Expect(obtainedProductOption.ProductId).To(Equal(productOption.ProductId))
			})

			It("Should update an existing product option", func() {
				err := obtainedProductOption.SetID(productOption.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = product2.AssignTo(&productOption)
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedProductOption.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedProductOption.ProductId).To(Equal(productOption.ProductId))
			})
		})

		Context("Delete operations", func() {
			It("Should delete productOption", func() {
				err := obtainedProductOption.SetID(productOption.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = productOption.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedProductOption.FetchById()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
			})

		})
	})
})
