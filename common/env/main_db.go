package env


type MainDB struct{
	Link string `json:"link"`
	Name string `json:"name"`
	MaxConn int `json:"max_conn"`
}
