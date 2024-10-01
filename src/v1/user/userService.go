package user

import (
	model "bjm/db/benjamit/models"
	"bjm/src/v1/user/dto"
	"bjm/utils"

	"gorm.io/gorm"
)

type UserService struct {
	_context *gorm.DB
}

func (s *UserService) CreateUser(
	reqModel *dto.CreateUserRequestModel,
	resModel *dto.CreateUserResponseModel,
) *dto.CreateUserResponseModel {
	insert, insertErr := s.insertUser(&model.User{
		Lastname:  reqModel.Lastname,
		Firstname: reqModel.Firstname,
		PrefixId:  0,
		Birthday:  reqModel.Birthday.Time(),
	})

	if insertErr != nil {
		resModel.Status = 500
		resModel.MessageDesc = insertErr.Error()
	}
	s.mapCreateUserResponse(insert, resModel)
	resModel.Status = 200
	resModel.MessageDesc = utils.HttpStatusCodes[200]

	return resModel
}

func (s *UserService) insertUser(data *model.User) (*model.User, error) {
	if err := s._context.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *UserService) mapCreateUserResponse(data *model.User, resModel *dto.CreateUserResponseModel) {
	resModel.Data.Username = data.Username
	resModel.Data.Firstname = data.Firstname
	resModel.Data.Lastname = data.Lastname
}
