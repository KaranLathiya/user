package response

type BlockUserDetails struct {
	ID          string `json:"id" validate:"required"`
	BlockedUser string `json:"blockedUser" validate:"required"`
	BlockedAt   string `json:"blockedAt" validate:"required"`
	Fullname    string `json:"fullname" validate:"required"`
}
