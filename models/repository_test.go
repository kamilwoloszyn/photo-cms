package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository test", func() {

	Context("Basic connection to db", func() {
		It("Should set handler", func() {
			err := models.GetHandler().Error
			Expect(err).To(BeNil())
			Expect(models.GetHandler()).NotTo(BeNil())
		})
	})
})
