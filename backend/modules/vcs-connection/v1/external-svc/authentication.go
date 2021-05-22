package external_svc

import (
	accountModel "alfred/modules/account/v1/models"
	userModel "alfred/modules/user-profile/v1/models"
	"alfred/util/userinfo"
	"context"
)

// ValidateUser : ValidateUser
func (s *vcsConnectionServer) ValidateUser(ctx context.Context) error {
	//Authentication check
	//it will check that context user id is present in user table or not
	//it will also check whether the user id is present in account table or not
	args := make(map[string]interface{})
	args["user_id"] = &accountModel.Account{}
	info := userinfo.ValidateUserInfo{
		Ctx:          ctx,
		RootTable:    &userModel.UserProfile{},
		RootTableTag: "id",
		Args:         args,
		Db:           s.db,
	}
	if err := info.ValidateUser(); err != nil {
		return err
	}
	return nil
}
