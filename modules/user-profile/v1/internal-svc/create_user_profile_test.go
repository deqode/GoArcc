package internal_svc_test

import (
	"context"
	"github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"github.com/deqode/GoArcc/protos/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("CreateUserProfile", func() {
	var (
		userProfileServer pb.UserProfileInternalServer
		profile           *pb.UserProfile
	)

	// this block will run after each it block
	BeforeEach(func() {
		userProfileServer = UserProfileServerIntTest
		profile = UsrProfile
	})

	Describe("Creating an user profile", func() {
		//Negative Test Cases
		Context("Get an error when nil User provided", func() {
			It("should return nil exception", func() {
				_, err := userProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "UserProfile not provided")))
			})
		})

		Context("Create a user when subject is empty", func() {
			It("It should return validation error", func() {

			})
		})

		Context("Get an error when user created with wrong email", func() {
			It("if email is not provided then ignore because its a non required field", func() {})
			It("Should return validation error when email is provided but not formatted", func() {})
		})

		Context("Get an error when user created with wrong Phone", func() {
			It("if phone is not provided then ignore because its a non required field", func() {})
			It("Should return validation error when phone number is not in proper", func() {})
		})

		Context("Get an error when name of user not provided", func() {
			It("if name is nil", func() {})
			It("if name length exceed maximum upto 100 character", func() {})
		})

		Context("Get an error when user-name of user not provided", func() {
			It("if user-name is empty return an error", func() {})
			It("if user-name length exceed maximum upto 100 character return error", func() {})
		})

		Context("Return proper error when user is creating from unknown source", func() {
			It("should return error if user is from unknown source", func() {
				_, err := userProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: &pb.UserProfile{
					Id:             profile.Id,
					Sub:            profile.Sub,
					Name:           profile.Name,
					UserName:       profile.UserName,
					Email:          profile.Email,
					PhoneNumber:    profile.PhoneNumber,
					ExternalSource: types.VCSProviders_UNKNOWN,
					ProfilePicUrl:  profile.ProfilePicUrl,
					TokenValidTill: profile.TokenValidTill,
				}})
				Expect(err.(pb.CreateUserProfileRequestValidationError).Cause().(pb.UserProfileValidationError).Reason()).Should(Equal("value must not be in list [0]"))
			})
		})

		Context("Get an error when wrong profile-address provided", func() {
			It("if not a valid url return error", func() {})

		})

		//Positive Test Cases
		Context("Create a user profile", func() {
			It("should return user_id as uuid", func() {
				_, err := userProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: &pb.UserProfile{
					Id:             profile.Id + "test",
					Sub:            profile.Sub + "test",
					Name:           profile.Name,
					UserName:       profile.UserName,
					Email:          profile.Email,
					PhoneNumber:    profile.PhoneNumber,
					ExternalSource: profile.ExternalSource,
					ProfilePicUrl:  profile.ProfilePicUrl,
					TokenValidTill: profile.TokenValidTill,
				}})
				Expect(err).To(BeNil(), "Error")
			})
		})

		Context("Get an error when user create with already existing sub", func() {
			It("should return already exists error", func() {
				request := profile
				_, err := userProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: request})
				Expect(err).To(Not(BeNil()), "Error")
			})
		})

	})
})
