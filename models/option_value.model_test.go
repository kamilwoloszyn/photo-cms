package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("OptionValue model", func() {
	var (
		option              models.Option
		optionValue         models.OptionValue
		productOption       models.ProductOption
		productOption2      models.ProductOption
		obtainedOptionValue models.OptionValue
		product             models.Product
		category            models.Category
		image               models.Image
		customer            models.Customer
		product2            models.Product
		image2              models.Image
	)

	BeforeEach(func() {
		option = models.Option{}
		optionValue = models.OptionValue{}
		productOption = models.ProductOption{}
	})

	JustBeforeEach(func() {
		option = CreateOption()
		optionValue = CreateOptionValue(option)
	})
	Describe("Basic crud testing", func() {

		Context("Create or update methods", func() {
			var obtainedOptionValue models.OptionValue
			BeforeEach(func() {
				option = models.Option{}
				optionValue = models.OptionValue{}
				obtainedOptionValue = models.OptionValue{}
			})
			AfterEach(func() {
				optionValue.Delete()
				option.Delete()
			})
			It("Should create a new option value", func() {
				err := obtainedOptionValue.SetID(optionValue.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedOptionValue.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedOptionValue.Value).To(Equal(optionValue.Value))
			})

			It("Should update an existing option value", func() {
				obtainedOptionValue.SetID(optionValue.GetID())
				optionValue.Value = "newValue"
				err := optionValue.UpdateInstance()
				Expect(err).NotTo(HaveOccurred())
				err = obtainedOptionValue.FetchById()
				Expect(err).ToNot(HaveOccurred())
				Expect(obtainedOptionValue.Value).To(Equal(optionValue.Value))
			})

		})
		Context("Delete operations", func() {
			AfterEach(func() {
				option.Delete()
			})
			It("Should delete an option value", func() {
				err := obtainedOptionValue.SetID(optionValue.GetID())
				Expect(err).ToNot(HaveOccurred())
				err = optionValue.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedOptionValue.FetchById()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
				Expect(obtainedOptionValue.Value).Should(BeEmpty())
			})
		})
	})

	Describe("Relationship tests", func() {
		BeforeEach(func() {
			product2 = models.Product{
				UnitPrice:   30,
				ProductName: "sample_image2",
				Quantity:    2,
			}
			image2 = models.Image{
				Name:     "generated_img_2",
				FullPath: "/tmp/generated_img_2",
				Size:     10020,
			}
		})
		JustBeforeEach(func() {
			category = CreateCategory()
			image = CreateImage()
			customer = CreateCustomer()
			product = CreateProductWithoutOrder(category, image, customer)
			CreateCustomImage(&image2)
			CreateCustomProductWithoutOrder(category, image, customer, &product2)
			productOption = CreateProductOption(product, optionValue)
			productOption2 = CreateProductOption(product2, optionValue)
		})

		AfterEach(func() {
			productOption.Delete()
			productOption2.Delete()
			product.Delete()
			product2.Delete()
			customer.Delete()
			image.Delete()
			image2.Delete()
			category.Delete()
			optionValue.Delete()
			option.Delete()
		})
		Context("Two product options", func() {
			It("Should contain 2 product option", func() {
				err := optionValue.GetProductOptions()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(optionValue.ProductOption)).To(Equal(2))
			})

		})
	})
})
