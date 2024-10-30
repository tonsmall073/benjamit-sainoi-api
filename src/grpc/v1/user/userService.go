package user

import (
	auth "bjm/auth/jwt"
	"bjm/db/benjamit/models"
	v1 "bjm/proto/v1"
	"bjm/utils"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	_context *gorm.DB
}

func (s *UserService) Login(
	reqModel *v1.LoginRequestModel,
	resModel *v1.LoginResponseModel,
) *v1.LoginResponseModel {
	if reqModel.Username == "" {
		resModel.MessageDesc = "username is empty or undefined"
		resModel.StatusCode = 400
		return resModel
	}
	if reqModel.Password == "" {
		resModel.MessageDesc = "password is empty or undefined"
		resModel.StatusCode = 400
		return resModel
	}

	data, dataErr := s.fetchUserByUsername(reqModel.Username)
	if dataErr != nil {
		resModel.MessageDesc = dataErr.Error()
		resModel.StatusCode = 400
		return resModel
	}
	resCheck := utils.CheckPasswordHash(reqModel.Password, data.Password)
	if !resCheck {
		resModel.MessageDesc = "the password is incorrect"
		resModel.StatusCode = 400
		return resModel
	}
	getToken, getTokenErr := auth.CreateToken(data.Username, data.UUID.String(), string(data.Role))
	if getTokenErr != nil {
		resModel.MessageDesc = getTokenErr.Error()
		resModel.StatusCode = 500
		return resModel
	}

	s.mapLoginResponseModel(data, getToken, resModel)
	resModel.StatusCode = 200
	resModel.MessageDesc = utils.HttpStatusCodes[200]
	return resModel
}

func (s *UserService) fetchUserByUsername(username string) (models.User, error) {
	user := models.User{}
	result := s._context.Preload("Prefix").Where("username = ? AND deleted_at IS NULL", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, errors.New("the username is incorrect")
		}
		return user, result.Error
	}
	return user, nil
}

func (s *UserService) mapLoginResponseModel(
	userData models.User,
	token string,
	resModel *v1.LoginResponseModel,
) {
	data := &v1.LoginDataListResponseModel{
		AccessToken: token,
		Uuid:        userData.UUID.String(),
		Username:    userData.Username,
		PrefixName:  userData.Prefix.Name,
		Nickname:    userData.Nickname,
		Firstname:   userData.Firstname,
		Lastname:    userData.Lastname,
		Birthday:    userData.Birthday.String(),
	}
	resModel.Data = data
}
