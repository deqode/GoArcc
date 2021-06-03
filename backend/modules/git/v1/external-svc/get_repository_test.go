package external_svc_test

import (
	"alfred/modules/git/v1/pb"
	vcsPb "alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetRepository", func() {
	var (
		ctx       context.Context
		gitServer pb.GitsServer
		vcs       *vcsPb.VCSConnection
	)

	BeforeEach(func() {
		ctx = CtxTest
		gitServer = GitExtServerTest
		vcs = VCSTest
	})

	Context("Get an error when request is null", func() {
		It("Should get an error", func() {
			_, err := gitServer.GetRepository(ctx, nil)
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Get an error When Request has empty fields", func() {
		It("Should get an error if account_id is empty", func() {
			_, err := gitServer.GetRepository(ctx, &pb.GetRepositoryRequest{
				Provider:  vcs.Provider,
				RepoName:  "goonj",
				AccountId: "",
				OwnerName: "shivang-saxena",
			})
			Expect(err).ShouldNot(BeNil())
		})

		It("Should get an error if repository name is empty", func() {
			_, err := gitServer.GetRepository(ctx, &pb.GetRepositoryRequest{
				Provider:  vcs.Provider,
				RepoName:  "",
				AccountId: vcs.AccountId,
				OwnerName: "shivang-saxena",
			})
			Expect(err).ShouldNot(BeNil())
		})

		It("Should get an error if username is empty", func() {
			_, err := gitServer.GetRepository(ctx, &pb.GetRepositoryRequest{
				Provider:  vcs.Provider,
				RepoName:  "goonj",
				AccountId: vcs.AccountId,
				OwnerName: "",
			})
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Get an error When Request has invalid fields", func() {
		It("Should get an error if account_id is wrong", func() {
			_, err := gitServer.GetRepository(ctx, &pb.GetRepositoryRequest{
				Provider:  vcs.Provider,
				RepoName:  "goonj",
				AccountId: "190925cf-58d0-4451-ab55-712704398439",
				OwnerName: "shivang-saxena",
			})
			Expect(err).ShouldNot(BeNil())
		})
		It("Should get an error if provider is wrong", func() {
			_, err := gitServer.GetRepository(ctx, &pb.GetRepositoryRequest{
				Provider:  types.VCSProviders_UNKNOWN,
				RepoName:  "goonj",
				AccountId: "190925cf-58d0-4451-ab55-712704398439",
				OwnerName: "shivang-saxena",
			})
			Expect(err).ShouldNot(BeNil())
		})
	})
	Context("Get an error When Access Token is empty", func() {
		It("Should get an error", func() {
		})
	})

	//Positive Test Cases
	Context("Get valid response when request is valid", func() {
		It("Should get an error if user has no repository of desired name", func() {
		})
		It("Should get an error if repository has 0 branches", func() {
		})
	})

})
