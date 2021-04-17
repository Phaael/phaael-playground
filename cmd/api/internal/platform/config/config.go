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
	DBHost = "localhost"
	DbUserName = "root"
	PassDB = "q1w2e3r4"
	DBPort = "3306"
}
