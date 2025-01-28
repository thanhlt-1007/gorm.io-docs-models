package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "time"
)

type User struct {
    ID uint             // Standard field for the primary key
    Name string         // A regular string field
    Email *string       // A pointer to a string, allowing for null values
    Age uint8           // A unsigned 8-bit integer
    Birthday *time.Time // A pointer to time.Time, can be null
    CreatedAt time.Time // Automatically managed by GORM for creation time
    UpdatedAt time.Time // Automatically managed by GORM for update time
    ignored string      // fields that aren't exported are ignored
}

func main() {
    db, err := gorm.Open(
        sqlite.Open("db.sqlite"),
        &gorm.Config{},
    )

    if err != nil {
        fmt.Printf("gorm.Open error %v\n", err)
        panic(err)
    }
    fmt.Println("gorm.Open success")

    db.AutoMigrate(&User{})
    fmt.Println("db.AutoMigrate User success")
}
