package external_svc_test

import (
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List All Supported VCS Providers", func() {
	var (
		ctx          context.Context
		vcsExtServer pb.VCSConnectionsServer
	)

	BeforeEach(func() {
		ctx = CtxTest
		vcsExtServer = VcsServerTest
	})

	By("By a RPC Call")
	Context("Get List of all VCS Providers", func() {
		It("should return nil exception", func() {
			res, err := vcsExtServer.ListAllSupportedVCSProviders(ctx, nil)
			if err != nil {
				Expect(err).ShouldNot(BeNil())
			} else {
				Expect(len(res.Providers)).ShouldNot(BeZero())
			}
		})
	})

})
