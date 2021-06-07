package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Image Model", func() {
	var (
		image    models.Image
		category models.Category
		customer models.Customer
		product  models.Product
		product2 models.Product
		image2   models.Image
	)

	BeforeEach(func() {
		image = CreateImage()
	})

	Describe("Basic CRUD test", func() {
		var obtainedImage models.Image

		BeforeEach(func() {
			obtainedImage = models.Image{}
		})
		Context("Create or update method", func() {
			AfterEach(func() {
				image.Delete()
			})
			It("Should create a new image", func() {
				obtainedImage.SetID(image.GetID())
				err := obtainedImage.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedImage.FullPath).To(Equal(image.FullPath))
			})

			It("Should update an existing image", func() {
				err := obtainedImage.SetID(image.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				image.FullPath = "/tmp/image_new"
				err = image.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedImage.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedImage.FullPath).To(Equal(image.FullPath))
			})
		})
		Context("Delete method", func() {
			It("Should delete an image", func() {
				err := obtainedImage.SetID(image.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = image.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedImage.FetchById()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
				Expect(obtainedImage.FullPath).To(Equal(""))
			})

		})

	})

	Describe("Relationship test", func() {
		BeforeEach(func() {
			category = CreateCategory()
			customer = CreateCustomer()
			product = CreateProductWithoutOrder(category, image, customer)
		})

		AfterEach(func() {
			product.Delete()
			product2.Delete()
			customer.Delete()
			category.Delete()
			image.Delete()
			image2.Delete()
		})

		Context("One product", func() {
			It("Should fetch one product", func() {
				err := image.GetProduct()
				Expect(err).ShouldNot(HaveOccurred())

			})
		})
	})

	Describe("One to one relationship test", func() {
		BeforeEach(func() {
			product2 = models.Product{
				UnitPrice:   14,
				ProductName: "sample_image2",
				Quantity:    3,
			}
			image2 = models.Image{
				Name:     "other_img",
				FullPath: "/tmp/other_img",
				Size:     1200,
			}
		})

		JustBeforeEach(func() {
			category = CreateCategory()
			customer = CreateCustomer()
			product = CreateProductWithoutOrder(category, image, customer)
			CreateCustomImage(&image2)
		})
		AfterEach(func() {
			product2.Delete()
			product.Delete()
			customer.Delete()
			category.Delete()
			image.Delete()
			image2.Delete()
		})

		Context("Two products to one image", func() {
			It("Should fail", func() {
				CreateCustomProductWithoutOrder(category, image2, customer, &product2)
				err := image.AssignTo(&product2)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
