package env

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)


var Conf = &conf{}

type conf struct {
	//Mysql       Mysql       `json:"mysql"`
	//Http        Http        `json:"http"`
	EmailClient EmailClient `json:"email_client"`
	//Cache       Cache       `json:"cache"`
	//Vrc         Vrc         `json:"vrc"`
	//Salt        Salt        `json:"salt"`
	//Template    Template    `json:"template"`
	//SecretKey   SecretKey   `json:"secret_key"`
	//FileStorage FileStorage `json:"file_storage"`
}

func (c *conf) Load(pathOfConfFile string) error {
	confFileData, err := ioutil.ReadFile(pathOfConfFile)
	if err != nil {
		logs.Error(err)
		return err
	}
	if err := json.Unmarshal(confFileData, c); err != nil {
		logs.Error(err)
		return err
	}
	logs.Debug("%+v", c)
	return nil
}
