package pkg

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	currentPath, _ = os.Getwd()
	StoragePath    = filepath.Join(currentPath, "storage")
	ErrorLogPath   = filepath.Join(StoragePath, "errorlog")
	LogPath        = filepath.Join(StoragePath, "log")
	YummyPath      = filepath.Join(currentPath, "yummy")
	RegisterPath   = filepath.Join(currentPath, "script", "order", "register")
	ReservedPath   = filepath.Join(RegisterPath, "reserved")

	WebPath   = filepath.Join(currentPath, "web")
	IndexPath = filepath.Join(WebPath, "*.html")
)
var (
	backendHost     = "backend"
	MobilePort      string
	PCPort          string
	DeliveryPort    string
	MonitorPort     string
	DeliveryAddress string

	TZ string

	MySQLDSN    string
	PostgresDSN string

	MongoDatabase  string
	MongoDSN       string
	AssembleNumber int
)

func init() {
	err := godotenv.Load(filepath.Join(currentPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AssembleNumber, err = strconv.Atoi(os.Getenv("ASSEMBLE_NUMBER"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TZ = os.Getenv("TZ")
	MobilePort = os.Getenv("MOBILE_BACK_PORT")
	PCPort = os.Getenv("PC_BACK_PORT")
	DeliveryPort = os.Getenv("DELIVERY_BACK_PORT")
	MonitorPort = os.Getenv("MONITOR_BACK_PORT")
	DeliveryAddress = backendHost + DeliveryPort

	PostgresDSN = "host=postgres" +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" port=" + os.Getenv("POSTGRES_BACK_PORT") +
		" TimeZone=" + TZ +
		" dbname=app sslmode=disable"

	MySQLDSN = os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") +
		"@tcp(mysql)/" + os.Getenv("MYSQL_DATABASE") +
		"?charset=utf8&parseTime=True&loc=Local"

	MongoDatabase = os.Getenv("MONGO_INITDB_DATABASE")
	MongoDSN = "mongodb://" +
		os.Getenv("MONGO_INITDB_ROOT_USERNAME") +
		":" +
		os.Getenv("MONGO_INITDB_ROOT_PASSWORD") +
		"@mongo:" +
		os.Getenv("MONG_BACK_PORT")

}
