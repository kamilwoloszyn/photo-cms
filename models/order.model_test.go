package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Order Model", func() {

	var (
		category       models.Category
		image          models.Image
		option         models.Option
		optionValue    models.OptionValue
		productOption  models.ProductOption
		customer       models.Customer
		deliveryMethod models.DeliveryMethod
		paymentMethod  models.PaymentMethod
		payment        models.Payment
		delivery       models.Delivery
		product        models.Product
		order          models.Order
	)

	BeforeEach(func() {
		category = CreateCategory()
		image = CreateImage()
		option = CreateOption()
		optionValue = CreateOptionValue(option)
		customer = CreateCustomer()
		product = CreateProductWithoutOrder(category, image, &customer)
		productOption = CreateProductOption(product, &optionValue)
		deliveryMethod = CreateDeliveryMethod()
		delivery = CreateDelivery(deliveryMethod)
		paymentMethod = CreatePaymentMethod()
		payment = CreatePayment(paymentMethod)
		order = CreateOrder(payment, delivery)
	})

	AfterEach(func() {
		order.Delete()
		payment.Delete()
		paymentMethod.Delete()
		delivery.Delete()
		deliveryMethod.Delete()
		productOption.Delete()
		product.Delete()
		category.Delete()
		image.Delete()
		optionValue.Delete()
		option.Delete()
		customer.Delete()
	})

	Describe("Basic crud testing", func() {
		It("Should be in db", func() {
			var obtainedOrder models.Order
			err := obtainedOrder.SetID(order.GetID())
			Expect(err).To(BeNil())
			err = obtainedOrder.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedOrder.Price).To(Equal(order.Price))
		})

	})

	// Describe("Relationship testing", func() {

	// })
})
