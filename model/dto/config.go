package dto

type Config struct {
	Port     string   `mapstructure:"PORT"`
	Database Database `mapstructure:",squash"`
	SMTP     SMTP     `mapstructure:",squash"`
	Twilio   Twilio   `mapstructure:",squash"`
	GooglAuth GoogleAuth `mapstructure:",squash"`
}
type Database struct {
	DBName string `mapstructure:"DATABASE_NAME"`
	DBUser string `mapstructure:"DATABASE_USER"`
}
type SMTP struct {
	EmailFrom     string `mapstructure:"SMTP_EMAIL_FROM"`
	EmailPassword string `mapstructure:"SMTP_EMAIL_PASSWORD"`
	Host          string `mapstructure:"SMTP_HOST"`
	Port          string `mapstructure:"SMTP_PORT"`
}

type Twilio struct {
	AcountSID   string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken   string `mapstructure:"TWILIO_AUTHTOKEN"`
	MessageFrom string `mapstructure:"TWILIO_MESSAGE_FROM"`
}


type GoogleAuth struct {
	ClientSecret   string `mapstructure:"GOOGLE_AUTH_CLIENT_SECRET"`
	ClientID   string `mapstructure:"GOOGLE_AUTH_CLIENT_ID"`
	RedirectURL string `mapstructure:"GOOGLE_AUTH_REDIRECT_URL"`
}