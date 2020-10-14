package env

type Rpc struct {
	AuthServerPort  int `json:"auth_server_port"`
	EmailServerPort int `json:"email_server_port"`
	MvpServerPort int `json:"mvp_server_port"`
}
