package main

import (
	"mkdb/models"
)

func main() {
	models.CreateDB()
	models.WriteData2DB()
	defer models.DB.Close()
}
