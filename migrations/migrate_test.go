package migrations_test

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Db Automigrate", func() {
	var db *gorm.DB

	BeforeEach(func() {

	})

	AfterEach(func() {
		db.Close()
	})
	Context("Migration into db", func() {
		It("Should migrate into db", func() {
			Migrate()
		})
	})

})
