package internal_svc_test

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

var _ = Describe("CreateAccount", func() {
	var (
		accountServer        pb.AccountInternalServer
		ctx                  context.Context
		userProfileIntServer usrPb.UserProfileInternalServer
		usrProfile           *usrPb.UserProfile
	)
	BeforeEach(func() {
		accountServer = AccountServerTest
		ctx = CtxTest
		userProfileIntServer = UserProfileIntServerTest

		id := "github_" + faker.Username()
		// Create a UserProfile before getting
		res, err := userProfileIntServer.CreateUserProfile(ctx, &usrPb.CreateUserProfileRequest{UserProfile: &usrPb.UserProfile{
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
		ctx = userinfo.NewContext(context.Background(), ui)
		usrProfile = res
	})

	Describe("Describe:Creating Account Correspond to user", func() {
		Context("Context:When user_id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.CreateAccount(context.Background(), &pb.CreateAccountRequest{Account: &pb.Account{
					UserId: "",
				}})
				Expect(err.(pb.CreateAccountRequestValidationError).Cause().(pb.AccountValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
		Context("Context:When request is nil", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.CreateAccount(context.Background(), &pb.CreateAccountRequest{Account: nil})
				Expect(err.(pb.CreateAccountRequestValidationError).Reason()).Should(Equal("value is required"))
			})
		})

		Context("Context:Happy Path", func() {
			It("It: should create the account", func() {
				_, err := accountServer.CreateAccount(ctx, &pb.CreateAccountRequest{Account: &pb.Account{
					Slug:   usrProfile.Sub + "_" + usrProfile.UserName,
					UserId: usrProfile.Id,
				}})
				Expect(err).Should(BeNil())
			})
		})
	})
})
