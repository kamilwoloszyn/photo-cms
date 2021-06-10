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
		product = CreateProductWithoutOrder(category, image, customer)
		productOption = CreateProductOption(product, optionValue)
		deliveryMethod = CreateDeliveryMethod()
		delivery = CreateDelivery(deliveryMethod)
		paymentMethod = CreatePaymentMethod()
		payment = CreatePayment(paymentMethod)
		order = CreateOrder(payment, delivery)
	})

	AfterEach(func() {
		productOption.Delete()
		product.Delete()
		order.Delete()
		payment.Delete()
		paymentMethod.Delete()
		delivery.Delete()
		deliveryMethod.Delete()
		category.Delete()
		image.Delete()
		optionValue.Delete()
		option.Delete()
		customer.Delete()
	})

	Describe("Basic crud testing", func() {
		var obtainedOrder models.Order

		BeforeEach(func() {
			obtainedOrder = models.Order{}
		})

		Context("Create or update operations", func() {
			It("Should create a new order", func() {
				err := obtainedOrder.SetID(order.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedOrder.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedOrder.Price).To(Equal(order.Price))
			})
			It("Should update an existing order", func() {
				err := obtainedOrder.SetID(order.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				order.Price = 100
				err = order.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedOrder.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedOrder.Price).To(Equal(order.Price))
				Expect(obtainedOrder.PaymentId).To(Equal(order.PaymentId))
				Expect(obtainedOrder.DeliveryId).To(Equal(order.DeliveryId))
			})
		})
		Context("Delete operations", func() {
			It("Should delete an existing order", func() {
				err := obtainedOrder.SetID(order.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = order.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedOrder.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedOrder.PaymentId).To(BeNil())
			})
		})

	})

	Describe("Relationship testing", func() {

		Context("One product", func() {
			It("Should contain one product", func() {
				err := order.GetProducts()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(order.Product)).To(Equal(1))
			})
		})

		Context("More than one product", func() {
			var (
				image2   models.Image
				product2 models.Product
			)
			BeforeEach(func() {
				image2 = models.Image{
					Name:     "custom_image",
					FullPath: "/tmp/custom_image",
					Size:     1300,
				}
			})

			JustBeforeEach(func() {
				CreateCustomImage(&image2)
				CreateCustomProductWithOrder(category, image2, customer, order, &product2)
				order.AssignTo(&product)
				order.AssignTo(&product2)
			})
			AfterEach(func() {
				product2.Delete()
				image2.Delete()
			})
			It("Should contain two products", func() {
				err := order.GetProducts()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(order.Product)).To(Equal(2))
			})
		})

		Context("Payment details", func() {
			var (
				checkerPayment models.Payment
				fetchedPayment models.Payment
			)
			It("Should fetch correct details of payment", func() {
				err := checkerPayment.SetID(*order.PaymentId)
				Expect(err).ShouldNot(HaveOccurred())
				err = checkerPayment.FetchByID()
				Expect(err).ShouldNot(HaveOccurred())
				err = order.GetPaymentDetails(&fetchedPayment)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(fetchedPayment.PaymentAmount).To(Equal(checkerPayment.PaymentAmount))
				Expect(fetchedPayment.PaymentMethodId).To(Equal(checkerPayment.PaymentMethodId))
				Expect(fetchedPayment.GetID()).To(Equal(checkerPayment.GetID()))
			})
		})

		Context("Delivery details", func() {
			var (
				checkerDelivery models.Delivery
				fetchedDelivery models.Delivery
			)
			It("Should fetch correct details of delivery", func() {
				err := checkerDelivery.SetID(*order.DeliveryId)
				Expect(err).ShouldNot(HaveOccurred())
				err = checkerDelivery.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				err = order.GetDeliveryDetails(&fetchedDelivery)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(fetchedDelivery.DeliveryMethodId).To(Equal(checkerDelivery.DeliveryMethodId))
				Expect(fetchedDelivery.DestinationAddress).To(Equal(checkerDelivery.DestinationAddress))
				Expect(fetchedDelivery.DestinationCity).To(Equal(checkerDelivery.DestinationCity))
				Expect(fetchedDelivery.TrackingCode).To(Equal(checkerDelivery.TrackingCode))
			})
		})
	})
})
