package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
		It("Should contain payment in db", func() {
			var obtainedPayment models.Payment
			err := obtainedPayment.SetID(payment.GetID())
			Expect(err).To(BeNil())
			err = obtainedPayment.FetchByID()
			Expect(err).To(BeNil())
			Expect(obtainedPayment.PaymentAmount).To(Equal(payment.PaymentAmount))
		})
	})
})
