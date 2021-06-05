package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Customer model", func() {
	var customer models.Customer

	BeforeEach(func() {
		customer = CreateCustomer()
	})

	AfterEach(func() {
		customer.Delete()
	})
	Describe("Crud Basic", func() {
		var obtainedCustomer models.Customer
		Context("Create or update operations", func() {
			It("Should create new customer", func() {
				err := obtainedCustomer.SetID(customer.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedCustomer.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedCustomer.City).To(Equal(customer.City))
			})
			It("Should update an existing customer", func() {
				var modifiedCustomer models.Customer
				err := modifiedCustomer.SetID(customer.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				customer.City = "Rzeszow"
				err = customer.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = modifiedCustomer.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(modifiedCustomer.City).To(Equal(customer.City))
			})

		})
		Context("Delete operation", func() {
			It("Should delete an existing customer", func() {
				var emptyCustomer models.Customer
				err := emptyCustomer.SetID(customer.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = customer.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = emptyCustomer.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(emptyCustomer.Address).To(Equal(""))
			})
		})
		Describe("Relationship test", func() {
			var product models.Product
			var anotherProduct models.Product
			var category models.Category
			var image models.Image
			var anotherImage models.Image

			BeforeEach(func() {
				category = CreateCategory()
				image = CreateImage()
				product = CreateProductWithoutOrder(category, image, customer)
				anotherProduct = models.Product{
					UnitPrice:   40,
					ProductName: "sample_image1",
					Quantity:    3,
				}
				anotherImage = models.Image{
					Name:     "generated_img1",
					FullPath: "/tmp/generated_img1",
					Size:     10300,
				}

			})

			AfterEach(func() {
				product.Delete()
				image.Delete()
				category.Delete()
				customer.Delete()
			})

			Context("One product", func() {
				It("Should fetch one product to the customer", func() {
					err := customer.GetProducts()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(customer.Products)).To(Equal(1))
					Expect(customer.Products[0].ProductName).To(Equal(product.ProductName))
				})
			})
			Context("More than one product", func() {
				JustBeforeEach(func() {
					CreateCustomImage(&anotherImage)
					CreateCustomProductWithoutOrder(category, anotherImage, customer, &anotherProduct)
				})
				AfterEach(func() {
					anotherProduct.Delete()
					anotherImage.Delete()
				})
				It("Should fetch two products", func() {
					err := customer.GetProducts()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(customer.Products)).To(Equal(2))
				})
			})
		})
	})
})
