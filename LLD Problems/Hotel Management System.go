RoomType = {standard, deluxe, familySuite}

Room{
	Owner *Customer
	RoomType int
	RoomNo int
	isAvailabe bool
}

Customer{
	customerID string
	BookedRooms []*Room
	name string
	email string
	phoneNo string	
}

Reservation{
	BookedRooms []Room
	Owner *Customer
	CheckinTime string
	CheckoutTime String
}

Hotel{
	map[int]*Room        // map <RoomNo, *Room>
	map[int]*Customer // map  <CustomerID , *Customer>
	RoomsList []Room
	CustomerList []*Customer
	ActiveReservation []Reservation

	Booking(Customer, CountofStandard, CountofDeluxe, CountofFamilySuite, Services{})Reservation
	CancelBooking(Reservation)
	Notification() //both for checkin date and Checkoutdate
	getOwnerOfRoom(RoomNo)*Customer
	SearchRoom()
 
}

