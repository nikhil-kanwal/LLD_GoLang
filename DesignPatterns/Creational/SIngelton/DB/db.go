package Db

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB
}

var lock = &sync.Mutex{}

var dbInstance *DB

// NewDB initializes a new database connection using GORM and PostgreSQL.
func GetDbInstance(dsn string) (*DB, error) {
	// Open the database connection
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		//  Need to check this twice because if let say two threads are waiting here
		//  to acquire lock then first thread may have created the connection and the second will create again
		if dbInstance == nil {
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				fmt.Printf("Failed to connect to the database: %v \n", err)
				return nil, err
			}
			fmt.Print("Single instance created.\n")
			dbInstance = &DB{Conn: db}
		} else {
			fmt.Print("Single instance already created.\n")
		}
	} else {
		fmt.Print("Single instance already created.\n")
	}
	// Return the DB instance
	return dbInstance, nil
}
