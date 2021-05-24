package models_test

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image Model", func() {
	var image models.Image

	BeforeEach(func() {
		image = CreateImage()
	})

	AfterEach(func() {
		image.Delete()
	})

	Describe("Basic CRUD test", func() {
		It("Should be stored into db", func() {
			var obtainedImage models.Image
			obtainedImage.SetID(image.GetID())
			err := obtainedImage.FetchById()
			Expect(err).To(BeNil())
			Expect(obtainedImage.FullPath).To(Equal(image.FullPath))
		})
	})
})
