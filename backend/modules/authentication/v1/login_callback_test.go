package authentication

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Login Callback Test", func() {
	Describe("Login Callback Test", func() {
		By("By Rpc Calls")

		Context("Get an error when request object is nil", func() {
			It("Failed precondition err must be responded", func() {

			})
		})
		Context("Get an error if state or code is empty", func() {
			It("Failed precondition err must be responded", func() {
			})
		})

		Context("Get an error if  code is invalid", func() {
			It("Unauthenticated error must be thrown", func() {
			})
		})

		Context("Get an error if user profile is not present", func() {
			It("gorm error must be thrown", func() {
			})
		})

		Context("Get CallBack Response when request object is valid", func() {
			It("Throw error if response object is nil", func() {
			})
			It("Throw error if user id is empty", func() {
			})
			It("Throw error if Access token is empty", func() {
			})
			It("Throw error if raw id  is empty", func() {
			})
		})
	})
})
