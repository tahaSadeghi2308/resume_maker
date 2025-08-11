package models

import "sync"

type Experience struct {
	JobTitle    string `json:job_title`
	Description string `json:description`
	StartDate   string `json:start_date`
	EndDate     string `json:end_date`
}

var ExperienceData = Experience{
	JobTitle:    "Software Engineer",
	Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Aperiam sed animi et laboriosam facere tempora rerum, eligendi veritatis nesciunt cum?",
	StartDate:   "2010",
	EndDate:     "2015",
}

var ExperienceMu sync.RWMutex
