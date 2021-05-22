package external_svc

import (
	userModel "alfred/modules/user-profile/v1/models"
	"alfred/util/userinfo"
	"context"
)

func (s *userProfilesServer) ValidateUser(ctx context.Context) error {
	//Authentication check
	//it will check that context user id is present in user table or not
	info := userinfo.ValidateUserInfo{
		Ctx:                ctx,
		RootTable:          &userModel.UserProfile{},
		RootTableTag:       "id",
		Args:               nil,
		SkipRootValidation: false,
		SkipArgsValidation: true,
		Db:                 s.db,
	}
	if err := info.ValidateUser(); err != nil {
		return err
	}
	return nil
}
