package external_svc_test

import (
	"alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authorize Test", func() {
	var (
		ctx          context.Context
		vcsExtServer pb.VCSConnectionsServer
	)

	BeforeEach(func() {
		ctx = CtxTest
		vcsExtServer = VcsServerTest
	})
	Context("Get an error when request object is empty", func() {
		It("Failed precondition err must be responded", func() {
			_, err := vcsExtServer.Authorize(ctx, nil)
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Get an error when provider object is invalid", func() {
		It("validation error must be return when", func() {
			res, err := vcsExtServer.Authorize(ctx, &pb.AuthorizeRequest{
				Provider: types.VCSProviders_UNKNOWN,
			})
			Expect(err).Should(BeNil())
			Expect(res.RedirectUrl).Should(BeEmpty())
		})
	})

	Context("Get an error when vcs config is not present", func() {
		Context("Successfully Redirect Url must be returned when request object is valid", func() {
			It("Return error if redirect url is empty", func() {
				res, err := vcsExtServer.Authorize(ctx, &pb.AuthorizeRequest{
					Provider: types.VCSProviders_GITHUB,
				})
				Expect(err).Should(BeNil())
				Expect(res.RedirectUrl).ShouldNot(BeEmpty())
			})
			It("Return error if response is nil", func() {
				res, err := vcsExtServer.Authorize(ctx, &pb.AuthorizeRequest{
					Provider: types.VCSProviders_GITHUB,
				})
				Expect(err).Should(BeNil())
				Expect(res).ShouldNot(BeNil())
			})
		})
	})
})
