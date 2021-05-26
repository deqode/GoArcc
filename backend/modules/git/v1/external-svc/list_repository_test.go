package external_svc_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("List Repository", func() {

	Context("Get an error when request is null", func() {
		It("Should get an error", func() {

		})
	})

	Context("Get an error When Request has empty fields", func() {
		It("Should get an error if account_id is empty", func() {
		})

		It("Should get an error if repository name is empty", func() {
		})
	})

	Context("Get an error When Request has invalid fields", func() {
		It("Should get an error if account_id is wrong", func() {
		})
		It("Should get an error if provider is wrong", func() {
		})
	})
	Context("Get an error When Access Token is empty", func() {
		It("Should get an error", func() {
		})
	})

	//Positive Test Cases
	Context("Get valid response when request is valid", func() {
		It("Should get an error if user has 0 repository ", func() {
		})
	})

})
