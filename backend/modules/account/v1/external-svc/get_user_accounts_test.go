package external_svc_test

import (
	"alfred/modules/account/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
				_, err := accountServer.GetUserAccounts(ctx, &pb.GetUserAccountsRequest{UserId: "7820-kwiwi-3939"})
				Expect(err).Should(Equal(status.Error(codes.PermissionDenied, "unauthenticated user")))
			})
		})

		Context("Get a record when id is provided", func() {
			It("should return requested field in the object", func() {
				_, err := accountServer.GetUserAccounts(ctx, &pb.GetUserAccountsRequest{UserId: account.UserId})
				Expect(err).Should(BeNil())
			})
		})
	})
})
