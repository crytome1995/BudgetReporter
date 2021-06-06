package date

import (
	"fmt"
	"time"

	"github.com/crytome1995/BudgetReporter/pkg/aws"
)

// GenerateDateRange return  the dates that should be used as a query for dinamoDB
func GenerateDateRange(t time.Time) []string {
	var dates []string
	day := t.Day()
	month := int(t.Month())
	year := t.Year()
	dates = append(dates, formatDate(day, month, year))
	// find the previous 6 days dates
	for r := 0; r < 6; {
		day, month, year = calcDate(day, month, year)
		dates = append(dates, formatDate(day, month, year))
		r++
	}
	return dates
}

// find the previous date based on provided date
func calcDate(d int, m int, y int) (determinedDay, determinedMonth, determinedYear int) {
	// first day of month back back one month
	if d == 1 {
		// first month of year go back one year
		if m == 1 {
			determinedMonth = 12
			determinedYear = y - 1
		} else {
			determinedMonth = m - 1
			determinedYear = y
		}
		determinedDay = lastDayOfTheMonth(y, determinedMonth).Day()
	} else {
		determinedDay = d - 1
		determinedMonth = m
		determinedYear = y
	}
	return determinedDay, determinedMonth, determinedYear
}

// format the date string
func formatDate(d int, m int, y int) string {
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

// call if we have to go back a month when subtracting a day
func lastDayOfTheMonth(year, month int) time.Time {
	if month++; month > 12 {
		month = 1
	}
	t := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	return t
}

// Sum returns the sum of all transactions amounts
func Sum(transactions []aws.Transaction) float64 {
	sum := float64(0)
	for i := 0; i < len(transactions); {
		sum += transactions[i].Amount
		i++
	}
	return sum
}
