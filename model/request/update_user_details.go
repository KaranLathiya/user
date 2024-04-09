package request

type UpdateUserPrivacy struct {
	Privacy string `json:"Privacy" validate:"required|in:public,private"`
}


type UpdateUserNameDetails struct {
	Firstname   string `json:"firstname" validate:"required|min_len:2|max_len:50" `
	Lastname    string `json:"lastname" validate:"required|min_len:2|max_len:50" `
	Username	string `json:"username" validate:"required|min_len:2|max_len:50" `
}

