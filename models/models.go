package models

import (
	"time"
	"github.com/google/uuid"
)

type PersonalDetails struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	Email        string    `json:"email"`
	CompanyID    uuid.UUID `json:"company_id"`
	Password     string    `json:"password"`
	JobTitle     string    `json:"job_title"`
	MobileNumber int       `json:"mobile_number"`
	CreatedAt    time.Time `json:"created_at"`
}

type CompanyDetails struct {
	ID          uuid.UUID `json:"company_id"`
	CompanyName string    `json:"companyname"`
	Branch      string    `json:"branch"`
	Address     string    `json:"address"`
	Address2    string    `json:"address2"`
	Country     string    `json:"country"`
	Province    string    `json:"province"`
	District    string    `json:"district"`
	SubDistrict string    `json:"subdistrict"`
	Zipcode     int       `json:"zipcode"`
	Website     string    `json:"website"`
	CreatedAt   time.Time `json:"created_at"`
}

type RequestRegister struct {
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	Email          string `json:"email"`
	JobTitle       string `json:"jobtitle"`
	MobileNumber   int    `json:"mobile_number"`
	Password       string `json:"Password"`
	CompanyName    string `json:"companyname"`
	Branch         string `json:"branch"`
	Address        string `json:"address"`
	Address2       string `json:"address2"`
	Country        string `json:"country"`
	Province       string `json:"province"`
	District       string `json:"district"`
	SubDistrict    string `json:"subdistrict"`
	Zipcode        int    `json:"zipcode"`
	Website        string `json:"website"`
	Requiresaction string `json:"requires_action"`
}
