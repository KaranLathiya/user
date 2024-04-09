package request

type UserListParameter struct {
	Fullname    string `json:"fullname" `
	Email       string `json:"email" `
	PhoneNumber string `json:"phoneNumber" `
	Limit       int    `json:"limit" `
	Offset      int    `json:"offset" `
	Order       string `json:"order" validate:"in:asc,desc"`
	OrderBy     string `json:"orderBy" validate:"in:fullname,email,phone,date"`
}
