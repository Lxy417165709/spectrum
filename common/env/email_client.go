package env

type EmailClient struct {
	EmailAddr string `json:"email_addr"`
	AuthCode  string `json:"auth_code"`
	SmtpAddr  string `json:"smtp_addr"`
	SmtpPort  int    `json:"smtp_port"`
}
