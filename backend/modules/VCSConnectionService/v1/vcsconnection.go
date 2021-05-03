package VCSConnectionService

import (
	"alfred/modules/VCSConnectionService/v1/models"
	"alfred/modules/VCSConnectionService/v1/pb"
	"alfred/protos/types"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	GithubOAuthScope = []string{"user", "gist", "repo", "repo_deployment", "repo_hook"}
)

func (s *VCSConnectionService) ListAllSupportedVCSProviders(context.Context, *empty.Empty) (*pb.ListAllSupportedVCSProvidersResponse, error) {
	conf := s.config.VCSConfig
	var providers []string
	for _, vcs := range conf {
		providers = append(providers, vcs.Provider)
	}
	return &pb.ListAllSupportedVCSProvidersResponse{
		Providers: providers,
	}, nil
}
func (s *VCSConnectionService) ListVCSConnection(ctx context.Context, in *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {
	var record []models.VCSConnection
	chain := s.db.Where("account_id = ?", in.AccountId)
	if in.Provider != types.GitProviders_UNKNOWN {
		chain = chain.Where("provider = ?", in.Provider)
	}
	result := chain.Find(&record)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	var accountVcsConnection []*pb.AccountVCSConnection
	for _, v := range record {
		accountVcsConnection = append(accountVcsConnection, &pb.AccountVCSConnection{
			Id:        v.ID,
			Provider:  v.Provider,
			AccountId: v.AccountId,
			Label:     v.Label,
			UserName:  v.UserName,
		})
	}
	return &pb.ListVCSConnectionResponse{
		VcsConnection: accountVcsConnection,
	}, nil
}

func (s *VCSConnectionService) GetVCSConnection(ctx context.Context, in *pb.GetVCSConnectionRequest) (*pb.AccountVCSConnection, error) {
	var record models.VCSConnection
	chain := s.db.Where("account_id = ?", in.AccountId).Where("id = ?", in.Id)
	result := chain.Find(&record)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	accountVcsConnection := &pb.AccountVCSConnection{
		Id:        record.ID,
		Provider:  record.Provider,
		AccountId: record.AccountId,
		Label:     record.Label,
		UserName:  record.UserName,
	}
	return accountVcsConnection, nil
}
