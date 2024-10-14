package services

import (
	"github.com/mereb/v4/database"
	"github.com/mereb/v4/models"
)

func CreatePersonService(person *models.Person) (*models.Person, error) {
	result := database.GetDB().Create(person)

	if result.Error != nil {
		return nil, result.Error

	}
	return person, nil
}

func GetAllPersonsService(perPage, page int) ([]models.Person, error) {
	var persons []models.Person
	result := database.GetDB().Limit(perPage).Offset(page).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func GetPersonService(id uint) (*models.Person, error) {
	var person models.Person
	result := database.GetDB().First(&person, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &person, nil
}

func UpdatePersonService(person *models.Person) error {
	result := database.GetDB().Save(person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeletePersonServices(id uint) error {
	result := database.GetDB().Delete(&models.Person{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
