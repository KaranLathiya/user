package constant

const (
	OTP_SENT             = "Successfully OTP sent."
	USER_DETAILS_UPDATED = "User details successfully updated."
	UNBLOCKED_USER       = "Successfully unblocked user"
	BLOCKED_USER         = "Successfully blocked user"
	ORGANIZATION_DELETED = "Organization deleted successfully."

	GOOGLE_ACCESS_TOKEN_REQUEST_URL = "https://oauth2.googleapis.com/token"
	GOOGLE_INFO_REQUEST_URL         = "https://www.googleapis.com/oauth2/v3/userinfo?access_token="
	ORGANIZATION_SERVICE_BASE_URL   = "http://localhost:9000/"

	EVENT_TYPE_SIGNUP              = "signup"
	EVENT_TYPE_LOGIN               = "login"
	EVENT_TYPE_GOOGLE_LOGIN        = "google_login"
	EVENT_TYPE_ORGANIZATION_DELETE = "organization_delete"

	SIGNUP_MODE_EMAIL        = "email"
	SIGNUP_MODE_GOOGLE_LOGIN = "google_login"
	SIGNUP_MODE_PHONE_NUMBER = "phone_number"
	
	LOGIN_TYPE_EMAIL        = "email"
	LOGIN_TYPE_PHONE_NUMBER = "phone_number"
)
