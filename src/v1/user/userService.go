package user

import (
	model "bjm/db/benjamit/models"
	"bjm/src/v1/user/dto"
	"bjm/utils"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	auth "bjm/auth/jwt"
)

type UserService struct {
	_context *gorm.DB
}

func (s UserService) CreateUser(
	reqModel *dto.CreateUserRequestModel,
	resModel *dto.CreateUserResponseModel,
) *dto.CreateUserResponseModel {
	endcodePass, endCodePassErr := utils.HashPassword(reqModel.Password)
	if endCodePassErr != nil {
		resModel.StatusCode = 500
		resModel.MessageDesc = endCodePassErr.Error()
		return resModel
	}
	uuid, uuidErr := uuid.Parse(reqModel.PrefixUuid)
	if uuidErr != nil {
		resModel.StatusCode = 400
		resModel.MessageDesc = "Invalid UUID format " + reqModel.PrefixUuid
		return resModel
	}
	getPrefix, getPrefixErr := s.fetchPrefixByUuid(uuid)
	if getPrefixErr != nil {
		resModel.StatusCode = 400
		resModel.MessageDesc = getPrefixErr.Error()
		return resModel
	}

	insert, insertErr := s.insertUser(model.User{
		Username:  reqModel.Username,
		Password:  endcodePass,
		Nickname:  reqModel.Nickname,
		Firstname: reqModel.Firstname,
		Lastname:  reqModel.Lastname,
		PrefixId:  int(getPrefix.ID),
		Birthday:  reqModel.Birthday,
		Email:     reqModel.Email,
	})

	if insertErr != nil {
		resModel.StatusCode = 500
		resModel.MessageDesc = insertErr.Error()
		return resModel
	}

	s.mapCreateUserResponseModel(insert, getPrefix, resModel)
	resModel.StatusCode = 201
	resModel.MessageDesc = utils.HttpStatusCodes[201]

	return resModel
}

func (s UserService) Login(
	reqModel *dto.LoginRequestModel,
	resModel *dto.LoginResponseModel,
) *dto.LoginResponseModel {
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

func (s UserService) GetProfile(
	uuid string, resModel *dto.GetProfileResponseModel,
) *dto.GetProfileResponseModel {
	data, dataErr := s.fetchUserByUuid(uuid)
	if dataErr != nil {
		resModel.StatusCode = 400
		resModel.MessageDesc = dataErr.Error()
		return resModel
	}
	s.mapGetProfileResponseModel(data, resModel)
	resModel.StatusCode = 200
	resModel.MessageDesc = utils.HttpStatusCodes[200]
	return resModel
}

func (s UserService) insertUser(data model.User) (model.User, error) {
	if err := s._context.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (s UserService) fetchPrefixByUuid(uuid uuid.UUID) (model.Prefix, error) {
	prefix := model.Prefix{}
	result := s._context.Where("uuid = ? AND deleted_at IS NULL", uuid).First(&prefix)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return prefix, errors.New("uuid " + uuid.String() + " prefix information not found")
		}
		return prefix, result.Error
	}
	return prefix, nil
}

func (s UserService) fetchUserByUsername(username string) (model.User, error) {
	user := model.User{}
	result := s._context.Preload("Prefix").Where("username = ? AND deleted_at IS NULL", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, errors.New("the username is incorrect")
		}
		return user, result.Error
	}
	return user, nil
}

func (s UserService) fetchUserByUuid(uuid string) (model.User, error) {
	user := model.User{}
	result := s._context.Preload("Prefix").Where("uuid = ? AND deleted_at IS NULL", uuid).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, errors.New("user information not found")
		}
		return user, result.Error
	}
	return user, nil
}

func (s UserService) mapCreateUserResponseModel(
	userData model.User,
	prefixData model.Prefix,
	resModel *dto.CreateUserResponseModel,
) {
	data := &dto.CreateUserDataListResponseModel{
		Username:   userData.Username,
		PrefixName: prefixData.Name,
		Nickname:   userData.Nickname,
		Firstname:  userData.Firstname,
		Lastname:   userData.Lastname,
		Birthday:   userData.Birthday,
	}
	resModel.Data = data
}

func (s UserService) mapLoginResponseModel(
	userData model.User,
	token string,
	resModel *dto.LoginResponseModel,
) {
	data := &dto.LoginDataListResponseModel{
		AccessToken: token,
		Username:    userData.Username,
		PrefixName:  userData.Prefix.Name,
		Nickname:    userData.Nickname,
		Firstname:   userData.Firstname,
		Lastname:    userData.Lastname,
		Birthday:    userData.Birthday,
	}
	resModel.Data = data
}

func (s UserService) mapGetProfileResponseModel(
	userData model.User,
	resModel *dto.GetProfileResponseModel,
) {
	data := &dto.GetProfileDataListResponseModel{
		PrefixName:    userData.Prefix.Name,
		Firstname:     userData.Firstname,
		Lastname:      userData.Lastname,
		Nickname:      userData.Nickname,
		Birthday:      userData.Birthday,
		Email:         userData.Email,
		LineId:        userData.LineId,
		MobilePhoneNo: userData.HomePhoneNo,
		HomePhoneNo:   userData.HomePhoneNo,
	}
	resModel.Data = data
}
