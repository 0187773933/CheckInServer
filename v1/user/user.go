package user

import (
	// "fmt"
	uuid "github.com/satori/go.uuid"
	// server "github.com/0187773933/GO_SERVER/v1/server"
	utils "github.com/0187773933/CheckInServer/v1/utils"
)

type CheckIn struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	ULID string `json:"ULID"`
	Date string `json:"date"`
	Time string `json:"time"`
	Type string `json:"type"`
	Result bool `json:"result"`
	TimeRemaining int `json:"time_remaining"`
	// PrintJob printer.PrintJob `json:"print_job"`
}

// type FailedCheckIn struct {
// 	Date string `json:"date"`
// 	Time string `json:"time"`
// 	Type string `json:"type"`
// 	DaysRemaining int `json:"remaining_days"`
// }

type DateOfBirth struct {
	Month string `json:"month"`
	Day int `json:"day"`
	Year int `json:"year"`
}

type Address struct {
	StreetNumber string `json:"street_number"`
	StreetName string `json:"street_name"`
	AddressTwo string `json:"address_two"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zipcode"`
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Address Address`json:"address"`
	DateOfBirth DateOfBirth `json:"date_of_birth"`
	Age int `json:"age"`
	Sex string `json:"sex"`
	Height string `json:"height"`
	EyeColor string `json:"eye_color"`
	Spouse bool `json:"spouse"`
	PhoneNumber string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
}

type User struct {
	Verified bool `json:"verified"`
	Username string `json:"username"`
	NameString string `json:"name_string"`
	SearchString string `json:"search_string"`
	UUID string `json:"uuid"`
	ULID string `json:"ulid"`
	Barcodes []string `json:"barcodes"`
	EmailAddress string `json:"email_address"`
	PhoneNumber string `json:"phone_number"`
	Identity Person `json:"identity"`
	AuthorizedAliases []Person `json:"authorized_aliases"`
	FamilySize int `json:"family_size"`
	FamilyMembers []Person `json:"family_members"`
	CreatedDate string `json:"created_date"`
	CreatedTime string `json:"created_time"`
	CheckIns []CheckIn `json:"check_ins"`
	FailedCheckIns []CheckIn `json:"failed_check_ins"`
	TotalGuestsAdmitted int `json:"total_guests_admitted"`
	TimeRemaining int `json:"time_remaining"`
	AllowedToCheckIn bool `json:"allowed_to_checkin"`
	Language string `json:"language"`
	SimilarUsers []User `json:"similar_users"`
	ECB string `json:"ecb"`
}

func New() ( new_user User ) {
	now := utils.GetNowTimeOBJ()
	new_user.Username = ""
	new_user.Verified = false
	new_user.UUID = uuid.NewV4().String()
	new_user.FamilySize = 1
	new_user.CreatedDate = utils.GetNowDateString( &now )
	new_user.CreatedTime = utils.GetNowTimeString( &now )
	return
}
