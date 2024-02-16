// 1) multiple floors in parking lot
// 2) multiple gates in parking lot
// 3) parking attendant at gates of entry and exit
// 4) entry gate should process parking ticket
// 5) exit gate should process payment on hourly basis
// 6) Different type of vechicle facility
// 7) Display board of vacant spots on floors and

// Classes: Floor, EntryGate, Exitgate, ParkingLot, VehicleTicket, Vehicle
// Methods: DisplayParkingLotLayout, CarEntry, CarExit, UpdateParkingAttendant, calculateRate

// vehicle come in parkinglot => VehicleEntry()

package main

import "fmt"

const (
	CAR   = 1
	BIKE  = 2
	TRUCK = 3
)
const (
	CAR_COUNT   = 1
	BIKE_COUNT  = 2
	TRUCK_COUNT = 3
)

type ParkingLot struct {
	floors         []Floor
	entry1, entry2 EntryGate
	exit1, exit2   ExitGate
	parkingName    string
	rate           int
}
type Floor struct {
	VehicleType []int
}

type EntryGate struct {
	attendantName string
}

type ExitGate struct {
	attendantName string
}

type VehicleTicket struct {
	date        string
	vehicleType int
	floorNumber int

	// vehicleNumber string
	// attendantName string
}

type Vehicle struct {
	vehicleType int
	date        string
}

func BuildFloorsLayout()[]Floor{
	floorSchema := make([]Floor , 3)
	for i, _ := range floorSchema{
		floorSchema[i].VehicleType = []int{CAR_COUNT , BIKE_COUNT , TRUCK_COUNT}
	}
	return floorSchema
}

func NewParkingLot(parkingLotName string, rateValue int) *ParkingLot {
	return &ParkingLot{
		floors: BuildFloorsLayout(),
		entry1: EntryGate{
			attendantName: "Sam",
		},
		entry2: EntryGate{
			attendantName: "Alex",
		},
		exit1: ExitGate{
			attendantName: "Robin",
		},
		exit2: ExitGate{
			attendantName: "Zoro",
		},
		rate: rateValue,
	}
}

func GetTime()string{
	return ""
}

func (p *ParkingLot) VehicleEntry(vehicle Vehicle, entryGate *EntryGate) *VehicleTicket {
	floorNumber_ := p.vehicleSpaceAvail(vehicle.vehicleType)
	if floorNumber_ == -1 {
		fmt.Println("Space not available for given vehicle type")
		// Panic() Can add panic state here
		return nil
	}
	return &VehicleTicket{
		date: GetTime(),
		vehicleType: vehicle.vehicleType,
		floorNumber: floorNumber_,
	}
}
func (p *ParkingLot) vehicleSpaceAvail(vehicleType int) int {
	for floorNumber, currentFloor := range p.floors {
		if currentFloor.VehicleType[vehicleType] > 0 {
			currentFloor.VehicleType[vehicleType]--
			return floorNumber
		}
	}
	return -1
}

func (p *ParkingLot)VehicleExit(ticket *VehicleTicket)int{
	p.floors[ticket.floorNumber].VehicleType[ticket.vehicleType]++
	return 0
	// CalculateRate(p.rate)
}

func Display(){
	// will display parking lot 
}

func main() {

}
