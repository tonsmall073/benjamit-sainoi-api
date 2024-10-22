package incomeAndExpense

import (
	"bjm/db/benjamit/models"
	"bjm/src/v1/incomeAndExpense/dto"
	"bjm/utils"
	"errors"

	"gorm.io/gorm"
)

type IncomeAndExpenseService struct {
	_context *gorm.DB
}

func (s *IncomeAndExpenseService) CreateList(
	reqModel *dto.CreateListRequestModel,
	resModel *dto.CreateListResponseModel,
	entrySource models.EntrySourceEnum,
	uuid string,
) *dto.CreateListResponseModel {
	resUser, resUserErr := s.fetchUserByUuid(uuid)
	if resUserErr != nil {
		resModel.MessageDesc = resUserErr.Error()
		resModel.StatusCode = 400
		return resModel
	}

	var productData = models.Product{}
	var productSellData = models.ProductSelling{}
	var insertData = models.IncomeAndExpense{
		Description:     reqModel.Description,
		TransactionType: reqModel.TransactionType,
		EntrySource:     entrySource,
		Amount:          s.adjustAmount(reqModel.Amount, reqModel.TransactionType),
		UserId:          int(resUser.ID),
	}

	if reqModel.ReferProductStatus {
		if reqModel.ReferProductUuid == "" {
			resModel.MessageDesc = `request body name "referProductStatus" is undefined or null`
			resModel.StatusCode = 400
			return resModel
		}
		if reqModel.ReferProductSellingUuid == "" {
			resModel.MessageDesc = `request body name "referProductSellingUuid" is undefined or null`
			resModel.StatusCode = 400
			return resModel
		}
		if reqModel.Quantity <= 0 {
			resModel.MessageDesc = `request body name "quantity" is less than equal to 0 or undefined or null`
			resModel.StatusCode = 400
			return resModel
		}
		resPro, resProErr := s.fetchProductByUuid(reqModel.ReferProductUuid)
		if resProErr != nil {
			resModel.MessageDesc = resProErr.Error()
			resModel.StatusCode = 400
			return resModel
		}

		resProSell, resProSellErr := s.fetchProductSellingByUuidAndProductId(reqModel.ReferProductSellingUuid, int(resPro.ID))
		if resProSellErr != nil {
			resModel.MessageDesc = resProSellErr.Error()
			resModel.StatusCode = 400
			return resModel
		}
		_, deductErr := s.productSellingDeductStockById(int(resProSell.ID), reqModel.Quantity, reqModel.TransactionType)
		if deductErr != nil {
			resModel.MessageDesc = deductErr.Error()
			resModel.StatusCode = 500
			return resModel
		}
		if reqModel.TransactionType == models.DEBIT {
			insertData.Amount = s.adjustAmount(float64(resProSell.SellPrice)*float64(reqModel.Quantity), reqModel.TransactionType)
		}
		if reqModel.TransactionType == models.CREDIT {
			insertData.Amount = s.adjustAmount(float64(resProSell.CostPrice)*float64(reqModel.Quantity), reqModel.TransactionType)
		}
		productData = resPro
		productSellData = resProSell
		insertData.ReferProductId = int(resPro.ID)
		insertData.Quantity = reqModel.Quantity
		insertData.ReferProductSellingId = int(resProSell.ID)
	}

	res, resErr := s.InsertIncomeAndExpense(insertData)
	if resErr != nil {
		resModel.MessageDesc = resErr.Error()
		resModel.StatusCode = 500
		return resModel
	}

	s.mapCreateListResponseModel(res, productData, productSellData, resModel)
	resModel.MessageDesc = utils.HttpStatusCodes[201]
	resModel.StatusCode = 201
	return resModel
}

func (s *IncomeAndExpenseService) InsertIncomeAndExpense(
	data models.IncomeAndExpense,
) (models.IncomeAndExpense, error) {
	if err := s._context.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (s *IncomeAndExpenseService) fetchProductByUuid(uuid string) (models.Product, error) {
	product := models.Product{}
	result := s._context.Where("uuid = ? AND deleted_at IS NULL", uuid).First(&product)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return product, errors.New("uuid " + uuid + " product information not found")
		}
		return product, result.Error
	}
	return product, nil
}

func (s *IncomeAndExpenseService) fetchProductSellingByUuidAndProductId(uuid string, productId int) (models.ProductSelling, error) {
	productSelling := models.ProductSelling{}
	result := s._context.Preload("UnitType").Where("uuid = ? AND product_id = ? AND deleted_at IS NULL", uuid, productId).First(&productSelling)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return productSelling, errors.New("uuid " + uuid + " product selling information not found or product information not found")
		}
		return productSelling, result.Error
	}
	return productSelling, nil
}

func (s *IncomeAndExpenseService) productSellingDeductStockById(id int, deductStock int, typeEnum models.TransactionTypeEnum) (models.ProductSelling, error) {
	var productSelling models.ProductSelling

	// ค้นหาสินค้า
	if err := s._context.First(&productSelling, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return productSelling, errors.New("product selling not found")
		}
		return productSelling, err
	}

	if typeEnum == models.DEBIT {
		// ตรวจสอบ stock ปัจจุบัน
		if productSelling.Stock < deductStock {
			return productSelling, errors.New("insufficient stock")
		}

		// หัก stock
		newStock := productSelling.Stock - deductStock
		productSelling.Stock = newStock

		// อัปเดต stock ในฐานข้อมูล
		if err := s._context.Save(&productSelling).Error; err != nil {
			return productSelling, err
		}

		return productSelling, nil
	} else if typeEnum == models.CREDIT {
		// บวก stock
		newStock := productSelling.Stock + deductStock
		productSelling.Stock = newStock

		// อัปเดต stock ในฐานข้อมูล
		if err := s._context.Save(&productSelling).Error; err != nil {
			return productSelling, err
		}
		return productSelling, nil
	} else {
		return productSelling, errors.New("transaction type not found")
	}
}

func (s *IncomeAndExpenseService) fetchUserByUuid(uuid string) (models.User, error) {
	user := models.User{}
	result := s._context.Where("uuid = ? AND deleted_at IS NULL", uuid).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, errors.New("user information not found")
		}
		return user, result.Error
	}
	return user, nil
}

func (s *IncomeAndExpenseService) adjustAmount(
	amount float64,
	typeEnum models.TransactionTypeEnum,
) float64 {
	if typeEnum == models.DEBIT {
		if amount < 0 {
			return -amount
		}
		return amount
	} else if typeEnum == models.CREDIT {
		if amount > 0 {
			return -amount
		}
		return amount
	}
	return 0
}

func (s *IncomeAndExpenseService) mapCreateListResponseModel(
	incomeData models.IncomeAndExpense,
	productData models.Product,
	productSellData models.ProductSelling,
	resModel *dto.CreateListResponseModel,
) {
	resModel.Data = &dto.CreateListDataListResponseModel{
		Amount:          incomeData.Amount,
		Description:     incomeData.Description,
		TransactionDate: incomeData.TransactionDate,
		ProductData: &dto.CreateListProductDataListResponseModel{
			Name: productData.Name,
		},
		ProductSellingData: &dto.CreateListProductSellingDataListResponseModel{
			SellPrice: productSellData.SellPrice,
			CostPrice: productSellData.CostPrice,
			UnitTypeData: &dto.CreateListUnitTypeDataListResponseModel{
				Name:   productSellData.UnitType.Name,
				NameEn: productSellData.UnitType.NameEn,
			},
		},
	}
}
