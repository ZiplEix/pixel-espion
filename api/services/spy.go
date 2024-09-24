package services

import (
	"fmt"
	"time"

	"github.com/ZiplEix/pixel-espion/database"
	"github.com/ZiplEix/pixel-espion/models"
	requestmodels "github.com/ZiplEix/pixel-espion/request_models"
	"github.com/sanity-io/litter"
)

func Pixel1(spyId string, clientIp string) (string, error) {
	var spy models.Spy
	if err := database.Db.First(&spy, spyId).Error; err != nil {
		return "", ServiceError{
			Code:    404,
			Message: "Spy not found: " + err.Error(),
		}
	}

	record := models.Record{
		Ip:    clientIp,
		SpyID: spy.ID,
		Time:  time.Now(),
	}

	if err := database.Db.Create(&record).Error; err != nil {
		return "", ServiceError{
			Code:    500,
			Message: "Error while creating record: " + err.Error(),
		}
	}

	fmt.Printf("Spy '%s' has been visited by '%s'\n", spy.Name, clientIp)
	litter.Dump(record)

	imagePath := "pixel/spy.png"

	return imagePath, nil
}

func NewSpy(req requestmodels.NewSpyRequest, userId uint) (uint, error) {
	spy := models.Spy{
		Name:   req.Name,
		Color:  req.Color,
		UserId: userId,
	}

	if err := database.Db.Create(&spy).Error; err != nil {
		return 0, ServiceError{
			Code:    500,
			Message: "Error while creating spy: " + err.Error(),
		}
	}

	return spy.ID, nil
}

func GetAllSpies(userId uint) ([]models.Spy, error) {
	var spies []models.Spy

	if err := database.Db.Where("user_id = ?", userId).Find(&spies).Error; err != nil {
		return nil, ServiceError{
			Code:    500,
			Message: "Error while fetching spies: " + err.Error(),
		}
	}

	return spies, nil
}

func GetSpy(spyId string) (*models.Spy, error) {
	var spy models.Spy

	if err := database.Db.First(&spy, "id = ?", spyId).Error; err != nil {
		return nil, ServiceError{
			Code:    404,
			Message: "Spy not found: " + err.Error(),
		}
	}

	return &spy, nil
}

func GetSpyRecords(spyId string) ([]models.Record, error) {
	var records []models.Record

	if err := database.Db.Where("spy_id = ?", spyId).Find(&records).Error; err != nil {
		return nil, ServiceError{
			Code:    500,
			Message: "Error while retrieving records: " + err.Error(),
		}
	}

	return records, nil
}

func GetAllRecords(userId uint) ([]models.Record, error) {
	var records []models.Record

	if err := database.Db.Joins("JOIN spies ON spies.id = records.spy_id").Where("spies.user_id = ?", userId).Find(&records).Error; err != nil {
		return nil, ServiceError{
			Code:    500,
			Message: "Error while retrieving user records: " + err.Error(),
		}
	}

	return records, nil
}

func UpdateSpy(spyId string, req requestmodels.NewSpyRequest, userId uint) error {
	var spy models.Spy

	if err := database.Db.First(&spy, spyId).Error; err != nil {
		return ServiceError{
			Code:    404,
			Message: fmt.Sprintf("Spy with ID %s not found", spyId),
		}
	}

	// Vérifiez que l'utilisateur est le propriétaire du spy
	if spy.UserId != userId {
		return ServiceError{
			Code:    403,
			Message: "Unauthorized to update this spy",
		}
	}

	spy.Name = req.Name
	spy.Color = req.Color

	if err := database.Db.Save(&spy).Error; err != nil {
		return ServiceError{
			Code:    500,
			Message: "Error while updating spy: " + err.Error(),
		}
	}

	return nil
}

func DeleteSpy(spyId string, userId uint) error {
	var spy models.Spy

	if err := database.Db.First(&spy, spyId).Error; err != nil {
		return ServiceError{
			Code:    404,
			Message: fmt.Sprintf("Spy with ID %s not found", spyId),
		}
	}

	// Vérifiez que l'utilisateur est le propriétaire du spy
	if spy.UserId != userId {
		return ServiceError{
			Code:    403,
			Message: "Unauthorized to delete this spy",
		}
	}

	if err := database.Db.Delete(&spy).Error; err != nil {
		return ServiceError{
			Code:    500,
			Message: "Error while deleting spy: " + err.Error(),
		}
	}

	return nil
}

func DeleteRecord(recordId string, userId uint) error {
	var record models.Record

	// Récupérer le record
	if err := database.Db.First(&record, recordId).Error; err != nil {
		return ServiceError{
			Code:    404,
			Message: fmt.Sprintf("Record with ID %s not found", recordId),
		}
	}

	// Vérifier si le spy auquel appartient ce record appartient à l'utilisateur
	var spy models.Spy
	if err := database.Db.First(&spy, record.SpyID).Error; err != nil {
		return ServiceError{
			Code:    404,
			Message: fmt.Sprintf("Spy not found for Record ID %s", recordId),
		}
	}

	if spy.UserId != userId {
		return ServiceError{
			Code:    403,
			Message: "Unauthorized to delete this record",
		}
	}

	// Supprimer le record
	if err := database.Db.Delete(&record).Error; err != nil {
		return ServiceError{
			Code:    500,
			Message: "Error while deleting record: " + err.Error(),
		}
	}

	return nil
}
