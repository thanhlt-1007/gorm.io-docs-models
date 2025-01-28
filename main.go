package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func main() {
    _, err := gorm.Open(
        sqlite.Open("db.sqlite"),
        &gorm.Config{},
    )

    if err != nil {
        panic(err)
    }}
