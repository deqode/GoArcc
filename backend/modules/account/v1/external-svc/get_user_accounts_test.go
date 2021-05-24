package external_svc_test

import (
	"alfred/modules/account/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("GetUserAccounts", func() {
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

	Describe("Describe:Categorizing with invalid id", func() {
		Context("Context:When user_id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.GetUserAccounts(ctx, &pb.GetUserAccountsRequest{UserId: ""})
				Expect(err.(pb.GetUserAccountsRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
		Context("Context:When user_id is wrong", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.GetUserAccounts(ctx, &pb.GetUserAccountsRequest{UserId: "wrongID"})
				Expect(gorm.ErrRecordNotFound).Should(Equal(err))
			})
		})

		Context("Get a record when id is provided", func() {
			It("should return requested field in the object", func() {
				_, err := accountServer.GetAccount(ctx, &pb.GetAccountRequest{Id: account.GetUserId()})
				Expect(err).Should(BeNil())
			})
		})
	})
})
