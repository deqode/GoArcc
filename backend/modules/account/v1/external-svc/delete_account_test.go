package external_svc_test

import (
	"alfred/modules/account/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Describe:Delete account test", func() {
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

	Describe("Categorizing with invalid id", func() {
		Context("Context:When id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.DeleteAccount(context.Background(), &pb.DeleteAccountRequest{Id: ""})
				Expect(err.(pb.DeleteAccountRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})

		//Positive Test Cases
		Context("Return confirmation when record soft deleted in DB", func() {
			It("Err must be nil", func() {
				_, err := accountServer.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: account.Id})
				Expect(err).To(BeNil(), "Error")
			})
		})
	})
})
