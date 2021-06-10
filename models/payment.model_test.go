package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Payment model", func() {
	var (
		paymentMethod models.PaymentMethod
		payment       models.Payment
	)

	BeforeEach(func() {
		paymentMethod = CreatePaymentMethod()
		payment = CreatePayment(paymentMethod)
	})
	AfterEach(func() {
		payment.Delete()
		paymentMethod.Delete()
	})
	Describe("Basic crud testing", func() {
		var (
			obtainedPayment models.Payment
		)

		BeforeEach(func() {
			obtainedPayment = models.Payment{}
		})
		Context("Add or update an operations", func() {
			It("Should create a new payment in db", func() {
				err := obtainedPayment.SetID(payment.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedPayment.FetchByID()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedPayment.PaymentAmount).To(Equal(payment.PaymentAmount))
			})
			It("Should update an existing payment", func() {
				err := obtainedPayment.SetID(payment.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				payment.PaymentAmount = 100
				err = payment.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedPayment.FetchByID()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedPayment.PaymentAmount).To(Equal(payment.PaymentAmount))
			})
		})
		Context("Delete operations", func() {
			It("Should delete existing payment", func() {
				err := obtainedPayment.SetID(payment.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = payment.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedPayment.FetchByID()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
				Expect(obtainedPayment.PaymentMethodId).To(Equal(""))
			})
		})

	})

	Describe("Relationship test", func() {
		var (
			deliveryMethod models.DeliveryMethod
			delivery       models.Delivery
			order          models.Order
			err            error
		)

		BeforeEach(func() {
			deliveryMethod = CreateDeliveryMethod()
			delivery = CreateDelivery(deliveryMethod)
			order = CreateOrder(payment, delivery)
		})
		JustAfterEach(func() {
			err = payment.AssignTo(&order)
		})

		AfterEach(func() {
			order.Delete()
			payment.Delete()
			delivery.Delete()
		})
		Context("One order", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(payment.Order.Price).To(Equal(order.Price))
		})
	})
})
