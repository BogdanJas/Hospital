package services

import "github.com/BogdanJas/Hospital/models"

type PatientService interface {
	CreatePatient(*models.Patient) error
	GetPatient(*int) (*models.Patient, error)
	GetAll() ([] *models.Patient, error)
	UpdatePatient(*models.Patient) error
	DeletePatient(*int) error
}