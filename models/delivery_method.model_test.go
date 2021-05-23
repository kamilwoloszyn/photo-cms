package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeliveryMethod model", func() {
	var deliveryMethod models.DeliveryMethod
	BeforeEach(func() {
		deliveryMethod = CreateDeliveryMethod()
	})

	AfterEach(func() {
		deliveryMethod.Delete()
	})
	Context("Basic crud operations", func() {
		It("Should be stored into db", func() {
			var obtainedMethod models.DeliveryMethod
			obtainedMethod.SetID(deliveryMethod.GetID())
			err := obtainedMethod.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedMethod.Name).To(Equal(deliveryMethod.Name))
		})
	})
})
