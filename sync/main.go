package main

import (
	"fmt"
	"os"

	"github.com/kainonly/cronx/bootstrap"
	"github.com/kainonly/cronx/common"
	"github.com/kainonly/cronx/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if err := sync("./config/values.yml"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sync(path string) (err error) {
	var v *common.Values
	if v, err = bootstrap.LoadStaticValues(path); err != nil {
		return
	}

	var db *gorm.DB
	if db, err = gorm.Open(
		sqlite.Open(v.Database.Path),
		&gorm.Config{},
	); err != nil {
		return
	}

	return db.AutoMigrate(model.Scheduler{})
}
