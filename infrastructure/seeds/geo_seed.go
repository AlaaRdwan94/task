package seeds

import (
	"github.com/InnoSoft/task/entity"
	"github.com/InnoSoft/task/infrastructure/db"
	"strconv"
	"time"
)

func Seed() {
	var newDB db.Database
	newDB = db.NewPostgres()
	db := newDB.Open()
	defer db.Close()
	user :=entity.User{
		FirstName: "test1",
		LastName:  "test2",
		Email:     strconv.Itoa(time.Now().UTC().Minute())+ "alaa.a.radwan1994@gmail.com",
		PassWord:  "123",
		Phone:     "012457485",
	}
	db.GormDB.Create(&user)
}
