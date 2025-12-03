package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	b := food{preptime: 15}
	c := food{preptime: 10}
	n := food{preptime: 12}

	switch order {
	case "burger":
		return b.preptime
	case "chips":
		return c.preptime
	case "nuggets":
		return n.preptime
	default:
		return 404
	}
}
