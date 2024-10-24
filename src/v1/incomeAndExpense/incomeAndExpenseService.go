package incomeAndExpense

import (
	"bjm/db/benjamit/models"
	"bjm/src/v1/incomeAndExpense/dto"
	"bjm/utils"
	"bjm/utils/enums"
	"errors"
	"time"

	"gorm.io/gorm"
)

type IncomeAndExpenseService struct {
	_context *gorm.DB
}

func (s *IncomeAndExpenseService) CreateList(
	reqModel *dto.CreateListRequestModel,
	resModel *dto.CreateListResponseModel,
	entrySource enums.EntrySourceEnum,
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

		s._context = s._context.Begin()

		_, deductErr := s.updateProductSellingDeductStockById(int(resProSell.ID), reqModel.Quantity, reqModel.TransactionType)
		if deductErr != nil {
			s._context.Rollback()
			resModel.MessageDesc = deductErr.Error()
			resModel.StatusCode = 500
			return resModel
		}
		if reqModel.TransactionType == enums.DEBIT {
			insertData.Amount = s.adjustAmount(float64(resProSell.SellPrice)*float64(reqModel.Quantity), reqModel.TransactionType)
		}
		if reqModel.TransactionType == enums.CREDIT {
			insertData.Amount = s.adjustAmount(float64(resProSell.CostPrice)*float64(reqModel.Quantity), reqModel.TransactionType)
		}
		productData = resPro
		productSellData = resProSell
		insertData.ReferProductId = int(resPro.ID)
		insertData.Quantity = reqModel.Quantity
		insertData.ReferProductSellingId = int(resProSell.ID)
	}

	res, resErr := s.insertIncomeAndExpense(insertData)
	if resErr != nil {
		s._context.Rollback()
		resModel.MessageDesc = resErr.Error()
		resModel.StatusCode = 500
		return resModel
	}

	s._context.Commit()
	s.mapCreateListResponseModel(res, productData, productSellData, resModel)
	resModel.MessageDesc = utils.HttpStatusCodes[201]
	resModel.StatusCode = 201
	return resModel
}

func (s *IncomeAndExpenseService) GetAllList(
	reqModel *dto.GetAllListRequestModel,
	resModel *dto.GetAllListResponseModel,
) *dto.GetAllListResponseModel {
	getAll, totalData, getAllErr := s.fetchAllIncomeAndExpenseBySearch(
		reqModel.Search,
		reqModel.Sort,
		reqModel.SortColumn,
		reqModel.StartDate,
		reqModel.EndDate,
		reqModel.Skip,
		reqModel.Take,
	)
	if getAllErr != nil {
		resModel.MessageDesc = getAllErr.Error()
		resModel.StatusCode = 500
		return resModel
	}

	s.mapGetAllListResponseModel(getAll, resModel)

	resModel.TotalData = totalData
	resModel.MessageDesc = utils.HttpStatusCodes[200]
	resModel.StatusCode = 200
	return resModel
}

func (s *IncomeAndExpenseService) insertIncomeAndExpense(
	data models.IncomeAndExpense,
) (models.IncomeAndExpense, error) {
	if err := s._context.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (s *IncomeAndExpenseService) fetchAllIncomeAndExpenseBySearch(
	search string,
	sort enums.SortEnum,
	sortColumn string,
	startDate time.Time,
	endDate time.Time,
	skip int,
	take int,
) ([]models.IncomeAndExpense, int, error) {
	var totalCount int64
	data := []models.IncomeAndExpense{}
	query := s._context.Preload("Product").
		Preload("ProductSelling").
		Preload("ProductSelling.UnitType").
		Where("income_and_expenses.deleted_at IS NULL").
		Joins("LEFT JOIN products ON income_and_expenses.refer_product_id = products.id").
		Joins("LEFT JOIN product_sellings ON income_and_expenses.refer_product_selling_id = product_sellings.id").
		Joins("LEFT JOIN unit_types ON product_sellings.unit_type_id = unit_types.id")

	if search != "" {
		query = query.Where(
			"income_and_expenses.description LIKE ? OR "+
				"products.name LIKE ? OR "+
				"CAST(income_and_expenses.amount AS TEXT) LIKE ? OR "+
				"CAST(product_sellings.sell_price AS TEXT) LIKE ? OR "+
				"CAST(product_sellings.cost_price AS TEXT) LIKE ? OR "+
				"unit_types.name LIKE ? OR "+
				"unit_types.name_en LIKE ? OR "+
				"CAST(income_and_expenses.quantity AS TEXT) LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%",
		)
	}
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("transaction_date BETWEEN ? AND ?", startDate, endDate)
	}

	if sort != "" && sortColumn != "" {
		query = query.Order(sortColumn + " " + string(sort))
	}

	resultCount := query.Model(&models.IncomeAndExpense{}).Count(&totalCount)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}

	if take >= 0 && skip >= 0 {
		query = query.Offset(skip).Limit(take)
	}

	result := query.Find(&data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return data, int(totalCount), errors.New("income and expense no data")
		}
		return data, int(totalCount), result.Error
	}

	return data, int(totalCount), nil
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

func (s *IncomeAndExpenseService) updateProductSellingDeductStockById(id int, qtyStock int, typeEnum enums.TransactionTypeEnum) (models.ProductSelling, error) {
	var productSelling models.ProductSelling

	// ค้นหาสินค้า
	if err := s._context.First(&productSelling, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return productSelling, errors.New("product selling not found")
		}
		return productSelling, err
	}

	if typeEnum == enums.DEBIT {
		// ตรวจสอบ stock ปัจจุบัน
		if productSelling.Stock < qtyStock {
			return productSelling, errors.New("insufficient stock")
		}

		// หัก stock
		newStock := productSelling.Stock - qtyStock
		productSelling.Stock = newStock

		// อัปเดต stock ในฐานข้อมูล
		if err := s._context.Save(&productSelling).Error; err != nil {
			return productSelling, err
		}

		return productSelling, nil
	} else if typeEnum == enums.CREDIT {
		// บวก stock
		newStock := productSelling.Stock + qtyStock
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
	typeEnum enums.TransactionTypeEnum,
) float64 {
	if typeEnum == enums.DEBIT {
		if amount < 0 {
			return -amount
		}
		return amount
	} else if typeEnum == enums.CREDIT {
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

func (s *IncomeAndExpenseService) mapGetAllListResponseModel(
	incomeData []models.IncomeAndExpense,
	resModel *dto.GetAllListResponseModel,
) {
	resModel.Data = make([]*dto.GetAllListDataListResponseModel, 0)

	for _, data := range incomeData {
		list := &dto.GetAllListDataListResponseModel{
			Amount:          data.Amount,
			Description:     data.Description,
			TransactionDate: data.TransactionDate,
			ProductData: &dto.GetAllListProductDataListResponseModel{
				Name: data.Product.Name,
			},
			ProductSellingData: &dto.GetAllListProductSellingDataListResponseModel{
				SellPrice: data.ProductSelling.SellPrice,
				CostPrice: data.ProductSelling.CostPrice,
				UnitTypeData: &dto.GetAllListUnitTypeDataListResponseModel{
					Name:   data.ProductSelling.UnitType.Name,
					NameEn: data.ProductSelling.UnitType.NameEn,
				},
			},
		}
		resModel.Data = append(resModel.Data, list)
	}

}
