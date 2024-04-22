package adapter

import (
	"time"

	log "github.com/rybakdigital/utility-library-go/logging/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var logger *log.Logger

type Adapter struct {
	IsConnected bool
	Logger      *log.Logger
	Config      *Config
	db          *gorm.DB
}

func NewAdapter(c *Config, l *log.Logger) *Adapter {
	// Assign logger
	l.Printf("Created new MySQL adapter")

	return &Adapter{Config: c, Logger: l, IsConnected: false}
}

func DefaultAdapter() *Adapter {
	// Create logger
	log := log.NewLogger("mysql-adapter")

	return NewAdapter(DefaultConfig(), log)
}

func (a *Adapter) GetDb() *gorm.DB {
	return a.db
}

func (a *Adapter) Connect() error {
	for i := 1; i <= a.Config.GetMaxAttempts(); i++ {
		a.Logger.Printf("Attempt %d: Attempting to establish db connection", i)
		db, err := a.IsDbConnectionAvailable()

		if err == nil {
			a.IsConnected = true
			a.db = db
			a.Logger.Printf("Attempt %d: Connection to db successfully established", i)
			break
		}

		a.Logger.Fatalf("Attempt %d: Failed to establish connection to db", i)

		if i < a.Config.GetMaxAttempts() {
			a.Logger.Printf("Will retry connecting in %d seconds...", a.Config.GetInterval()/time.Second)
		} else {
			a.Logger.Fatalf("Could not establish connection to db after %d attempts", i)
			return err
		}

		time.Sleep(a.Config.GetInterval())
	}

	return nil
}

func (a *Adapter) IsDbConnectionAvailable() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(a.Config.Dsn.Dsn()), &gorm.Config{})

	if err != nil {
		a.Logger.Printf("Could not connect to db %v", err)

		return db, err
	}

	return db, nil
}
