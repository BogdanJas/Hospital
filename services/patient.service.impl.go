package services

import (
	"context"
	"errors"

	"github.com/BogdanJas/Hospital/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PatientServiceImpl struct {
	patientCollection *mongo.Collection
	ctx 			context.Context
}

func NewPatientService(patientcollection *mongo.Collection, ctx context.Context) PatientService {
	return &PatientServiceImpl{
		patientCollection: patientcollection,
		ctx: ctx,
	}
}

func (u *PatientServiceImpl) CreatePatient(patient *models.Patient) error{
	_,err := u.patientCollection.InsertOne(u.ctx,patient)
	return err
}

func (u *PatientServiceImpl) GetPatient(id *int) (*models.Patient,error){
	var patient *models.Patient
	query := bson.D{bson.E{Key: "patient_id", Value:id}}
	err := u.patientCollection.FindOne(u.ctx,query).Decode(&patient)
	return patient,err
}

func (u *PatientServiceImpl) GetAll() ([] *models.Patient, error){
	var patients []*models.Patient
	cursor,err := u.patientCollection.Find(u.ctx, bson.D{{}})
	if err!=nil{
		return nil,err
	}
	for cursor.Next(u.ctx){
		var patient models.Patient
		err := cursor.Decode(&patient)
		if err!= nil{

			return nil,err
		}
		patients = append(patients,&patient)
	}

	if err := cursor.Err(); err!= nil{
		return nil,err
	}
	cursor.Close(u.ctx)
	if len(patients) == 0 {
		return nil, errors.New("documents not found")
	}
	return patients,nil
}

func (u *PatientServiceImpl) UpdatePatient(patient *models.Patient) error{
	filter := bson.D{bson.E{Key: "patient_id", Value: patient.Id}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key:"patient_id", Value:patient.Id},bson.E{Key:"patient_name", Value:patient.Name},bson.E{Key:"patient_surname", Value:patient.Surname},bson.E{Key:"patient_age", Value:patient.Age},bson.E{Key:"patient_address", Value:patient.Address},bson.E{Key:"patient_phoneNumber", Value:patient.PhoneNumber},bson.E{Key:"patient_groupOfBlood", Value:patient.BloodGroup},bson.E{Key:"patient_diseas", Value:patient.Diseas}}}}
	result,_ := u.patientCollection.UpdateOne(u.ctx,filter,update)
	if result.MatchedCount !=1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *PatientServiceImpl) DeletePatient(id *int) error{
	filter := bson.D{bson.E{Key: "patient_id", Value: id}}
	result, _ := u.patientCollection.DeleteOne(u.ctx,filter)
	if result.DeletedCount !=1 {
		return errors.New("no matched document found for update")
	}
	return nil
}