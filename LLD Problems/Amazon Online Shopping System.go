package main

import (
	"time"
)

type Customer struct {
	CustomerID int
	ShoppingCart
	Search Search
}

func (c *Customer) GetShoppingCart() ShoppingCart {
	return c.ShoppingCart
}

func (c *Customer) AddItemsToShoppingCart(item Item) {
	c.ShoppingCart.AddItem(item)
}

func (c *Customer) UpdateItemFromShoppingCart(item Item) {
	c.ShoppingCart.UpdateItem(item)
}

func (c *Customer) RemoveItemFromShoppingCart(item Item) {
	c.ShoppingCart.DeleteItem(item)
}

type Guest struct {
	Customer
}

func (g *Guest) CreateNewAccount() Account {
	// Implementation
	return Account{}
}

type User struct {
	Customer
	Account Account
}

type Seller struct {
	User
}

func (s *Seller) AddProduct(product Product) {
	// Implementation
}

type Buyer struct {
	User
	Order Order
}

func (b *Buyer) AddReview(review ProductReview) {
	// Implementation
}

func (b *Buyer) PlaceOrder(cart ShoppingCart) OrderStatus {
	// Implementation
	return OrderStatus{}
}

type Account struct {
	Name             string
	Email            string
	PhoneNumber      string
	UserName         string
	Password         string
	ShippingAdresses []Address
	AccountStatus    AccountStatus
}

type Address struct {
	PinCode int
	Street  string
	City    string
	State   string
	Country string
}

type ShoppingCart struct {
	Items     []Item
	CartValue float64
}

func (s *ShoppingCart) AddItem(item Item) {
	// Implementation
}

func (s *ShoppingCart) UpdateItem(item Item) {
	// Implementation
}

func (s *ShoppingCart) DeleteItem(item Item) {
	// Implementation
}

func (s *ShoppingCart) CheckoutItems() {
	// Implementation
}

func (s *ShoppingCart) GetItems() []Item {
	return s.Items
}

func (s *ShoppingCart) GetCartValue() float64 {
	return s.CartValue
}

type Item struct {
	Product Product
	Qty     int
}

type Product struct {
	ProductID          int
	ProductDescription string
	Name               string
	ProductCategory    ProductCategory
	Seller             Seller
	Cost               float64
	ProductReviews     []ProductReview
}

type ProductCategory int

const (
	Electronics ProductCategory = iota
	Furniture
	Grocery
	Mobile
)

type ProductReview struct {
	Details  string
	Reviewer Buyer
	Rating   int
}

type Search struct{}

func (s *Search) SearchByName(name string) []Product {
	// Implementation
	return nil
}

func (s *Search) SearchByCategory(category ProductCategory) []Product {
	// Implementation
	return nil
}

type Order struct {
	OrderID      int
	OrderItem    []Item
	OrderValue   float64
	Buyer        Buyer
	OrderDate    time.Time
	Notification NotificationService
	OrderLog     []OrderLog
}

func (o *Order) PlaceOrder() OrderStatus {
	// Implementation
	return OrderStatus{}
}

func (o *Order) TrackOrder() OrderStatus {
	// Implementation
	return OrderStatus{}
}

func (o *Order) AddOrderLogs() {
	// Implementation
}

func (o *Order) MakePayment(paymentType PaymentType) PaymentStatus {
	// Implementation
	return PaymentStatus{}
}

type OrderStatus int

const (
	Packed OrderStatus = iota
	Shipped
	Enroute
	OutForDelivery
	Delivered
	Cancelled
)

type PaymentStatus int

const (
	Success PaymentStatus = iota
	Error
	CancelledPayment
	RefundInitiated
	Refunded
)

type PaymentType int

const (
	CreditCard PaymentType = iota
	DebitCard
	NetBanking
	UPI
)

type OrderLog struct {
	OrderDetail string
	CreatedDate time.Time
	Status      OrderStatus
}

type NotificationDomain struct {
	NotificationID   string
	NotificationType NotificationType
	User             User
}

type NotificationService struct{}

func (n *NotificationService) SendNotification(notificationDomain NotificationDomain) bool {
	// Implementation
	return false
}

type NotificationType int

const (
	Email NotificationType = iota
	WhatsApp
)

type Notification interface {
	SendNotification(messageAttribute MessageAttributes) bool
}

type EmailNotification struct{}

func (e *EmailNotification) SendNotification(messageAttribute MessageAttributes) bool {
	// Implementation
	return false
}

