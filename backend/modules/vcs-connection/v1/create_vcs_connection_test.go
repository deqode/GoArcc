package vcs_connection_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Create Vcs Connection Test", func() {

	Describe("Creating a vcs connection", func() {

		//Negative Test Cases
		By("By a internal RPC Call")
		Context("Get an error when request object is nil", func() {
			It("should return nil exception", func() {
			})
		})
		Context("Get an error when vcs connection in request object is nil", func() {
			It("should return nil exception", func() {
			})
		})
		Context("Get an error when vcs object is invalid", func() {
			It("should return an error when account id is empty", func() {
			})
		})
		Context("Get an error when vcs object is invalid", func() {
			It("should return error when username  is empty", func() {
			})
		})

		Context("Get an error when vcs object is invalid", func() {
			It("should return an error when account id is invalid", func() {
			})
		})
		Context("Validate Vcs Object", func() {
			It("should return  an error when username  is invalid", func() {
			})
		})

		Context("Create a Vcs when request object is valid", func() {
			It("Fail if response object is nil", func() {
			})
		})
		Context("Create a Vcs when request object is valid", func() {
			It("Fail if username is empty", func() {
			})
		})
		Context("Create a Vcs when request object is valid", func() {
			It("Fail if account Id is empty", func() {
			})
		})

		Context("Create a Vcs when request object is valid", func() {
			It("Fail if access token is empty", func() {
			})
		})

		Context("Create a Vcs when request object is valid", func() {
			It("Fail if provider  is empty", func() {
			})
		})
		Context("Create a Vcs when request object is valid", func() {
			It("Fail if id   is empty", func() {
			})
		})

	})
})
