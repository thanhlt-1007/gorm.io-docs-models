package main

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-models/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

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

    db.AutoMigrate(&models.User{})
    fmt.Println("db.AutoMigrate models.User success")
}
