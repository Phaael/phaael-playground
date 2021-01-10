package config

var (
	DataBaseName string
	DBHost       string
	DbUserName   string
	PassDB       string
	DBPort       string
)

func init() {
	DataBaseName = "appdb"
	DBHost = "mysql"
	DbUserName = "phaael"
	PassDB = "123456"
	DBPort = "3306"
}
