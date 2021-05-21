package external_svc_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Callback Test", func() {
	Describe("Callback Test", func() {
		By("By Rpc Calls")
		Context("Get an error when request object is empty", func() {
			It("Failed precondition err must be responded", func() {

			})
		})
		Context("Get an error when Request object is invalid", func() {
			It("validation err must be responded when account length is invalid", func() {

			})
			It("Failed precondition err must be responded when code length is invalid", func() {

			})
		})

		Context("Get an error when account id is invalid", func() {
			It("Failed precondition err must be responded", func() {

			})
		})
		Context("Get an error when code is invalid", func() {
			It("Failed precondition err must be responded", func() {

			})
		})

		Context("Successfully Get Response when Request object is valid", func() {
			It("Check whether response id is valid or not", func() {

			})
			It("Check whether username is valid or not", func() {

			})
			It("Check whether label is valid or not", func() {

			})
			It("Check whether provider of request matches from provider of response", func() {

			})
		})
	})
})
