package models

import "time"

type Diseases struct {
	Fracture  bool `json:"fracture" bson:"fracture"`
	Heartburn bool `json:"heartburn" bson:"heartburn"`
	Headache  bool `json:"headache" bson:"headache"`
	Sneeze    bool `json:"sneeze" bson:"sneeze"`
	Heat      bool `json:"heat" bson:"heat"`
	Rash      bool `json:"rash" bson:"rash"`
}

type Address struct {
	State      string `json:"state" bson:"state"`
	City       string `json:"city" bson:"city"`
	Street     string `json:"street" bson:"street"`
	Postalcode string `json:"postcode" bson:"postcode"`
}

type Patient struct {
	Id          int       `json:"id" bson:"patient_id"`
	Name        string    `json:"name" bson:"patient_name"`
	Surname     string    `json:"surname" bson:"patient_surname"`
	Age         int       `json:"age" bson:"patient_age"`
	Address     Address   `json:"address" bson:"patient_address"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	VisitDate   time.Time `json:"visitDate" bson:"visitDate"`
	BloodGroup  string    `json:"groupOfBlood" bson:"patient_groupOfBlood"`
	Diseases    Diseases  `json:"diseases" bson:"patient_diseases"`
}
