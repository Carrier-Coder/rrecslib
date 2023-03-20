package dataSets

type DailyData struct {

	// Letters
	DPSLetters       int `json:"DPSLetters,string"`
	WSSLetters       int `json:"WSSLetters,string"`
	BoxholderLetters int `json:"BoxholderLetters,string"`
	CasedDPSLetters  int `json:"CasedDPSLetters,string"`

	// Flats
	PreBundledFlats int `json:"PreBundledFlats,string"`
	DPSFlats        int `json:"DPSFlats,string"`
	WSSFlats        int `json:"WSSFlats,string"`
	BoxholderFlats  int `json:"BoxholderFlats,string"`
	CasedDPSFlats   int `json:"CasedDPSFlats,string"`

	//DoorTrips
	DoorTrips int `json:"DoorTrips,string"`

	//Driveways
	LessSwimPool     int `json:"LessSwimPool,string"`
	SwimPool         int `json:"SwimPool,string"`
	FootballField    int `json:"FootballField,string"`
	TwoFootballField int `json:"TwoFootballField,string"`
	QuarterMile      int `json:"QuarterMile,string"`
	HalfMile         int `json:"HalfMile,string"`

	//Parcels
	MailBoxParcels int `json:"MailBoxParcels,string"`
	LockerParcels  int `json:"LockerParcels,string"`
	DoorParcels    int `json:"DoorParcels,string"`

	//Accountables
	Certified           int `json:"Certified,string"`
	Registered          int `json:"Registered,string"`
	SignatureExpress    int `json:"SignatureExpress,string"`
	NonSignatureExpress int `json:"NonSignatureExpress,string"`
	CODs                int `json:"CODs,string"`
	PostageDueCustom    int `json:"PostageDueCustom,string"`
	PostageDue          int `json:"PostageDue,string"`

	SignatureParcel int `json:"SignatureParcel,string"`

	//Skipped boxes (box coverage factor)
	RegCurbBoxesSkipped  int `json:"RegCurbBoxesSkipped,string"`
	RegSdwkBoxesSkipped  int `json:"RegSdwkBoxesSkipped,string"`
	RegOtherBoxesSkipped int `json:"RegOtherBoxesSkipped,string"`

	CBUBoxesSkipped  int `json:"CBUBoxesSkipped,string"`
	CENTBoxesSkipped int `json:"CENTBoxesSkipped,string"`
	NPUBoxesSkipped  int `json:"NPUBoxesSkipped,string"`
	DETBoxesSkipped  int `json:"DETBoxesSkipped,string"`

	// miscellaneous trips to door (mdd entry)
	DoorMisc         int `json:"DoorMisc,string"`
	DoorPickupEvents int `json:"DoorPickupEvents,string"`

	//Additional Trips
	AddTripOther    int `json:"AddTripOther,string"`
	AddTripSidewalk int `json:"AddTripSidewalk,string"`
	AddTripCbu      int `json:"AddTripCbu,string"`
	AddTripCent     int `json:"AddTripCent,string"`
	AddTripNpu      int `json:"AddTripNpu,string"`
	AddTripDetVpo   int `json:"AddTripDetVpo,string"`

	//Scans
	//SPMs count as one step scans?
	OneStepScan int `json:"OneStepScan,string"`
	TwoStepScan int `json:"TwoStepScan,string"`
	StampSale   int `json:"StampSale,string"`
	RuralReach  int `json:"RuralReach,string"`

	//Package Pickups
	EventForms int `json:"EventForms,string"`
	Events     int `json:"Events,string"`
	PackageNum int `json:"PackageNum,string"`

	ExtraSteps int `json:"ExtraSteps,string"`

	//actual time in standard minutes
	Loading            float64 `json:"Loading,string"`
	EndOfShift         float64 `json:"EndOfShift,string"`
	Deviation          float64 `json:"Deviation,string"`
	SafetyServiceTalks float64 `json:"SafetyServiceTalks,string"`
	FudgeTime          float64 `json:"FudgeTime,string"`
}
