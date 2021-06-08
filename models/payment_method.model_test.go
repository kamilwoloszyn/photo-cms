package models_test

import (
	"time"

	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Payment Method model", func() {
	var (
		paymentMethod models.PaymentMethod
	)

	BeforeEach(func() {
		paymentMethod = CreatePaymentMethod()
	})

	Describe("Basic crud tests", func() {
		var obtainedPaymentMethod models.PaymentMethod
		BeforeEach(func() {
			obtainedPaymentMethod = models.PaymentMethod{}
		})
		AfterEach(func() {
			paymentMethod.Delete()
		})
		Context("Create or update operations", func() {
			It("Should create a new payment method", func() {
				err := obtainedPaymentMethod.SetID(paymentMethod.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedPaymentMethod.FetchByID()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedPaymentMethod.KeyMd5).To(Equal(paymentMethod.KeyMd5))
			})

			It("Should update an existing payment method", func() {
				err := obtainedPaymentMethod.SetID(paymentMethod.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				paymentMethod.PosId = "1111111"
				err = paymentMethod.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedPaymentMethod.FetchByID()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedPaymentMethod.PosId).To(Equal(paymentMethod.PosId))
			})
		})

		Context("Delete operation", func() {
			It("Should delete existing payment method", func() {
				err := obtainedPaymentMethod.SetID(paymentMethod.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = paymentMethod.Delete()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedPaymentMethod.FetchByID()
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
				Expect(obtainedPaymentMethod.KeyMd5).To(Equal(""))
			})
		})

	})

	Describe("Relationship test", func() {
		var (
			payment models.Payment
		)
		BeforeEach(func() {
			payment = CreatePayment(paymentMethod)
		})
		AfterEach(func() {
			payment.Delete()
			paymentMethod.Delete()
		})
		Context("One payment", func() {
			It("Should have one payment", func() {
				err := paymentMethod.GetPayments()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(paymentMethod.Payment)).To(Equal(1))
				Expect(paymentMethod.Payment[0].PaymentAmount).To(Equal(payment.PaymentAmount))
				Expect(paymentMethod.Payment[0].PaymentMethodId).To(Equal(paymentMethod.GetID()))
			})
		})

		Context("Two payments", func() {
			var (
				payment2 models.Payment
				time     time.Time
			)

			BeforeEach(func() {
				time = time.AddDate(2021, 12, 10)
				payment2 = models.Payment{
					PaymentDate:   &time,
					PaymentAmount: 1200,
				}
			})
			JustBeforeEach(func() {
				CreateCustomPayment(paymentMethod, &payment2)
			})
			AfterEach(func() {
				payment2.Delete()
			})
			It("Should have two payents", func() {
				err := paymentMethod.GetPayments()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(paymentMethod.Payment)).To(Equal(2))
			})
		})
	})
})
