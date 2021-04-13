package VCSConnectionService

import (
	"alfred/modules/VCSConnectionService/v1/pb"
	"alfred/protos/types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"strings"
)

func (s *VCSConnectionService) Authorize(ctx context.Context, in *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	var redirectUrl string
	switch in.Provider {
	case types.GitProviders_GITHUB:
		conf := s.config.VCSConfig
		// Format required field in url
		mp := map[string]string{
			"scope":         conf["vcs_github"].Scope,
			"client_id":     conf["vcs_github"].ClientID,
			"redirect_uri":  conf["vcs_github"].RedirectUri,
			"response_type": conf["vcs_github"].ResponseType,
		}
		reg := regexp.MustCompile(`[{}]`)
		fields := reg.Split(conf["vcs_github"].URLTemplate, -1)
		for i := 1; i < len(fields); i = i + 2 {
			val, ok := mp[fields[i]]
			if ok {
				fields[i] = val
			} else {
				return nil, status.Error(codes.Internal, "Internal ERROR")
			}
		}
		redirectUrl = strings.Join(fields, "")
	}

	// todo - find account id to identify user account
	return &pb.AuthorizeResponse{
		RedirectUrl:  redirectUrl,
		TempJwtToken: "",
	}, nil
}
