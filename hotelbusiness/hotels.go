//go:build !solution

package hotelbusiness

import (
	"cmp"
	"slices"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

type move struct {
	day     int
	checkIn bool
}

func ComputeLoad(guests []Guest) []Load {
	moves := make([]move, 0, len(guests)*2)
	for _, guest := range guests {
		moves = append(moves, move{guest.CheckInDate, true})
		moves = append(moves, move{guest.CheckOutDate, false})
	}
	slices.SortFunc(moves, func(a, b move) int {
		return cmp.Compare(a.day, b.day)
	})

	loads := []Load{}
	curGuestCnt := 0
	curday := 0
	prevGuestCnt := 0
	for i := 0; i < len(moves); {
		curday = moves[i].day
		for i < len(moves) && moves[i].day == curday {
			if moves[i].checkIn {
				curGuestCnt++
			} else {
				curGuestCnt--
			}
			i++
		}
		if curGuestCnt != prevGuestCnt {
			loads = append(loads, Load{curday, curGuestCnt})
			prevGuestCnt = curGuestCnt
		}
	}
	return loads
}
