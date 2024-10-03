package user

import (
	model "bjm/db/benjamit/models"
	"bjm/src/v1/user/dto"
	"bjm/utils"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
		Birthday:  reqModel.Birthday.Time(),
		Email:     reqModel.Email,
	})

	if insertErr != nil {
		resModel.StatusCode = 500
		resModel.MessageDesc = insertErr.Error()
		return resModel
	}

	s.mapCreateUserResponseModel(insert, getPrefix, resModel)
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
	result := s._context.Where("uuid = ?", uuid).First(&prefix)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {

			return prefix, errors.New("uuid " + uuid.String() + " Prefix information not found.")
		}
		return prefix, result.Error
	}
	return prefix, nil
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
	}
	resModel.Data = data
}
