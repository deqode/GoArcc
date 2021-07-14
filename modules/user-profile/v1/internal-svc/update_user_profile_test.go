package internal_svc_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goarcc/modules/user-profile/v1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("UpdateUserProfile", func() {
	var (
		userProfileServer pb.UserProfileInternalServer
		//ctx               context.Context
		//profile        *pb.UserProfile
	)
	BeforeEach(func() {
		//profile = UsrProfile
		userProfileServer = UserProfileServerIntTest
		//ctx = CtxTest
	})

	Describe("Update a user profile", func() {
		//Negative Test Cases
		By("By a internal RPC Call")
		Context("Get an error when request object is nil", func() {
			It("should return nil exception", func() {
				_, err := userProfileServer.UpdateUserProfile(context.Background(), &pb.UpdateUserProfileRequest{UserProfile: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "UserProfile to update is not provided")))
			})
		})

		Context("Get an error when user profile object is nil", func() {
			It("should return nil exception", func() {
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
			})
		})
		Context("Get an error when user profile does not exist", func() {
			It("should return failed precondition error when profile not exist", func() {
			})
		})

		Context("Update user profile when update mask is correct and id is correct", func() {
			It("user profile should be updated successfully", func() {
			})
		})

		Context("Update user profile when update mask is correct and id is correct", func() {
			It("check the response of object that user profile is updated or not", func() {
			})
		})
	})
})
