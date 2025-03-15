package usecase

import "productManagement/internal/repository"

type UseCase struct {
	ProductUseCase ProductUseCase
	//CategoryUseCase      CategoryUseCase
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		ProductUseCase: NewProductUseCase(repo.ProductRepository),
		//CategoryUseCase:      NewCategoryUsecase(repo.CategoryRepository),
	}
}
