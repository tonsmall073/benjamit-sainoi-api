package prefix

import (
	"bjm/db/benjamit/models"
	"bjm/src/v1/prefix/dto"
	"bjm/utils"

	"gorm.io/gorm"
)

type PrefixService struct {
	_context *gorm.DB
}

func (s PrefixService) GetAllPrefix(resModel *dto.GetAllPrefixResponseModel) *dto.GetAllPrefixResponseModel {

	data, dataErr := s.fetchAllPrefixByActive(true)
	if dataErr != nil {
		resModel.StatusCode = 500
		resModel.MessageDesc = dataErr.Error()
		return resModel
	}
	s.mapGetAllPrefixResponseModel(data, resModel)
	resModel.StatusCode = 200
	resModel.MessageDesc = utils.HttpStatusCodes[200]
	return resModel
}

func (s PrefixService) fetchAllPrefixByActive(active bool) ([]models.Prefix, error) {
	prefixes := []models.Prefix{}
	if err := s._context.Where("active = ?", active).Find(&prefixes).Error; err != nil {
		return prefixes, err
	}
	return prefixes, nil
}

func (s PrefixService) mapGetAllPrefixResponseModel(data []models.Prefix, resModel *dto.GetAllPrefixResponseModel) {
	dataSub := []*dto.GetAllPrefixDataListResponseModel{}
	for _, item := range data {
		obj := &dto.GetAllPrefixDataListResponseModel{
			Uuid: item.UUID.String(),
			Name: item.Name,
		}
		dataSub = append(dataSub, obj)
	}
	resModel.Data = dataSub
}
