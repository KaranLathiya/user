
package request

type BlockUser struct {
	BlockedUser   string `json:"blockedUser" validate:"required"`
}