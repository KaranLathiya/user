package request

type UserListParameter struct {
	Fullname    string `json:"fullname" `
	Email       *string `json:"email" `
	PhoneNumber *string `json:"phoneNumber" `
	Limit       int    `json:"limit" `
	Offset      int    `json:"offset" `
	Order       string `json:"order" validate:"in:asc,desc"`
	OrderBy     string `json:"orderBy" validate:"in:fullname,email,phone,date"`
}


type UpdateUserPrivacy struct {
	Privacy string `json:"Privacy" validate:"required|in:public,private"`
}


type UpdateUserNameDetails struct {
	Firstname   string `json:"firstname" validate:"required|min_len:2|max_len:50" `
	Lastname    string `json:"lastname" validate:"required|min_len:2|max_len:50" `
	Username	string `json:"username" validate:"required|min_len:2|max_len:50" `
}
