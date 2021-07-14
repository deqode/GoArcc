package external_svc_test

import (
	"context"
	"github.com/deqode/GoArcc/modules/authentication/v1/pb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("Login Callback", func() {
	var (
		authServer pb.AuthenticationsServer
		//authenticator *external_svc.Authenticator
	)
	BeforeEach(func() {
		authServer = AuthServerTest
		//authenticator = AuthenticatorTest
	})

	Describe("Login Callback Test", func() {
		By("By Rpc Calls")
		Context("Get an error when request object is nil", func() {
			It("Failed precondition err must be responded", func() {
				_, err := authServer.LoginCallback(context.Background(), nil)
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Request is Nil")))
			})
		})
		Context("Get an error if state or code is empty", func() {
			It("Failed precondition err must be responded", func() {
				_, err := authServer.LoginCallback(context.Background(), &pb.LoginCallbackRequest{Code: "", State: ""})
				Expect(err).ShouldNot(BeNil())
			})
		})

		Context("Get an error if user profile is not present", func() {
			It("gorm error must be thrown", func() {

			})
		})

		Context("Get CallBack Response when request object is valid", func() {
			It("Throw error if response object is nil", func() {
			})
			It("Throw error if user id is empty", func() {
			})
			It("Throw error if Access token is empty", func() {
			})
			It("Throw error if raw id  is empty", func() {
			})
		})
	})
})
