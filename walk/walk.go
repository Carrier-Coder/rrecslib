package walk

import (
	"fmt"

	ds "rrecslib.rrecsulator.com/dataSets"
	"rrecslib.rrecsulator.com/standards"
)

// Time for taking things to the door.
// check walking and driveway for those specific elements
type WalkData struct {
	Steps        int
	FeetPerStep  float64
	DismountFeet int
	OfficeFeet   int
}

func (ww *WalkData) Populate(dd ds.DailyData, fd ds.FixedData) {
	ww.Steps = dd.ExtraSteps
	ww.OfficeFeet = fd.OfficeFeet
	ww.DismountFeet = fd.DismountFeet
	ww.FeetPerStep = fd.FeetPerStep
}

func (ww *WalkData) Report() string {
	out := "\n************ Walking Time **************\n"
	out += fmt.Sprintf("Package Deliveries: %d steps\n", ww.Steps)
	out += fmt.Sprintf("Stride length: %4.2f feet per step\n", ww.FeetPerStep)
	out += fmt.Sprintf("Measured Office feet: %d feet\n", ww.OfficeFeet)
	out += fmt.Sprintf("Measured Dismount feet: %d feet\n", ww.DismountFeet)
	out += fmt.Sprintf("\nWalking Total Time %4.2f\n", ww.GetTime())
	return out
}

func (ww *WalkData) GetTime() float64 {
	out := 0.0
	out += float64(ww.Steps) * ww.FeetPerStep
	out += float64(ww.OfficeFeet)
	out += float64(ww.DismountFeet)
	return out * standards.WALKING_STANDARD
}
