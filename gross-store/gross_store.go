package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	gross_map := map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }

    return gross_map
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    amt := units[unit]
	if amt != 0 {
        bill[item] += amt
        return true
    } else {
        return false
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    uAmt := units[unit]
    bAmt := bill[item]
    newAmt := bAmt - uAmt
	if bAmt == 0 || uAmt == 0 || newAmt < 0{
        return false
    } else if newAmt > 0{
        bill[item] = newAmt
        return true
    } else {
        delete(bill, item)
        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	amt := bill[item]

    if amt != 0 {
        return amt, true
    } else {
        return 0, false
    }
}
