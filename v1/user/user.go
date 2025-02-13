package user

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	server "github.com/0187773933/GO_SERVER/v1/server"
)

type CheckIn struct {
	Date string `json:"date"`
	Time string `json:"time"`
	UUID string `json:"uuid"`
	Additional []string `json:"additional"`
}

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
	PhoneNumber string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
	UUID string `json:"uuid"`
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
	AuthorizedAliases []string `json:"authorized_aliases"`
	FamilySize int `json:"family_size"`
	FamilyMembers []Person `json:"family_members"`
	Spouse Person `json:"spouse"`
	Children []Person `json:"children"`
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

func New( s *server.Server ) ( new_user User ) {
	now := s.LOG.GetNowTimeOBJ()
	new_user.Username = ""
	new_user.Verified = false
	new_user.UUID = uuid.NewV4().String()
	new_user.FamilySize = 1
	new_user.CreatedDate = s.LOG.FormatDateString( &now )
	new_user.CreatedTime = s.LOG.FormatTimeString( &now )
	return
}

func ( user *User ) GetFamilyName() string {
	if user.Spouse.FirstName != "" {
		return fmt.Sprintf("%s and %s %s", user.Identity.FirstName, user.Spouse.FirstName, user.Identity.LastName)
	}
	return fmt.Sprintf("%s %s", user.Identity.FirstName, user.Identity.LastName)
}