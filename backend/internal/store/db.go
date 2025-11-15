package store
import(
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// initialize db connection
func InitDB(){
	// localhost for now
	host:="localhost"
	port:=5432
	user:="chat_app"
	password:="chat_app"
	dbname:="chat_app"

	dsn:=fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		host,user,password,dbname,port,
	)

	// gorm logger 
	newLogger:=logger.New(
		log.New(os.Stdout,"\r\n",log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful: true,
		},
	)
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err!=nil{
		log.Fatalf("failed to connect database: %v",err)

	}
	DB = db
}