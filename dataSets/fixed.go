package dataSets

import "rrecslib.rrecsulator.com/standards"

type FixedData struct {

	//data from the mini mail survey
	RawLetters int     `json:"RawLetters,string"`
	RawFlats   int     `json:"RawFlats,string"`
	Form3982   int     `json:"Form3982,string"`
	MiscTime   float64 `json:"MiscTime,string"`
	OfficeFeet int     `json:"OfficeFeet,string"`

	//number of mailboxes on the route
	RegCurbBox  int `json:"RegCurbBox,string"`
	RegSdwkBox  int `json:"RegSdwkBox,string"`
	RegOtherBox int `json:"RegOtherBox,string"`

	//vehicle information
	DrivePOV   bool `json:"DrivePOV,string"`
	DriveGOVLH bool `json:"DriveGOVLH,string"`

	BlueBox int `json:"BlueBox,string"`

	//Centralized boxes on the route
	CBUs            int `json:"CBUs,string"`
	CBUBoxes        int `json:"CBUBoxes,string"`
	CENTs           int `json:"CENTs,string"`
	CENTBoxes       int `json:"CENTBoxes,string"`
	NPUs            int `json:"NPUs,string"`
	NPUBoxes        int `json:"NPUBoxes,string"`
	DETs            int `json:"DETs,string"`
	DETBoxes        int `json:"DETBoxes,string"`
	CollectionBoxes int `json:"CollectionBoxes,string"`

	//Authorized fixed dismounts
	DismountOther    int `json:"DismountOther,string"`
	DismountSidewalk int `json:"DismountSidewalk,string"`
	DismountCbu      int `json:"DismountCbu,string"`
	DismountCent     int `json:"DismountCent,string"`
	DismountNpu      int `json:"DismountNpu,string"`
	DismountDetVpo   int `json:"DismountDetVpo,string"`

	LowVolumePouches  int `json:"LowVolumePouches,string"`
	HighVolumePouches int `json:"HighVolumePouches,string"`

	//Route mileage and drive time
	Miles     float64 `json:"Miles,string"`
	DriveTime float64 `json:"DriveTime,string"`

	//Daily steps or walking distance in feet
	FeetPerStep  float64 `json:"FeetPerStep,string"`
	DismountFeet int     `json:"DismountFeet,string"`

	Bundle standards.Bundling `json:"Bundle,string"`

	WithdrawMail bool `json:"WithdrawMail,string"`
}
