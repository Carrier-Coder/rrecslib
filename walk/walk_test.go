package walk

import (
	"math"
	"testing"

	ds "rrecslib.rrecsulator.com/dataSets"

	"rrecslib.rrecsulator.com/standards"
)

func sampleWalkData() WalkData {
	return WalkData{
		Steps:       520,
		FeetPerStep: standards.AVG_FEET_PER_STEP,
	}
}

func TestGetTime(t *testing.T) {
	ww := sampleWalkData()
	got := ww.GetTime()
	want := 5.577

	if math.Abs(got-want) > .001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestPopulateAndTime(t *testing.T) {

	dd := ds.DailyData{}
	dd.ExtraSteps = 100
	fd := ds.FixedData{}
	fd.OfficeFeet = 100
	fd.DismountFeet = 100
	fd.FeetPerStep = 2.5

	ww := WalkData{}
	ww.Populate(dd, fd)
	got := ww.GetTime()
	want := 1.931

	if math.Abs(got-want) > .001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
