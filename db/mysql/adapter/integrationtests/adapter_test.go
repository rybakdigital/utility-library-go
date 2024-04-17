package utility_library_go_test

import (
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rybakdigital/utility-library-go/db/mysql/adapter"
)

var _ = Describe("Db/Mysql/Adapter/Integrationtests/Adapter", func() {
	Context("connection available", func() {
		It("should fail to establish connection, DSN not configured properly", func() {
			a := adapter.DefaultAdapter()
			_, err := a.IsDbConnectionAvailable()
			Expect(err).To(HaveOccurred())
		})

		It("should be able to connect to db", func() {
			logger := log.Default()
			logger.SetPrefix("mysql-adapter")
			c := adapter.NewConfig(adapter.NewDsn("testuser", "testpassword", "test").SetPort("8886"), 5, 5)
			a := adapter.NewAdapter(c, logger)
			_, err := a.IsDbConnectionAvailable()
			Expect(err).ToNot(HaveOccurred())
		})

		It("should establish connection", func() {
			logger := log.Default()
			logger.SetPrefix("mysql-adapter")
			c := adapter.NewConfig(adapter.NewDsn("testuser", "testpassword", "test").SetPort("8886"), 5, 5)
			a := adapter.NewAdapter(c, logger)
			err := a.Connect()
			Expect(err).ToNot(HaveOccurred())
			Expect(a.IsConnected).To(BeTrue())
		})
	})
})
