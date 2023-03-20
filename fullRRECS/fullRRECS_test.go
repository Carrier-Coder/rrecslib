package fullRRECS

import (
	"math"
	"testing"

	"rrecslib.rrecsulator.com/dataSets"
	"rrecslib.rrecsulator.com/standards"
)

func sampleDailyData() dataSets.DailyData {
	// data from 21 Jul
	dd := dataSets.DailyData{

		//letters
		CasedDPSLetters:  100,
		DPSLetters:       350,
		WSSLetters:       0,
		BoxholderLetters: 0,

		//flats
		PreBundledFlats: 0,
		DPSFlats:        0,
		WSSFlats:        0,
		BoxholderFlats:  0,

		//door trips
		DoorTrips: 15,

		//driveways
		LessSwimPool:     2,
		SwimPool:         4,
		FootballField:    2,
		TwoFootballField: 1,
		QuarterMile:      1,
		HalfMile:         1,

		//parcels
		MailBoxParcels: 17,
		LockerParcels:  0,
		DoorParcels:    18,

		RegCurbBoxesSkipped: 15,

		StampSale:  0,
		RuralReach: 1,

		OneStepScan: 20,
		TwoStepScan: 20,

		ExtraSteps: 744,

		SafetyServiceTalks: 5,
		Loading:            23,
		EndOfShift:         29,
	}
	return dd
}

func sampleFixedData() dataSets.FixedData {
	fd := dataSets.FixedData{
		//from survey
		RawLetters: 27,
		RawFlats:   29,
		Form3982:   5,
		MiscTime:   10,
		OfficeFeet: 400,

		RegCurbBox:       238,
		DrivePOV:         true,
		DismountOther:    2,
		DismountFeet:     100,
		LowVolumePouches: 2,
		Miles:            92,
		DriveTime:        192.1,
		FeetPerStep:      2.5,
		Bundle:           standards.TWO_BUNDLE,
	}
	return fd
}

func TestFull(t *testing.T) {
	dd := sampleDailyData()
	fd := sampleFixedData()

	got := GetTime(fd, dd)
	want := 407.238

	if math.Abs(got-want) > 0.1 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}

}
