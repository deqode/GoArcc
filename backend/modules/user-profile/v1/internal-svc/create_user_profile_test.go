package internal_svc_test

import (
	"alfred/modules/user-profile/v1/pb"
	"alfred/protos/types"
	"context"
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

	// this block will run after Before Suite
	AfterSuite(func() {})

	Describe("Creating an user profile", func() {

		//Negative Test Cases
		By("By a internal RPC Call")
		Context("Get an error when nil User provided", func() {
			It("should return nil exception", func() {
				_, err := userProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "UserProfile not provided")))
			})
		})

		Context("Create a user when subject is empty", func() {
			It("It should return validation error", func() {
				//request := Profile
				//request.Sub = ""
				//_, err := UserProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: request})
				//Expect(err.(pb.CreateUserProfileRequestValidationError).Cause().(pb.UserProfileValidationError).Field()).Should(Equal(pb.UserProfileValidationError{}))
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
				request := profile
				request.ExternalSource = types.VCSProviders_UNKNOWN
				_, err := userProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: request})
				Expect(err.(pb.CreateUserProfileRequestValidationError).Cause().(pb.UserProfileValidationError).Reason()).Should(Equal("value must not be in list [0]"))
			})
		})

		Context("Get an error when wrong profile-address provided", func() {
			It("if not a valid url return error", func() {})

		})

		//Positive Test Cases
		Context("Create a user profile", func() {
			It("should return user_id as uuid", func() {
				//request := profile
				//user, err := userProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: request})
				//Expect(err).To(BeNil(), "Error")
				//Expect(user).To(Not(BeNil()), "User Object Must present")
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
