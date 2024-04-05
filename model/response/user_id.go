package response 

type UserID struct{
	UserID string `json:"userID"  validate:"required" `
}