package external_svc_test

import (
	"alfred/modules/account/v1/pb"
	usrPb "alfred/modules/user-profile/v1/pb"
	"alfred/protos/types"
	"alfred/util/userinfo"
	"context"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Delete A Account", func() {
	var (
		accountServer      pb.AccountsServer
		ctx                context.Context
		accountInServer    pb.AccountInternalServer
		usrProfileInServer usrPb.UserProfileInternalServer
		account            *pb.Account
	)

	// This Before Each run for outermost Describe
	BeforeEach(func() {
		accountServer = AccountServerTest
		accountInServer = AccountServerInTest
		usrProfileInServer = UsrProfileInServerTest
	})

	Describe("Get an error if wrong id is provided", func() {
		Context("Context:When id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.DeleteAccount(CtxTest, &pb.DeleteAccountRequest{Id: ""})
				Expect(err.(pb.DeleteAccountRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
	})

	Describe("Get a non nil error if valid id is provided", func() {
		// BeforeEach will run after each It Block in this Describe
		BeforeEach(func() {
			id := "github_" + faker.Username()
			// Create a UserProfile before getting
			res, err := usrProfileInServer.CreateUserProfile(ctx, &usrPb.CreateUserProfileRequest{UserProfile: &usrPb.UserProfile{
				Id:             id,
				Sub:            id,
				Name:           faker.Name(),
				UserName:       faker.Username(),
				Email:          faker.Email(),
				PhoneNumber:    faker.Phonenumber(),
				ExternalSource: types.VCSProviders_GITHUB,
				ProfilePicUrl:  faker.URL(),
				TokenValidTill: nil,
			}})
			if err != nil {
				return
			}
			ui := userinfo.UserInfo{
				ID:          res.Id,
				Email:       res.Email,
				Sub:         res.Sub,
				TokenExpiry: time.Time{},
			}
			ctx = context.Background()
			ctx = userinfo.NewContext(ctx, ui)
			// Create A Account
			acc, err := accountInServer.CreateAccount(ctx, &pb.CreateAccountRequest{Account: &pb.Account{
				Slug:   res.Sub + "_" + res.UserName,
				UserId: res.Id,
			}})
			if err != nil {
				return
			}
			account = acc
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
