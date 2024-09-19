package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {

	grossUnits := map[string]int{}

	grossUnits["quarter_of_a_dozen"] = 3
	grossUnits["half_of_a_dozen"] = 6
	grossUnits["dozen"] = 12
	grossUnits["small_gross"] = 120
	grossUnits["gross"] = 144
	grossUnits["great_gross"] = 1728

	return grossUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	if _, existsUnits := units[unit]; !existsUnits {

		return false
	}

	if _, existsBill := bill[item]; existsBill {

		bill[item] += units[unit]
		return true
	}
	bill[item] = units[unit]
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	if _, existsBill := bill[item]; !existsBill {

		return false
	}

	if _, existsUnits := units[unit]; !existsUnits {

		return false
	}

	if newUnit := bill[item] - units[unit]; newUnit < 0 {

		return false
	}

	if newUnit := bill[item] - units[unit]; newUnit == 0 {

		delete(bill, item)
		return true
	}

	bill[item] -= units[unit]

	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, existsBill := bill[item]

	if !existsBill {
		return 0, false
	}

	return quantity, true
}
