package external_svc_test

import (
	"alfred/modules/account/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var _ = Describe("UpdateUserAccounts", func() {
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

	Describe("Update an account", func() {
		//Negative Test Cases
		By("By a internal or external RPC Call")
		Context("Get an error when request object is nil", func() {
			It("should return nil exception", func() {
				_, err := accountServer.UpdateAccount(ctx, &pb.UpdateAccountRequest{Account: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Account to update is not provided")))
			})
		})

		Context("Get an error when update mask is incorrect ", func() {
			It("should return failed precondition error", func() {
			})
		})

		Context("Get an error when update mask contain id ", func() {
			It("should return failed precondition error", func() {
			})
		})
		Context("Get an error when id is incorrect", func() {
			It("should return failed precondition error", func() {
				_, err := accountServer.UpdateAccount(ctx, &pb.UpdateAccountRequest{Account: &pb.Account{
					Id:     "7328-dshj-328-shds",
					Slug:   account.Slug,
					UserId: account.UserId,
				}})
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})
		})
		Context("Get an error when account does not exist", func() {
			It("should return failed precondition error when profile not exist", func() {
			})
		})

		Context("Update account when update mask is correct and id is correct", func() {
			It("user Account should be updated successfully", func() {
			})
		})

		Context("Update user account when update mask is correct and id is correct", func() {
			It("check the response of object that account is updated or not", func() {
			})
		})

	})
})
