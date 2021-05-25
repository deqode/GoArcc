package external_svc_test

import (
	"alfred/modules/authentication/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login", func() {
	var (
		authServer pb.AuthenticationsServer
		//authenticator *external_svc.Authenticator
	)
	BeforeEach(func() {
		authServer = AuthServerTest
		//authenticator = AuthenticatorTest
	})
	Describe("Getting External VCS login url", func() {
		By("By Rpc Calls")

		// positive test case
		Context("Get an error if returned url is empty", func() {
			It("Failed precondition err must be responded", func() {
				res, err := authServer.Login(context.Background(), nil)
				Expect(err).Should(BeNil())
				Expect(res).ShouldNot(BeNil())
			})
		})
	})
})
