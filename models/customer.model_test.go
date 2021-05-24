package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Customer model", func() {
	var customer models.Customer

	BeforeEach(func() {
		customer = CreateCustomer()
	})

	AfterEach(func() {
		customer.Delete()
	})
	Describe("Crud Basic", func() {
		It("Should be stored into db", func() {
			var obtainedCustomer models.Customer
			err := obtainedCustomer.SetID(customer.GetID())
			Expect(err).To(BeNil())
			err = obtainedCustomer.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedCustomer.City).To(Equal(customer.City))
		})
	})
})
