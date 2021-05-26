package internal_svc_test

import (
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Vcs Connection Test", func() {
	var (
		ctx              context.Context
		vcsIntServerTest pb.VCSConnectionInternalServer
		//account          *accPb.Account
		vcs *pb.VCSConnection
	)

	BeforeEach(func() {
		ctx = CtxTest
		vcsIntServerTest = VcsIntServerTest
		//account = Account
		vcs = VCSTest
	})
	//Negative Test Cases
	By("By a internal RPC Call")
	Context("Get an error when request object is nil", func() {
		It("should return nil exception", func() {
			_, err := vcsIntServerTest.CreateVCSConnection(ctx, &pb.CreateVCSConnectionRequest{VcsConnection: nil})
			Expect(err).ShouldNot(BeNil())
		})
	})
	Context("Get an error when vcs object is invalid", func() {
		It("should return an error when account id is empty", func() {
			_, err := vcsIntServerTest.CreateVCSConnection(ctx, &pb.CreateVCSConnectionRequest{
				VcsConnection: &pb.VCSConnection{
					Id:                 vcs.Id,
					Label:              vcs.Label,
					Provider:           vcs.Provider,
					ConnectionId:       vcs.ConnectionId,
					AccessToken:        vcs.AccessToken,
					RefreshToken:       vcs.RefreshToken,
					AccessTokenExpiry:  nil,
					RefreshTokenExpiry: nil,
					Revoked:            false,
					AccountId:          "",
					UserName:           vcs.UserName,
				},
			})
			Expect(err).ShouldNot(BeNil())
		})
	})
	Context("Get an error when vcs object is invalid", func() {
		It("should return error when username  is empty", func() {
			_, err := vcsIntServerTest.CreateVCSConnection(ctx, &pb.CreateVCSConnectionRequest{
				VcsConnection: &pb.VCSConnection{
					Id:                 vcs.Id,
					Label:              vcs.Label,
					Provider:           vcs.Provider,
					ConnectionId:       vcs.ConnectionId,
					AccessToken:        vcs.AccessToken,
					RefreshToken:       vcs.RefreshToken,
					AccessTokenExpiry:  nil,
					RefreshTokenExpiry: nil,
					Revoked:            false,
					AccountId:          vcs.AccountId,
					UserName:           "",
				},
			})
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Get an error when vcs object is invalid", func() {
		It("should return an error when account id is invalid", func() {
			_, err := vcsIntServerTest.CreateVCSConnection(ctx, &pb.CreateVCSConnectionRequest{
				VcsConnection: &pb.VCSConnection{
					Id:                 vcs.Id,
					Label:              vcs.Label,
					Provider:           vcs.Provider,
					ConnectionId:       vcs.ConnectionId,
					AccessToken:        vcs.AccessToken,
					RefreshToken:       vcs.RefreshToken,
					AccessTokenExpiry:  nil,
					RefreshTokenExpiry: nil,
					Revoked:            false,
					AccountId:          "84340-sjkj",
					UserName:           "",
				},
			})
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Create a Vcs when request object is valid", func() {
		It("Fail if response object is nil", func() {
			res, err := vcsIntServerTest.CreateVCSConnection(ctx, &pb.CreateVCSConnectionRequest{
				VcsConnection: &pb.VCSConnection{
					Id:                 vcs.Id,
					Label:              vcs.Label,
					Provider:           vcs.Provider,
					ConnectionId:       vcs.ConnectionId,
					AccessToken:        vcs.AccessToken,
					RefreshToken:       vcs.RefreshToken,
					AccessTokenExpiry:  vcs.AccessTokenExpiry,
					RefreshTokenExpiry: vcs.RefreshTokenExpiry,
					Revoked:            vcs.Revoked,
					AccountId:          vcs.AccountId,
					UserName:           vcs.UserName,
				},
			})
			Expect(err).Should(BeNil())
			Expect(res).ShouldNot(BeNil())
			//Expect(res).Should(Equal(vcs))
		})
	})
	Context("Create a Vcs when request object is valid", func() {
		It("Fail if username is empty", func() {
		})
	})
	Context("Create a Vcs when request object is valid", func() {
		It("Fail if account Id is empty", func() {
		})
	})

	Context("Create a Vcs when request object is valid", func() {
		It("Fail if access token is empty", func() {
		})
	})

	Context("Create a Vcs when request object is valid", func() {
		It("Fail if provider  is empty", func() {
		})
	})
	Context("Create a Vcs when request object is valid", func() {
		It("Fail if id   is empty", func() {
		})
	})

})
