package date

import (
	"testing"
	"time"

	"github.com/crytome1995/BudgetReporter/pkg/aws"
)

func TestBegginingofYearBeginningofMonthDate(t *testing.T) {
	t.Log("TestBegginingofYearBeginningofMonthDate")
	testDate := time.Date(2021, time.January, 1, 1, 1, 1, 1, time.UTC)
	expectedDates := []string{"2021-01-01", "2020-12-31", "2020-12-30", "2020-12-29", "2020-12-28", "2020-12-27", "2020-12-26"}
	actualDates := GenerateDateRange(testDate)
	if len(actualDates) != 7 {
		t.Fatalf("Expected length of 7 got %d", len(actualDates))
	}
	for i, v := range actualDates {
		if v != expectedDates[i] {
			t.Fatalf("Expected date %s actual date %s", expectedDates[i], v)
		}
	}
}

func TestBeginningofMonthDate(t *testing.T) {
	t.Log("TestBeginningofMonthDate")
	testDate := time.Date(2021, time.February, 1, 1, 1, 1, 1, time.UTC)
	expectedDates := []string{"2021-02-01", "2021-01-31", "2021-01-30", "2021-01-29", "2021-01-28", "2021-01-27", "2021-01-26"}
	actualDates := GenerateDateRange(testDate)
	if len(actualDates) != 7 {
		t.Fatalf("Expected length of 7 got %d", len(actualDates))
	}
	for i, v := range actualDates {
		if v != expectedDates[i] {
			t.Fatalf("Expected date %s actual date %s", expectedDates[i], v)
		}
	}
}

func TestSumTransactions(t *testing.T) {
	t.Log("TestSumTransactions")
	expected := 21.00
	transaction1 := aws.Transaction{"ethan", "ethan", 10.50, "ethan", "ethan", "2021-01-01"}
	transaction2 := aws.Transaction{"ethan", "ethan", 10.50, "ethan", "ethan", "2021-01-01"}
	transactions := []aws.Transaction{transaction1, transaction2}
	actual := Sum(transactions)
	if expected != actual {
		t.Fatalf("Expected sum %f actual sum %f", expected, actual)
	}

}
