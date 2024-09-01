//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	if len(guests) == 0 {
		return nil
	}

	maxCheckOutDate := 0
	for _, g := range guests {
		if g.CheckOutDate > maxCheckOutDate {
			maxCheckOutDate = g.CheckOutDate
		}
	}

	table := make([]int, maxCheckOutDate+1)
	for _, guest := range guests {
		table[guest.CheckInDate]++
		table[guest.CheckOutDate]--
	}

	res := make([]Load, len(table))

	guestcount := 0
	index := 0
	for date, people := range table {
		guestcount += people
		if people != 0 {
			res[index] = Load{StartDate: date, GuestCount: guestcount}
			index++
		}
	}
	return res[:index]
}
