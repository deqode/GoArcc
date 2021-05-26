package external_svc_test

import (
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Vcs Connection Test", func() {
	var (
		ctx          context.Context
		vcsExtServer pb.VCSConnectionsServer
		vcs          *pb.VCSConnection
	)

	BeforeEach(func() {
		ctx = CtxTest
		vcsExtServer = VcsServerTest
		vcs = VCSTest
	})
	//Negative Test Cases
	By("By a internal RPC Call")
	Context("Get an error when request object is nil", func() {
		It("should return nil exception", func() {
			_, err := vcsExtServer.GetAccountVCSConnection(ctx, nil)
			Expect(err).ShouldNot(BeNil())
		})
	})
	Context("Get an error when account id is empty ", func() {
		It("should return nil exception", func() {
			_, err := vcsExtServer.GetAccountVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: "",
				Provider:  vcs.Provider,
			})
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Get an error when id is invalid or account id is invalid", func() {
		It("should return an no record found error ", func() {
			_, err := vcsExtServer.GetAccountVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        "489-shjc-si",
				AccountId: "ruier-439-dfd",
				Provider:  vcs.Provider,
			})
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("When Request object is valid", func() {
		It("Fail in case id is not present in response", func() {
			res, err := vcsExtServer.GetAccountVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: vcs.AccountId,
				Provider:  vcs.Provider,
			})
			Expect(err).Should(BeNil())
			Expect(res.UserName).ShouldNot(BeEmpty())
			Expect(res.AccountId).ShouldNot(BeEmpty())
			Expect(res.Id).ShouldNot(BeEmpty())
		})
	})
	Context("When Request object is valid", func() {
		It("Fail in case provider is not present in response", func() {
			res, err := vcsExtServer.GetAccountVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: vcs.AccountId,
				Provider:  vcs.Provider,
			})
			Expect(err).Should(BeNil())
			Expect(res.UserName).ShouldNot(BeEmpty())
			Expect(res.AccountId).ShouldNot(BeEmpty())
			Expect(res.Provider).ShouldNot(BeZero())
		})
	})

	Context("When Request object is valid", func() {
		It("Fail in case username is not present in response", func() {
			res, err := vcsExtServer.GetAccountVCSConnection(ctx, &pb.GetVCSConnectionRequest{
				Id:        vcs.Id,
				AccountId: vcs.AccountId,
				Provider:  vcs.Provider,
			})
			Expect(err).Should(BeNil())
			Expect(res.UserName).ShouldNot(BeEmpty())
		})
	})
})
