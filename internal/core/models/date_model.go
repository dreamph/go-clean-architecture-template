package models

import "time"

//type MonthData struct {
//	Year  int
//	Month int
//}

type CalDate struct {
	Year      int
	Month     int
	StartDate time.Time
	EndDate   time.Time
	Skip      bool
}

type CalDateRange struct {
	StartDate time.Time
	EndDate   time.Time
}

type CalCriteriaDateRange struct {
	MonthFrom int `json:"monthFrom"`
	MonthTo   int `json:"monthTo"`
	YearFrom  int `json:"yearFrom"`
	YearTo    int `json:"yearTo"`
}
