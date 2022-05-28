package libs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
	logger "github.com/sirupsen/logrus"
)

var (
	Config = struct {
		LogLevel string `default:"info" env:"Loglevel"`
		Bindata  bool   `default:"true" env:"Bindata"`
		Debug    bool   `default:"true" env:"Debug"`
		HTTPS    bool   `default:"false" env:"HTTPS"`
		Certpath string `default:"" env:"Certpath"`
		Certkey  string `default:"" env:"Certkey"`
		Port     int    `default:"8085" env:"PORT"`
		Host     string `default:"0.0.0.0" env:"Host"`
		Admin    struct {
			UserName        string `env:"AdminUserName" default:"username"`
			Name            string `env:"AdminName" default:"name"`
			Pwd             string `env:"AdminPwd" default:"123456"`
			RoleName        string `env:"AdminRoleName" default:"admin"`
			RoleDisplayName string `env:"TenantRoleDisplayName" default:"admin"`
		}

		DB
	}{}
)

type DB struct {
	Prefix string `env:"DBPrefix" default:""`

	Name     string `env:"DBName" default:"MedicalQuestionnaireService"`
	Adapter  string `env:"DBAdapter" default:"mysql"`
	Host     string `env:"DBHost" default:"medicalquestionnaireservice.cz7nrjgqkgdb.ap-northeast-1.rds.amazonaws.com"`
	Port     string `env:"DBPort" default:"3306"`
	User     string `env:"DBUser" default:""`
	Password string `env:"DBPassword" default:"password"`
}

func InitConfig(p, m string) {

	path, _ := os.Getwd()

	configPath := filepath.Join(path, "application.yaml")
	if p != "" {
		configPath = p
	}

	if err := configor.Load(&Config, configPath); err != nil {
		logger.Println(fmt.Sprintf("Config Path:%s ,Error:%s", configPath, err.Error()))
		return
	}

	// 测试环境打包指令
	// go build -ldflags "-X main.Model=Test"
	if len(m) == 0 {
		m = "Debug"
	}
	switch m {
	case "Test":
		Config.Host = "0.0.0.0"
		Config.Port = 8085
	case "Produce":
		Config.Host = "0.0.0.0"
		Config.Port = 8085
	default:
		Config.Host = "0.0.0.0"
		Config.Port = 8085
	}

	if Config.Debug {
		fmt.Println(fmt.Sprintf("config項目：%v", Config))
	}
}
