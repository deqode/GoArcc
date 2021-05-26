package internal_svc_test

import (
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Vcs Connection Test", func() {
	var (
		ctx          context.Context
		vcsIntServer pb.VCSConnectionInternalServer
		vcs          *pb.VCSConnection
	)

	BeforeEach(func() {
		ctx = CtxTest
		vcsIntServer = VcsIntServerTest
		vcs = VCSTest
	})

	//Negative Test Cases
	By("By a internal RPC Call")
	Context("Get an error when request object is nil", func() {
		It("should return nil exception", func() {
			_, err := vcsIntServer.GetVCSConnection(ctx, nil)
			Expect(err).ShouldNot(BeNil())
		})
	})
	Context("Get an error when account id is empty ", func() {
		It("should return nil exception", func() {
			_, err := vcsIntServer.GetVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: "",
				Provider:  vcs.Provider,
			})
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Get an error when id is invalid or account id is invalid", func() {
		It("should return an no record found error ", func() {
			_, err := vcsIntServer.GetVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: "89382-dskns-2i2",
				Provider:  vcs.Provider,
			})
			Expect(err).ShouldNot(BeNil())
			//Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Context("When Request object is valid", func() {
		It("Fail in case access token is not present in response", func() {
			res, err := vcsIntServer.GetVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: vcs.AccountId,
				Provider:  vcs.Provider,
			})
			Expect(err).Should(BeNil())
			Expect(res.AccessToken).ShouldNot(BeEmpty())
		})
	})

	Context("When Request object is valid", func() {
		It("Fail in case user_name is not present in response", func() {
			res, err := vcsIntServer.GetVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: vcs.AccountId,
				Provider:  vcs.Provider,
			})
			Expect(err).Should(BeNil())
			Expect(res.UserName).ShouldNot(BeEmpty())
		})
	})

	Context("When Request object is valid", func() {
		It("Fail in case accountId is not present in response", func() {
			res, err := vcsIntServer.GetVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: vcs.AccountId,
				Provider:  vcs.Provider,
			})
			Expect(err).Should(BeNil())
			Expect(res.AccountId).ShouldNot(BeEmpty())
		})
	})
})
