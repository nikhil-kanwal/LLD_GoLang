package main

import (
	"fmt"

	db "github.com/nikhil-kanwal/LLD_GoLang/DesignPatterns/Creational/SIngelton/DB"
)

func main() {

	fmt.Print("initializing DB\n")
	dsn := "user=postgres password=postgres dbname=user_module port=5432 sslmode=disable"

	// Initialize the DB instance
	for i := 0; i < 10; i++ {
		go func() {
			_, err := db.GetDbInstance(dsn)
			if err != nil {
				fmt.Printf("error initializing database: %+v\n", err)
			}
		}()

	}
	fmt.Print("Database connection established successfully\n")
	fmt.Scanln()

}
