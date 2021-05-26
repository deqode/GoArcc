package external_svc_test

import (
	pb "alfred/modules/account/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Describe:GetAccount", func() {
	var (
		accountServer pb.AccountsServer
		ctx           context.Context
		account       *pb.Account
	)
	BeforeEach(func() {
		accountServer = AccountServerTest
		ctx = CtxTest
		account = Account
	})

	Describe("Get an account", func() {
		By("internal or external call")
		Context("Get an error when ctx is wrong", func() {
			It("Should get an error of permission denied", func() {

			})
		})
		Context("Get an error when id is empty", func() {
			It("it should return validation error", func() {
				_, err := accountServer.GetAccount(ctx, &pb.GetAccountRequest{Id: ""})
				Expect(err.(pb.GetAccountRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
		Context("Get an error when id is wrong", func() {
			It("should return not found error", func() {
				_, err := accountServer.GetAccount(ctx, &pb.GetAccountRequest{Id: "4389-jdk-8549"})
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})
		})

		Context("Get a record when id is provided", func() {
			It("should return requested field in the object", func() {
				_, err := accountServer.GetAccount(ctx, &pb.GetAccountRequest{Id: account.GetId()})
				Expect(err).Should(BeNil())
			})
		})
	})
})
