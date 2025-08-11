package models

import "sync"

type PersonalInfo struct {
	Title       string `json:title`
	Description string `json:description`
	Address     string `json:addr`
	Email       string `json:email`
	PhoneNumber string `json:phone_number`
}

// using defeult value for it 
var PersonalData = PersonalInfo{
	Title: "",
	Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Aperiam sed animi et laboriosam facere tempora rerum, eligendi veritatis nesciunt cum?",
	Address: "",
	Email: "taha@gmail.com",
	PhoneNumber: "0912000000000",
}

var PersonalMu sync.RWMutex