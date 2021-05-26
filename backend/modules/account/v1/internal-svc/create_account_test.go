package internal_svc_test

import (
	"alfred/modules/account/v1/pb"
	usrPb "alfred/modules/user-profile/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateAccount", func() {
	var (
		accountServer pb.AccountInternalServer
		ctx           context.Context
		usrProfile    *usrPb.UserProfile
	)
	BeforeEach(func() {
		accountServer = AccountServerTest
		ctx = CtxTest
		usrProfile = UsrProfile
	})

	Describe("Describe:Creating Account Correspond to user", func() {
		Context("Context:When user_id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.CreateAccount(ctx, &pb.CreateAccountRequest{Account: &pb.Account{
					UserId: "",
				}})
				Expect(err.(pb.CreateAccountRequestValidationError).Cause().(pb.AccountValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
		Context("Context:When request is nil", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.CreateAccount(ctx, &pb.CreateAccountRequest{Account: nil})
				Expect(err.(pb.CreateAccountRequestValidationError).Reason()).Should(Equal("value is required"))
			})
		})

		// Positive Test Case
		Context("Context:Happy Path", func() {
			It("It: should create the account", func() {
				res, err := accountServer.CreateAccount(ctx, &pb.CreateAccountRequest{Account: &pb.Account{
					Slug:   usrProfile.Sub + "_" + usrProfile.UserName,
					UserId: usrProfile.Id,
				}})
				Expect(err).Should(BeNil())
				Expect(res.UserId).ShouldNot(BeEmpty())
			})
		})
	})
})
