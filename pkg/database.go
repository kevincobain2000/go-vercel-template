package pkg

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	log.Info("db init")
	DB()
}

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func DB() *gorm.DB {
	dbOnce.Do(func() {
		log.Info("db once")
		db = getDB()
	})
	return db
}

func getDB() *gorm.DB {
	if db != nil {
		return db
	}
	db = openConnection(dsn())

	sqlDB := attachSQLInstance()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(30)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(30)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// Default DB close on mysql is 8 hours, so we set way before that (1 min)
	// This can be increased to 1 hour as well
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(1))

	attachSQLInstance()

	return db
}

func openConnection(dsn string) *gorm.DB {
	conn := mysql.Open(dsn)
	gc := &gorm.Config{
		Logger: logger.Default.LogMode(
			logger.LogLevel(Config().DBLogLevel),
		),
	}
	db, err := gorm.Open(conn, gc)
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	return db
}

func attachSQLInstance() *sql.DB {
	sqlDB, err := db.DB()
	_ = sqlDB

	if err != nil {
		log.Fatal("cannot get to sql database")
	}
	return sqlDB
}

func dsn() string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		Config().DBUsername,
		Config().DBPassword,
		Config().DBHost,
		Config().DBPort,
		Config().DBDatabase,
	)
	log.Info("dsn" + dsn)
	return dsn
}
