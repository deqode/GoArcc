package external_svc

import (
	model "alfred/modules/account/v1/models"
	userModel "alfred/modules/user-profile/v1/models"
	"alfred/util/userinfo"
	"context"
)

func (s *accountExtServer) ValidateUser(ctx context.Context) error {
	//Authentication check
	//it will check that context user id is present in user table or not
	//it will also check whether the context user id is present in account table or not
	args := make(map[string]interface{})
	args["user_id"] = &model.Account{}
	info := userinfo.ValidateUserInfo{
		Ctx:                ctx,
		RootTable:          &userModel.UserProfile{},
		RootTableTag:       "id",
		Args:               args,
		SkipRootValidation: false,
		SkipArgsValidation: false,
		Db:                 s.db,
	}
	if err := info.ValidateUser(); err != nil {
		return err
	}
	return nil
}
