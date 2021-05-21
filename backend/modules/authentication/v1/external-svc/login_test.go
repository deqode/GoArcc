package external_svc_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Login", func() {
	Describe("Getting External VCS login url", func() {
		By("By Rpc Calls")
		Context("Get an error when request object is nil", func() {
			It("Failed precondition err must be responded", func() {

			})
		})

		// positive test case
		Context("Get an error if returned url is empty", func() {
			It("Failed precondition err must be responded", func() {
			})
		})
	})
})
