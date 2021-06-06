package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Delivery model", func() {
	var (
		deliveryMethod models.DeliveryMethod
		delivery       models.Delivery
		paymentMethod  models.PaymentMethod
		payment        models.Payment
		order          models.Order
	)

	BeforeEach(func() {
		deliveryMethod = CreateDeliveryMethod()
		delivery = CreateDelivery(deliveryMethod)
	})

	Describe("Basic crud operations", func() {
		Context("Update or delete operations", func() {
			AfterEach(func() {
				delivery.Delete()
				deliveryMethod.Delete()
			})
			It("Should create new delivery", func() {
				obtainedDelivery := models.Delivery{}
				err := obtainedDelivery.SetID(delivery.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedDelivery.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedDelivery.DestinationCity).To(Equal(delivery.DestinationCity))
			})
			It("Should update existing delivery", func() {
				obtainedDelivery := models.Delivery{}
				err := obtainedDelivery.SetID(delivery.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				delivery.DestinationAddress = "Poziomkowa 20"
				err = delivery.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedDelivery.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedDelivery.DestinationAddress).To(Equal(delivery.DestinationAddress))
			})
		})
		Context("Delete operations", func() {
			AfterEach(func() {
				deliveryMethod.Delete()
			})
			It("Should delete delivery", func() {
				obtainedDelivery := models.Delivery{}
				err := obtainedDelivery.SetID(delivery.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = delivery.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedDelivery.FetchById()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
			})
		})

	})

	Describe("Relationship test", func() {
		BeforeEach(func() {
			paymentMethod = CreatePaymentMethod()
			payment = CreatePayment(paymentMethod)
			order = CreateOrder(payment, delivery)
		})
		AfterEach(func() {
			order.Delete()
		})
		Context("One order", func() {
			It("Should fetch exactly one order", func() {
				err := delivery.GetOrders()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(delivery.Order)).To(Equal(1))
				Expect(delivery.Order[0].Price).To(Equal(order.Price))
			})

		})
	})

})
