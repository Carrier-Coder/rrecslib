package doorTrip

import (
	"math"
	"testing"
)

func sampleDoorData() DoorData {
	return DoorData{

		Trips:        19,
		Parcels:      20,
		Accountables: 1,
		Misc:         2,
		PickupEvents: 1,
	}
}

// adding another misc time entry ends up subtracting time?
// yes, if door trip already counted, than adding an additional misc
// actually decreases the eval (effectively a door trip was counted as unique
// with additional fixed gather time, but it really wasnt
func TestDoorMiscTimebug(t *testing.T) {
	dd := sampleDoorData()
	got := dd.GetTime()
	dd.Misc++
	got2 := dd.GetTime()
	if got2 > got {
		t.Errorf("%v before %v after", got, got2)
	}
}

func TestDoorMisctime(t *testing.T) {
	dd := sampleDoorData()
	got := dd.doorMiscTime()
	want := 2 * 0.0854

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetTime(t *testing.T) {
	dd := sampleDoorData()
	got := dd.GetTime()
	want := 20.0448

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
