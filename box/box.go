package box

import (
	"fmt"

	ds "rrecslib.rrecsulator.com/dataSets"
	s "rrecslib.rrecsulator.com/standards"
)

// BoxData contains all the info for calculating box times
type BoxData struct {
	RegCurbBox  int
	RegSdwkBox  int
	RegOtherBox int

	// Box Coverage Factor
	RegCurbBoxesSkipped  int
	RegSdwkBoxesSkipped  int
	RegOtherBoxesSkipped int

	VerifyFlats   int
	VerifyLetters int

	Bundle s.Bundling
}

func (bd *BoxData) Populate(dd ds.DailyData, fd ds.FixedData) {
	bd.RegCurbBox = fd.RegCurbBox
	bd.RegSdwkBox = fd.RegSdwkBox
	bd.RegOtherBox = fd.RegOtherBox

	bd.RegCurbBoxesSkipped = dd.RegCurbBoxesSkipped
	bd.RegSdwkBoxesSkipped = dd.RegSdwkBoxesSkipped
	bd.RegOtherBoxesSkipped = dd.RegOtherBoxesSkipped

	bd.VerifyFlats = fd.RawFlats + dd.PreBundledFlats + dd.DPSFlats
	bd.VerifyFlats += dd.WSSFlats

	bd.VerifyLetters = fd.RawLetters + dd.DPSLetters + dd.WSSLetters + dd.CasedDPSLetters

	bd.Bundle = fd.Bundle
}

func (bd *BoxData) Report() string {
	//The curbBoxTime function subtracts the skipped boxes.
	//The box Skipped Time is already calculated out...
	out := "\n********** Box Time **************\n"
	out += fmt.Sprintf("Curb Box Time: %4.2f\n", bd.curbBoxTime())
	out += fmt.Sprintf("Sidewalk Box Time: %4.2f\n", bd.sidewalkBoxTime())
	out += fmt.Sprintf("Other Box Time: %4.2f\n", bd.otherBoxTime())
	out += fmt.Sprintf("Verify Letter Time: %4.2f\n", bd.verifyLetterTime())
	out += fmt.Sprintf("Verify Flat Time: %4.2f\n", bd.verifyFlatTime())
	out += fmt.Sprintf("\nBox Skipped Time: %4.2f\n", bd.skippedBoxTime())
	out += fmt.Sprintf("\nBox Time Total: %4.2f\n", bd.GetTime())
	return out
}

func (bd *BoxData) GetTime() float64 {

	curbTime := bd.curbBoxTime()
	sdwkTime := bd.sidewalkBoxTime()
	otherTime := bd.otherBoxTime()

	verifyTime := bd.verifyLetterTime() + bd.verifyFlatTime()
	boxTime := 0.0
	boxTime = (curbTime + sdwkTime + otherTime + verifyTime)

	return boxTime

}

func (bd *BoxData) skippedBoxTime() float64 {
	time := 0.0
	time += s.RegBox[s.CURB][bd.Bundle] * float64(bd.RegCurbBoxesSkipped)
	time += s.RegBox[s.SIDEWALK][bd.Bundle] * float64(bd.RegSdwkBoxesSkipped)
	time += s.RegBox[s.OTHER][bd.Bundle] * float64(bd.RegOtherBoxesSkipped)
	return time
}

func (bd *BoxData) curbBoxTime() float64 {
	boxes := bd.RegCurbBox - bd.RegCurbBoxesSkipped
	return s.RegBox[s.CURB][bd.Bundle] * float64(boxes)
}

func (bd *BoxData) sidewalkBoxTime() float64 {
	boxes := bd.RegSdwkBox - bd.RegSdwkBoxesSkipped
	//S103 part of dismounts (could dismount for multiple sidewalk boxes)
	return s.RegBox[s.SIDEWALK][bd.Bundle] * float64(boxes)
}

func (bd *BoxData) otherBoxTime() float64 {
	boxes := bd.RegOtherBox - bd.RegOtherBoxesSkipped
	//S107 part of dismounts (dismount for multiple other boxes for example)
	return s.RegBox[s.OTHER][bd.Bundle] * float64(boxes)
}

func (bd *BoxData) verifyLetterTime() float64 {
	return float64(bd.VerifyLetters) * s.VERIFY_LETTER
}

func (bd *BoxData) verifyFlatTime() float64 {
	return float64(bd.VerifyFlats) * s.VERIFY_FLAT
}
