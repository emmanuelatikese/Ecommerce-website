package response

import (
	"time"

)

func DateRange(startDate time.Time, endDate time.Time) []string{
	currentDate := startDate
	var dateList []string
	counter := -7
	for currentDate.Before(endDate){
		dateList = append(dateList, currentDate.Format("2006-01-02"))
		counter++
		currentDate = currentDate.AddDate(0,0,counter)
	}
	return dateList
}