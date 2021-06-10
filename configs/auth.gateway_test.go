package configs_test

import (
	"github.com/kamilwoloszyn/photo-cms/configs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth gateway", func() {
	var (
		config configs.AuthGatewayConfig
	)
	Context("Config", func() {
		It("Should load config correctly", func() {
			err := config.Load()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(config.Host).ToNot(BeEmpty())
			Expect(config.Port).ToNot(BeEmpty())
		})
	})
	Context("Connection good config", func() {
		It("Should pass", func() {
			gateway, err := config.TestConnection()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(gateway.Url).ToNot(BeEmpty())
		})

	})

	Context("Connection, bad config", func() {
		var (
			badConfig configs.AuthGatewayConfig
		)
		BeforeEach(func() {
			badConfig = configs.AuthGatewayConfig{}
		})
		It("Should fail", func() {
			_, err := badConfig.TestConnection()
			Expect(err).Should(HaveOccurred())
		})
	})

})
