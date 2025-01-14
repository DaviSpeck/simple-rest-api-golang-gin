package usecase

import "simple-rest-api-golang-gin/model"

type ProductUsecase struct {
	//Repository
}

func NewProductUseCase() ProductUsecase {
	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
