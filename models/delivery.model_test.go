package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Devliery model", func() {
	var deliveryMethod models.DeliveryMethod
	var delivery models.Delivery
	BeforeEach(func() {
		deliveryMethod = CreateDeliveryMethod()
		delivery = CreateDelivery(deliveryMethod)
	})

	AfterEach(func() {
		delivery.Delete()
		deliveryMethod.Delete()
	})

	Describe("Basic crud operations", func() {
		It("Should be in db", func() {
			obtainedDelivery := models.Delivery{}
			err := obtainedDelivery.SetID(delivery.GetID())
			Expect(err).To(BeNil())
			err = obtainedDelivery.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedDelivery.DestinationCity).To(Equal(delivery.DestinationCity))
		})
	})

})
