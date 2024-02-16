// CAR SUV TRUCK VAN = {VehicleType }
// VEHICLE{vehicletype, locationAddress, isAvailabe, []*ActiveReservations, Barcode, member}
// MEMBER{name, email, phonenumber, []*ActiveReservations}
// RESERVATIONS{id, *MEMBER, Vehicle, PickDate, ReturnDate}

// member -> Vehicle
// Vehicle -> member

// CAR_RENTAL{
// 	map<id , RESERVATIONS>
// 	[]*Vehicle
// 	RegisterMember()
// 	ReserveVehicle(RetualInsurance, AdditionEquipments) //member CheckVehicaleAvailble
// 	CancelReservation() //member
// 	PickUpCar() // reservations
// 	ReturnCar() //resevations delete reservation after completing
// 	notifyDueDate()
// 	notifyPickUpDate()
// }

//Extension: 
// 1) Number of Kilometers if run and charge according to it
// 2) Enter Location, User can get the list of cars availabe during that days
// 3) User should be able to select car from given list

package main
import "fmt"

const{
	CAR int = iota // 0
	SUV            // 1
	TRUCK          // 2
}

type Vehicle struct{
	vehicleType int
	locationAddress string
	isAvailable bool
	activeReservations []*Reservations
	barcode string
	member *Member
}

type Member struct{
	memberID string //unique
	name string
	email string 
	phonenumber string 
	activeReservations []*Reservations
}

type Reservations struct{
	id int
	member *Member
	vehicle *Vehicle
	pickupDate string
	returnDate string
}


struct CarRental struct{
	idMap map[int]*Reservations
	vehicleList []*Vehicle
	memberList []*Member
}

// addNewVehicleToRentalSystem()
// removeVehicleFromRentalSystem()

func(c* CarRental)RegisterNewMember(name string, phoneNumber string, email string){
	// create a member
	newMember = Member{}
	c.memberList = append(c.memberList , newMember)
}
func (c* CarRental)ReserveVehicle(member Member, pickupDate string, returnDate string)Reservations, err{
	// ifvehicle availabe or not
	if VehicleAvailabe() == false{
		return err
	}
	// Build Reservation and return with success err
}

func(c *CarRental)CancelReservation(id){
	rerveration := c.idMap[id]
	// remove reservation from member, car and then rental system
}

func(c *CarRental)carPickup(id){
	rerveration := c.idMap[id]
	// logic for pickupcar
}

func(c *CarRental)returnCar(id){
	// remove reservation from car, member and carrental
	// calculate charge, per hour and late return extra charge
}


func (c* CarRental)Notification(id){ //cronjob scheduledjob for everyday or startinofWeek
	currentDate := time.Time()
	for i := 0; i < len(c.vehicleList); i++{
		// logic for pickupDate
		// logic for returnDate notificatons
	}
}


func main{
	// initRentalSystem
	// Add all vehicles in rental system, with details
	
	// add member and signup, generate memberID
	// someone has memberID, grab *Member object using map

	// makeReservation
	//return vehicle
}