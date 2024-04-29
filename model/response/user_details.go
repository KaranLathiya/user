package response

type UserID struct {
	UserID string `json:"userID"  validate:"required" `
}

type Username struct {
	Username string `json:"username,omitempty"  validate:"required" `
}

type UserDetails struct {
	UserID      string  `json:"userID" `
	Firstname   string  `json:"firstname" `
	Lastname    string  `json:"lastname" `
	Fullname    string  `json:"fullname" `
	Username    string  `json:"username" `
	Email       *string `json:"email,omitempty" `
	PhoneNumber *string `json:"phoneNumber,omitempty" `
	CountryCode *string `json:"countryCode,omitempty" `
	Privacy     string  `json:"privacy,omitempty" `
	CreatedAt   string  `json:"createdAt,omitempty" `
	UpdatedAt   *string `json:"updatedAt,omitempty" `
}

type User struct {
	UserID      string  `json:"userID" `
	Firstname   string  `json:"firstname" `
	Lastname    string  `json:"lastname" `
	Fullname    string  `json:"fullname" `
	Username    string  `json:"username" `
	Email       *string `json:"email" `
	PhoneNumber *string `json:"phoneNumber" `
	CountryCode *string `json:"countryCode" `
}
