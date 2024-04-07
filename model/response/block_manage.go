package response

type BlockUserDetails struct{
	BlockedUser   string `json:"blockedUser" validate:"required"`
	BlockedAt string `json:"blockedAt" validate:"required"`
}