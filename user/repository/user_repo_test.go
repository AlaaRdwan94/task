package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/InnoSoft/task/entity"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open("postgres", db)
	userRepoobj := NewUserRepo(gdb, nil)
	user :=entity.User{
		FirstName: "test1",
		LastName:  "test2",
		Email:     "alaa.a.radwan1994@gmail.com",
		PassWord:  "123456",
		Phone:     "012457485",
	}
	mock.ExpectQuery(
		"SELECT(.*)").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "first_name", "email","last_name","pass_word","phone"}).
				AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.DeletedAt, user.FirstName, user.Email,user.LastName,user.PassWord,user.Phone))
	res, err := userRepoobj.GetUser(1)

	require.NoError(t, err)
	require.Equal(t, res, &user)

}


func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open("postgres", db)
	userRepoobj := NewUserRepo(gdb, nil)
	user :=entity.User{
		FirstName: "test1",
		LastName:  "test2",
		Email:     "alaa.a.radwan1994@gmail.com",
		PassWord:  "123456",
		Phone:     "012457485",
	}
	mock.ExpectQuery(
		"SELECT(.*)").
		WithArgs(user.Email).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "first_name", "email","last_name","pass_word","phone"}).
				AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.DeletedAt, user.FirstName, user.Email,user.LastName,user.PassWord,user.Phone))
	res, err := userRepoobj.GetUserByEmail(user.Email)

	require.NoError(t, err)
	require.Equal(t, res, &user)
}
