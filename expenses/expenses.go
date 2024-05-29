package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var results []Record

    for _, item := range in {
        if predicate(item) {
            results = append(results, item)
        }
    }

    return results
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(r Record) bool {
	return func(r Record) bool {
        if r.Day >= p.From && r.Day <= p.To {
            return true
        } else {
            return false
        }
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
        if c == r.Category {
            return true
        } else {
            return false
        }
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    sum := 0.0
	filtered := Filter(in, ByDaysPeriod(p))

    for _, item := range filtered {
        sum += item.Amount
    }
    
    return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	sum := 0.0
    var err error
    
    filtered_days := Filter(in, ByDaysPeriod(p))
    if len(filtered_days) > 0 {
        filtered_cat := Filter(filtered_days, ByCategory(c))
        if len(filtered_cat) > 0 {
            for _, item := range filtered_cat {
            sum += item.Amount
            }
        } else {
            err = errors.New("Unknown Category")
        }
    } 
    return sum, err
}
