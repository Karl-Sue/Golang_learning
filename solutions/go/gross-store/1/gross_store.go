package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int {
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	num, ok := units[unit]
    if ok {
        _ , ok2 := bill[item]
        if !ok2 {
            bill[item] = num
        } else {
            bill[item] += num
        }
        return true
    }
    return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	num, ok := bill[item]
    if !ok {
        return false
    }
    num2 , ok2 := units[unit]
    if !ok2 {
        return false
    }
    if num < num2 {
        return false
    } else if num == num2 {
        delete(bill,item)
        return true
    } else {
        bill[item] -= num2
        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty , ok := bill[item]
    return qty, ok
}
