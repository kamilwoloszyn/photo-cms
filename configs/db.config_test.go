package configs_test

import (
	"github.com/kamilwoloszyn/photo-cms/configs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Env variables test", func() {

	It("Should get variables", func() {
		dbConfig, err := configs.LoadDbConfig()
		Expect(err).To(BeNil())
		Expect(dbConfig.Database).NotTo(BeEmpty())
		Expect(dbConfig.Dbpassword).NotTo(BeEmpty())
		Expect(dbConfig.Dbuser).NotTo(BeEmpty())
		Expect(dbConfig.Port).NotTo(BeEmpty())
		Expect(dbConfig.Host).NotTo(BeEmpty())
		Expect(dbConfig.HandlerName).NotTo(BeEmpty())
	})
})