type WhatsAppNotification struct{}

func (w *WhatsAppNotification) SendNotification(messageAttribute MessageAttributes) bool {
	// Implementation
	return false
}

type SMSNotification struct{}

func (s *SMSNotification) SendNotification(messageAttribute MessageAttributes) bool {
	// Implementation
	return false
}

type MessageAttributes struct {
	To      string
	From    string
	Message string
}

func main() {
	// Main function
}


########################################################

package main

import (
	"fmt"
	"time"
)

// ProductCategory represents the category of a product.
type ProductCategory int

const (
	Electronics ProductCategory = iota
	Books
	HomeAppliances
	Clothing
	// Add more categories as needed
)

// Product represents a product available for sale.
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Category    ProductCategory
	Stock       int
	Reviews     []*Review
}

// Review represents a review for a product.
type Review struct {
	UserID  int
	Rating  int
	Comment string
}

// Order represents an order made by a user.
type Order struct {
	ID              int
	UserID          int
	Products        []*Product
	ShippingAddress string
	OrderStatus     OrderStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// OrderStatus represents the status of an order.
type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Shipped
	Delivered
	Cancelled
)

// User represents a user of the system.
type User struct {
	ID           int
	Username     string
	Email        string
	Password     string
	IsRegistered bool
	Orders       []*Order
	ShoppingCart ShoppingCart
}

// ShoppingCart represents the shopping cart of a user.
type ShoppingCart struct {
	Items []*Product
}

// ShoppingSystem represents the Amazon shopping system.
type ShoppingSystem struct {
	Products       []*Product
	Users          map[int]*User
	Orders         []*Order
	NotificationService NotificationService
}

// NewShoppingSystem creates a new ShoppingSystem instance.
func NewShoppingSystem() *ShoppingSystem {
	return &ShoppingSystem{
		Products:              []*Product{},
		Users:                 make(map[int]*User),
		Orders:                []*Order{},
		NotificationService:   &AmazonNotificationService{},
	}
}

// AddProduct allows users to add new products to sell.
func (s *ShoppingSystem) AddProduct(user *User, product *Product) error {
	// Implementation
	return nil
}

// SearchProductsByName searches for products by their name.
func (s *ShoppingSystem) SearchProductsByName(name string) []*Product {
	// Implementation
	return nil
}

// SearchProductsByCategory searches for products by their category.
func (s *ShoppingSystem) SearchProductsByCategory(category ProductCategory) []*Product {
	// Implementation
	return nil
}

// RegisterUser registers a user to become a member.
func (s *ShoppingSystem) RegisterUser(user *User) error {
	// Implementation
	return nil
}

// AddProductToCart adds a product to the user's shopping cart.
func (s *ShoppingSystem) AddProductToCart(user *User, product *Product) error {
	// Implementation
	return nil
}

// RemoveProductFromCart removes a product from the user's shopping cart.
func (s *ShoppingSystem) RemoveProductFromCart(user *User, product *Product) error {
	// Implementation
	return nil
}

// ModifyProductInCart modifies a product in the user's shopping cart.
func (s *ShoppingSystem) ModifyProductInCart(user *User, product *Product) error {
	// Implementation
	return nil
}

// Checkout checks out and buys items in the shopping cart.
func (s *ShoppingSystem) Checkout(user *User) (*Order, error) {
	// Implementation
	return nil, nil
}

// CancelOrder cancels an order if it has not shipped.
func (s *ShoppingSystem) CancelOrder(order *Order) error {
	// Implementation
	return nil
}

// AddReview adds a review for a product.
func (s *ShoppingSystem) AddReview(product *Product, userID int, rating int, comment string) error {
	// Implementation
	return nil
}

// SpecifyShippingAddress specifies a shipping address for an order.
func (s *ShoppingSystem) SpecifyShippingAddress(order *Order, address string) error {
	// Implementation
	return nil
}

// AmazonNotificationService represents the notification service for Amazon.
type AmazonNotificationService struct{}

// NotifyUser notifies the user of a change in order or shipping status.
func (ans *AmazonNotificationService) NotifyUser(user *User, message string) error {
	// Implementation
	fmt.Printf("Notification sent to user %s: %s\n", user.Username, message)
	return nil
}

func main() {
	// Initialize the shopping system
	shoppingSystem := NewShoppingSystem()

	// Simulate usage of the shopping system
	// (Creating users, adding products, searching, adding to cart, checking out, etc.)
}
