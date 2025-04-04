// temperature_usecase.go
package application

import (
	"github.com/vicpoo/APILuz/Luz/domain"
	"github.com/vicpoo/APILuz/Luz/domain/entities"
)

type TemperatureUseCase struct {
	repo domain.TemperatureRepository
}

func NewTemperatureUseCase(repo domain.TemperatureRepository) *TemperatureUseCase {
	return &TemperatureUseCase{repo: repo}
}

func (uc *TemperatureUseCase) SaveTemperature(temp entities.Temperature) error {
	return uc.repo.Save(temp)
}

func (uc *TemperatureUseCase) GetAllTemperatures() ([]entities.Temperature, error) {
	return uc.repo.GetAll()
}
