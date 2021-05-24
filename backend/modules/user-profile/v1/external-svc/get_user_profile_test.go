package external_svc_test

import (
	"alfred/modules/user-profile/v1/external-svc"
	"alfred/modules/user-profile/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "log"
)

var _ = Describe("GetUserProfile", func() {
	var (
		userProfileServer pb.UserProfilesServer
		ctx               context.Context
	)

	// this block will run after each it block
	BeforeEach(func() {
		userProfileServer = external_svc.UserProfileServerTest
		ctx = external_svc.CtxTest
	})

	// this block will run after Before Suite
	AfterSuite(func() {})

	By("internal or external call")
	// positive
	Context("Get a record when id is provided", func() {
		It("should return requested field in the object", func() {
			_, err := userProfileServer.GetUserProfile(ctx, nil)
			Expect(err).Should(BeNil())
		})
	})
})
