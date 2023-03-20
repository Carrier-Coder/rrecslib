package fullRRECS

import (
	"bytes"
	"fmt"

	"rrecslib.rrecsulator.com/accountable"
	"rrecslib.rrecsulator.com/actual"
	"rrecslib.rrecsulator.com/averaged"
	"rrecslib.rrecsulator.com/box"
	"rrecslib.rrecsulator.com/boxholder"
	"rrecslib.rrecsulator.com/cbu"
	"rrecslib.rrecsulator.com/dataSets"
	"rrecslib.rrecsulator.com/dismount"
	"rrecslib.rrecsulator.com/doorTrip"
	"rrecslib.rrecsulator.com/driveway"
	"rrecslib.rrecsulator.com/flat"
	"rrecslib.rrecsulator.com/letter"
	"rrecslib.rrecsulator.com/markup"
	"rrecslib.rrecsulator.com/misc"
	"rrecslib.rrecsulator.com/parcel"
	"rrecslib.rrecsulator.com/pickup"
	"rrecslib.rrecsulator.com/pouch"
	"rrecslib.rrecsulator.com/reload"
	"rrecslib.rrecsulator.com/strapout"
	"rrecslib.rrecsulator.com/vehicle"
	"rrecslib.rrecsulator.com/walk"
	"rrecslib.rrecsulator.com/withdraw"
)

//all events that can produce a time
type timeable interface {
	Populate(dataSets.DailyData, dataSets.FixedData)
	Report() string
	GetTime() float64
}

// GetTime returns daily eval for a set of fixed and daily data
func GetTime(fd dataSets.FixedData, dd dataSets.DailyData) float64 {
	elements := []timeable{
		&actual.ActualData{},
		&accountable.AccountableData{},
		&averaged.AveragedData{},
		&box.BoxData{},
		&boxholder.BoxholderData{},
		&cbu.CBUData{},
		&dismount.DismountData{},
		&doorTrip.DoorData{},
		&driveway.DrivewayData{},
		&flat.FlatData{},
		&letter.LetterData{},
		&markup.MarkupData{},
		&misc.MiscData{},
		&parcel.ParcelData{},
		&pickup.PickupData{},
		&pouch.PouchData{},
		&strapout.StrapoutData{},
		&reload.ReloadData{},
		&vehicle.VehicleData{},
		&walk.WalkData{},
		&withdraw.WithdrawData{},
	}

	totalTime := 0.0
	for _, tb := range elements {
		tb.Populate(dd, fd)
		totalTime += tb.GetTime()
	}
	return totalTime
}

//GenearteReport returns a detailed report for a set of fixed and daily data
func GenerateReport(fd dataSets.FixedData, dd dataSets.DailyData) string {

	elements := []timeable{
		&actual.ActualData{},
		&accountable.AccountableData{},
		&averaged.AveragedData{},
		&box.BoxData{},
		&boxholder.BoxholderData{},
		&cbu.CBUData{},
		&dismount.DismountData{},
		&doorTrip.DoorData{},
		&driveway.DrivewayData{},
		&flat.FlatData{},
		&letter.LetterData{},
		&markup.MarkupData{},
		&misc.MiscData{},
		&parcel.ParcelData{},
		&pickup.PickupData{},
		&pouch.PouchData{},
		&strapout.StrapoutData{},
		&reload.ReloadData{},
		&vehicle.VehicleData{},
		&walk.WalkData{},
		&withdraw.WithdrawData{},
	}

	totalTime := 0.0
	var b bytes.Buffer

	for _, tb := range elements {
		tb.Populate(dd, fd)
		b.WriteString(tb.Report())
		totalTime += tb.GetTime()
	}
	outString := fmt.Sprintf("Days Total Eval Time: %4.2f minutes\n\n See below for a more detailed breakout\n", totalTime)

	//return outString
	return outString + b.String()
}
