package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Product model", func() {
	var (
		category      models.Category
		image         models.Image
		option        models.Option
		optionValue   models.OptionValue
		productOption models.ProductOption
		customer      models.Customer
		product       models.Product
	)

	BeforeEach(func() {
		category = CreateCategory()
		image = CreateImage()
		option = CreateOption()
		optionValue = CreateOptionValue(option)
		customer = CreateCustomer()
		product = CreateProductWithoutOrder(category, image, customer)

	})
	AfterEach(func() {
		product.Delete()
		customer.Delete()
		optionValue.Delete()
		category.Delete()
		image.Delete()
		option.Delete()
	})
	Describe("Basic crud testing", func() {
		var (
			productObtained models.Product
		)
		BeforeEach(func() {
			productObtained = models.Product{}
		})
		Context("Create or update operations", func() {
			It("Should create a new product", func() {
				err := productObtained.SetID(product.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = productObtained.FetchByID()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(productObtained.ProductName).To(Equal(product.ProductName))
			})
			It("Should update an existing product", func() {
				err := productObtained.SetID(product.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				product.ProductName = "changed"
				err = product.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = productObtained.FetchByID()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(productObtained.ProductName).To(Equal(product.ProductName))
			})
		})

		Context("Delete operations", func() {
			It("Should delete an existing product", func() {
				err := productObtained.SetID(product.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = productObtained.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = productObtained.FetchByID()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
				Expect(productObtained.ProductName).To(Equal(""))
			})
		})

	})

	Describe("Relationship tests", func() {

		BeforeEach(func() {
			productOption = CreateProductOption(product, optionValue)
		})

		AfterEach(func() {
			productOption.Delete()
		})
		Context("One product option", func() {
			It("Should get one product option", func() {
				err := product.GetProductOptions()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(product.ProductOption)).To(Equal(1))
				Expect(product.ProductOption[0].OptionValueId).To(Equal(optionValue.GetID()))
			})
		})

		Context("Two product options", func() {
			var (
				option2        models.Option
				optionValue2   models.OptionValue
				productOption2 models.ProductOption
			)

			BeforeEach(func() {
				optionValue2 = models.OptionValue{
					Value:      "Normal",
					ExtraPrice: 0,
				}
				option2 = models.Option{
					Name: "Material",
				}
				productOption2 = models.ProductOption{}
			})

			JustBeforeEach(func() {
				CreateCustomOption(&option2)
				CreateCustomOptionValue(option2, &optionValue2)
				productOption2 = CreateProductOption(product, optionValue2)
			})

			AfterEach(func() {
				productOption2.Delete()
				optionValue2.Delete()
				option2.Delete()
			})
			It("Should get two product options", func() {
				err := product.GetProductOptions()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(product.ProductOption)).To(Equal(2))
			})
		})
		Context("Customer details", func() {
			var (
				customerChecker  models.Customer
				obtainedCustomer models.Customer
			)
			It("Should get a customer details", func() {
				err := customerChecker.SetID(customer.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = customerChecker.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				err = product.GetCustomerDetails(&obtainedCustomer)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedCustomer.Address).To(Equal(customerChecker.Address))
				Expect(obtainedCustomer.City).To(Equal(customerChecker.City))
				Expect(obtainedCustomer.EmailAddress).To(Equal(customerChecker.EmailAddress))
				Expect(obtainedCustomer.PhoneNumber).To(Equal(customerChecker.PhoneNumber))
				Expect(obtainedCustomer.Regon).To(Equal(customerChecker.Regon))
			})
		})

		Context("OrderDetails", func() {
			var (
				paymentMethod  models.PaymentMethod
				payment        models.Payment
				deliveryMethod models.DeliveryMethod
				delivery       models.Delivery
				order          models.Order
				orderChecker   models.Order
				obtainedOrder  models.Order
				errAssign      error
			)
			BeforeEach(func() {
				paymentMethod = CreatePaymentMethod()
				payment = CreatePayment(paymentMethod)
				deliveryMethod = CreateDeliveryMethod()
				delivery = CreateDelivery(deliveryMethod)
				order = CreateOrder(payment, delivery)
				errAssign = order.AssignTo(&product)
			})
			AfterEach(func() {
				productOption.Delete()
				product.Delete()
				order.Delete()
				payment.Delete()
				paymentMethod.Delete()
				delivery.Delete()
				deliveryMethod.Delete()
			})
			It("Should get an order details", func() {
				Expect(errAssign).ShouldNot(HaveOccurred())
				err := orderChecker.SetID(order.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = orderChecker.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				err = product.GetOrderDetails(&obtainedOrder)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedOrder.Price).To(Equal(orderChecker.Price))
				Expect(obtainedOrder.Fvat).To(Equal(orderChecker.Fvat))
				Expect(obtainedOrder.PaymentId).To(Equal(orderChecker.PaymentId))
				Expect(obtainedOrder.ID).To(Equal(orderChecker.GetID()))
			})
		})

		Context("Category details ", func() {
			var (
				categoryChecker  models.Category
				obtainedCategory models.Category
			)
			It("Should get category details", func() {
				err := categoryChecker.SetID(category.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = categoryChecker.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				err = product.GetCategoryDetails(&obtainedCategory)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedCategory.CategoryName).To(Equal(categoryChecker.CategoryName))
			})
		})

		Context("Image details", func() {
			var (
				imageChecker  models.Image
				obtainedImage models.Image
			)
			It("Should get image details", func() {
				err := imageChecker.SetID(image.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = imageChecker.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				err = product.GetImageDetails(&obtainedImage)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedImage.FullPath).To(Equal(imageChecker.FullPath))
				Expect(obtainedImage.Name).To(Equal(imageChecker.Name))
				Expect(obtainedImage.Size).To(Equal(imageChecker.Size))
				Expect(obtainedImage.ID).To(Equal(imageChecker.GetID()))
			})
		})

	})

})
