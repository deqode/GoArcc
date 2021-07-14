package external_svc_test

import (
	"context"
	"github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"github.com/deqode/GoArcc/util/userinfo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	_ "log"
	"time"
)

var _ = Describe("GetUserProfile", func() {
	var (
		userProfileServer pb.UserProfilesServer
		ctx               context.Context
	)

	// this block will run after each it block
	BeforeEach(func() {
		userProfileServer = UserProfileServerTest
		ctx = CtxTest
	})

	// negative
	By("internal or external call")
	Context("Get an error when id is wrong", func() {
		It("Return record not found error", func() {
			ui := userinfo.UserInfo{
				ID:          "Id",
				Email:       "Email",
				Sub:         "Sub",
				TokenExpiry: time.Time{},
			}
			newCtx := userinfo.NewContext(context.Background(), ui)
			_, err := userProfileServer.GetUserProfile(newCtx, nil)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})
	// positive
	Context("Get a record when id is provided", func() {
		It("should return requested field in the object", func() {
			_, err := userProfileServer.GetUserProfile(ctx, nil)
			Expect(err).Should(BeNil())
		})
	})
})
