package migrate

import (
	"go-im/app/appInit"
	"go-im/app/model"
)

func CreateTable() {
	appInit.DB.AutoMigrate(
		//!!do not delete the line, gen generate code at here
		&model.User{},
	)
}
