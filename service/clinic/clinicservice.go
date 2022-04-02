package clinic

import (
	"BCC_Intership_BuildUlang/domain"

	"gorm.io/gorm"
)

type clinic struct {
	db *gorm.DB
}

func NewClinicRepository(db *gorm.DB) *clinic {
	return &clinic{db: db}
}

func (c *clinic) Create(newClinic *domain.Clinic) (*domain.Clinic, error) {
	if err := c.db.Create(&newClinic).Error; err != nil {
		return nil, err
	}
	return newClinic, nil
}

func (c *clinic) Update(oldClinic, newClinic *domain.Clinic) (*domain.Clinic, error) {
	if err := c.db.Model(&oldClinic).Updates(&newClinic).Error; err != nil {
		return nil, err
	}
	return newClinic, nil
}

func (c *clinic) Delete(deleteClinic *domain.Clinic) (*domain.Clinic, bool, error) {
	if err := c.db.Delete(&deleteClinic).Error; err != nil {
		return nil, false, err
	}
	return deleteClinic, true, nil
}

func (c *clinic) FindAll() ([]*domain.Clinic, error) {
	var result []*domain.Clinic
	if err := c.db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (c *clinic) FindById(id string) (*domain.Clinic, error) {
	var result *domain.Clinic
	if err := c.db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (c *clinic) FindByCity(city string) ([]*domain.Clinic, error) {
	var result []*domain.Clinic
	if err := c.db.Where("address = ?", city).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
