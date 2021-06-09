package migrations_test

import (
	"github.com/kamilwoloszyn/photo-cms/migrations"
	"github.com/kamilwoloszyn/photo-cms/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Db Automigrate", func() {
	var db *gorm.DB

	BeforeEach(func() {
		err := models.ConnectDB()
		Expect(err).ShouldNot(HaveOccurred())
		db = models.GetHandler()
	})

	AfterEach(func() {
		statement, _ := db.DB()
		statement.Close()
	})
	Describe("Migration into db", func() {
		It("Should migrate into db", func() {
			err := migrations.Migrate(db)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

})
