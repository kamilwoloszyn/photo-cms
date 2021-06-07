package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("DeliveryMethod model", func() {
	var (
		deliveryMethod  models.DeliveryMethod
		delivery        models.Delivery
		anotherDelivery models.Delivery
	)
	BeforeEach(func() {
		deliveryMethod = CreateDeliveryMethod()
	})

	AfterEach(func() {
		deliveryMethod.Delete()
	})
	Describe("Basic crud operations", func() {
		Context("Create or Update operations", func() {
			var obtainedDeliveryMethod models.DeliveryMethod

			It("Should be stored into db", func() {
				err := obtainedDeliveryMethod.SetID(deliveryMethod.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedDeliveryMethod.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedDeliveryMethod.Name).To(Equal(deliveryMethod.Name))
			})

			It("Should update an existing delivery method", func() {
				err := obtainedDeliveryMethod.SetID(deliveryMethod.GetID())
				Expect(err).ShouldNot(HaveOccurred())
				deliveryMethod.FixedPrice = 40
				err = deliveryMethod.UpdateInstance()
				Expect(err).ShouldNot(HaveOccurred())
				err = obtainedDeliveryMethod.FetchById()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(obtainedDeliveryMethod.FixedPrice).To(Equal(deliveryMethod.FixedPrice))
			})
			Context("Delete operations", func() {
				It("Should delete a delivery method from db", func() {
					err := obtainedDeliveryMethod.SetID(deliveryMethod.GetID())
					Expect(err).ShouldNot(HaveOccurred())
					err = deliveryMethod.Delete()
					Expect(err).ShouldNot(HaveOccurred())
					err = obtainedDeliveryMethod.FetchById()
					Expect(err).To(Equal(gorm.ErrRecordNotFound))
				})
			})
		})
	})
	Describe("Relationship tests", func() {
		BeforeEach(func() {
			delivery = CreateDelivery(deliveryMethod)
		})
		AfterEach(func() {
			delivery.Delete()
		})
		Context("One delivery", func() {
			It("Should get one delivery", func() {
				err := deliveryMethod.GetDeliveries()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(deliveryMethod.Delivery[0].DestinationAddress).To(Equal(delivery.DestinationAddress))
			})
		})
		Context("More deliveries", func() {
			BeforeEach(func() {
				anotherDelivery = models.Delivery{
					ShippedVia:               "Jakub",
					TrackingCode:             "123793002742342",
					DestinationPostalCode:    "37-620",
					DestinationCountryRegion: "Podkarpackie",
					DestinationAddress:       "Truskawkowa 1",
					DestinationCity:          "Horyniec-Zdroj",
				}
				CreateCustomDelivery(deliveryMethod, &anotherDelivery)
			})
			AfterEach(func() {
				anotherDelivery.Delete()
				delivery.Delete()
				deliveryMethod.Delete()

			})
			It("Should get two deliveries", func() {
				err := deliveryMethod.GetDeliveries()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(deliveryMethod.Delivery)).To(Equal(2))
			})
		})
	})

})
