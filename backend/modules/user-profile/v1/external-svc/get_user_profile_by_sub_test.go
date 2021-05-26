package external_svc_test

import (
	pb "alfred/modules/user-profile/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("GetUserProfile By unique External Sub", func() {
	var (
		userProfileServer pb.UserProfilesServer
		ctx               context.Context
		usrProfile        *pb.UserProfile
	)

	// this block will run after each it block
	BeforeEach(func() {
		userProfileServer = UserProfileServerTest
		ctx = CtxTest
		usrProfile = UsrProfile
	})

	Describe("Get a user-profile", func() {
		By("internal or external call")
		Context("Get an error when sub is empty", func() {
			It("it should return validation error", func() {
				_, err := userProfileServer.GetUserProfileBySub(ctx, &pb.GetUserProfileBySubRequest{Sub: ""})
				Expect(err.(pb.GetUserProfileBySubRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
		Context("Get an error when sub_id is wrong", func() {
			It("should return not found error", func() {
				_, err := userProfileServer.GetUserProfileBySub(ctx, &pb.GetUserProfileBySubRequest{Sub: "fji5895"})
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})
		})

		Context("Get a record when sub_id is provided", func() {
			It("should return requested field in the object", func() {
				usr, err := userProfileServer.GetUserProfileBySub(ctx, &pb.GetUserProfileBySubRequest{Sub: usrProfile.Sub})
				Expect(err).Should(BeNil())
				Expect(usr.Sub).Should(Equal(usrProfile.Sub))
			})
		})
	})
})
