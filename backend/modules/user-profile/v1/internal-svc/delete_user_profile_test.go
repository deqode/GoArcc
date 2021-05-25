package internal_svc_test

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("DeleteUserProfile", func() {
	var (
		userProfileServer pb.UserProfileInternalServer
		ctx               context.Context
		profile           *pb.UserProfile
	)
	BeforeEach(func() {
		profile = UsrProfile
		userProfileServer = UserProfileServerIntTest
		ctx = CtxTest
	})

	Describe("Deleting a user profile", func() {
		//Negative Cases
		By("By an internal RPC Call")
		Context("Get an error when no user_id provided", func() {
			It("Should return validation error", func() {
				_, err := userProfileServer.DeleteUserProfile(ctx, &pb.DeleteUserProfileRequest{Id: ""})
				Expect(err.(pb.DeleteUserProfileRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})

		Context("Get an error when request is nil", func() {
			It("should return error ", func() {
				_, err := userProfileServer.DeleteUserProfile(ctx, nil)
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Request can't be nil")))
			})
		})

		Context("Return nil when wrong user_id provided", func() {
			It("should not perform any action in DB and return nil", func() {
				_, err := userProfileServer.DeleteUserProfile(ctx, &pb.DeleteUserProfileRequest{Id: "fejfoeiw0r8290r420"})
				Expect(err).To(BeNil(), "Error")
			})
		})

		//Positive Test Cases
		Context("Return confirmation when record soft deleted in DB", func() {
			It("should return a boolean value", func() {
				_, err := userProfileServer.DeleteUserProfile(ctx, &pb.DeleteUserProfileRequest{Id: profile.Id})
				Expect(err).To(BeNil(), "Error")
			})
		})
	})

})
