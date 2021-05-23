package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Payment Method model", func() {
	var paymentMethod models.PaymentMethod

	BeforeEach(func() {
		paymentMethod = CreatePaymentMethod()
	})

	AfterEach(func() {
		paymentMethod.Delete()
	})
	Context("Basic crud testing", func() {
		It("Should be into db", func() {
			var obtainedPaymentMethod models.PaymentMethod
			err := obtainedPaymentMethod.SetID(paymentMethod.GetID())
			Expect(err).To(BeNil())
			err = obtainedPaymentMethod.FetchByID()
			Expect(err).To(BeNil())
			Expect(obtainedPaymentMethod.KeyMd5).To(Equal(paymentMethod.KeyMd5))
		})
	})
})
