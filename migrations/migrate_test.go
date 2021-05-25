package migrations_test

import (
	"github.com/jinzhu/gorm"
	"github.com/kamilwoloszyn/photo-cms/migrations"
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Db Automigrate", func() {
	var db *gorm.DB

	BeforeEach(func() {
		err := models.Connect()
		Expect(err).To(BeNil())
		db = models.GetHandler()
	})

	AfterEach(func() {
		db.Close()
	})
	Describe("Migration into db", func() {
		It("Should migrate into db", func() {
			migrations.Migrate(db)
		})
	})

})
