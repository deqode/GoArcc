package external_svc_test

import (
	"alfred/modules/account/v1/pb"
	"alfred/util/userinfo"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("GetUserAccounts", func() {
	var (
		accountServer pb.AccountsServer
		ctx           context.Context
		//account       *pb.Account
	)
	BeforeEach(func() {
		accountServer = AccountServerTest
		ctx = CtxTest
		//account = Account
	})

	Describe("UserId Must be in Context", func() {
		Context("Get a record when id is provided by context", func() {
			It("should return error if context is wrong", func() {
				ui := userinfo.UserInfo{
					ID:          "Id",
					Email:       "Email",
					Sub:         "Sub",
					TokenExpiry: time.Time{},
				}
				newCtx := userinfo.NewContext(context.Background(), ui)
				_, err := accountServer.GetUserAccounts(newCtx, nil)
				Expect(err).ShouldNot(BeNil())
			})

			It("should return record if context is write", func() {
				res, err := accountServer.GetUserAccounts(ctx, nil)
				Expect(err).Should(BeNil())
				Expect(res).ShouldNot(BeNil())
			})
		})
	})
})
