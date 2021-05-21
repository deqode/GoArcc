package external_svc_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Get Vcs Connection Test", func() {

	Describe("Get vcs connection", func() {

		//Negative Test Cases
		By("By a internal RPC Call")
		Context("Get an error when request object is nil", func() {
			It("should return nil exception", func() {
			})
		})
		Context("Get an error when account id is empty ", func() {
			It("should return nil exception", func() {
			})
		})
		Context("Get an error when id is empty", func() {
			It("should return an error when account id is empty", func() {
			})
		})

		Context("Get an error when id is invalid or account id is invalid", func() {
			It("should return an no record found error ", func() {
			})
		})

		Context("When Request object is valid", func() {
			It("Fail in case id is not present in response", func() {
			})
		})
		Context("When Request object is valid", func() {
			It("Fail in case provider is not present in response", func() {
			})
		})

		Context("When Request object is valid", func() {
			It("Fail in case username is not present in response", func() {
			})
		})
	})
})
