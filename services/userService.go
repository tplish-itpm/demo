package services

import (
	"github.com/tplish-itpm/demo/models"
	"github.com/tplish-itpm/demo/pkg/e"
	"github.com/tplish-itpm/demo/pkg/utils"
)

func Register(user *models.User) e.Error {
	u, err := models.FindUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if u != nil {
		return e.ErrNewCode(e.ErrUserExist)
	}

	user.Password, err = utils.Encrypt(user.Password)
	if err != nil {
		return err
	}

	err = models.AddUser(user)
	if err != nil {
		return e.ErrNewErr(err)
	}
	return nil
}

func Login(user *models.User) e.Error {
	u, err := models.FindUserByUsernameWithRole(user.Username)
	if err != nil {
		return err
	}
	if u == nil || utils.Compare(user.Password, u.Password) != nil {
		return e.ErrNewCode(e.ErrNamePwd)
	}
	*user = *u
	return nil
}
