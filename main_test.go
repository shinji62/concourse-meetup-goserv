package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/shinji62/concourse-meetup-goserv"
	"io/ioutil"
	"net/http"
)

var _ = Describe("Main", func() {
	var backend *ghttp.Server
	var getReq *http.Request

	BeforeEach(func() {
		backend = ghttp.NewServer()
		backend.AppendHandlers(main.BarHandler)
		getReq, _ = http.NewRequest("GET", backend.URL()+"/bar", nil)

	})

	It("Should respond to Bar", func() {
		res, err := http.DefaultClient.Do(getReq)
		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(res.StatusCode).Should(Equal(200))
		Expect(string(body)).Should(Equal("Hello, World"))

	})
	AfterEach(func() {
		backend.Close()
	})

})
