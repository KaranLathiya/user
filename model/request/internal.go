package request

type UserIDs struct {
	UserIDs []string `json:"userIDs" validate:"required"`
}

type CreateOTPForDeleteOrganization struct {
	OrganizationID string `json:"organizationID" validate:"required"`
	OwnerID        string `json:"ownerID" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type UserID struct {
	UserID string `json:"userID" validate:"required"`
}