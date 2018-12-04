package data

import (
	"database/sql"
	"testing"
)

func Test_UserCreate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	if users[0].Id == 0 {
		t.Errorf("No id or created_at in user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "User not created.")
	}
	if users[0].Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created.")
	}
	setup()
}

func Test_UserDelete(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	if err := users[0].Delete(); err != nil {
		t.Error(err, "- Cannot delete user.")
	}
	_, err := UserByEmail(users[0].Email)
	if err != sql.ErrNoRows {
		t.Error(err, "- User not deleted.")
	}
	setup()
}

func Test_UserUpdate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	users[0].Name = "Random User"
	if err := users[0].Update(); err != nil {
		t.Error(err, "- Cannot update user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "- User not updated.")
	}
	if u.Name != "Random User" {
		t.Errorf("- User not updated.")
	}
	setup()
}

func Test_UserByUuid(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user")
	}
	u, err := UserByUUID(users[0].Uuid)
	if err != nil {
		t.Error(err, "User not created.")
	}
	if users[0].Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created.")
	}
	setup()
}

func Test_Users(t *testing.T) {
	setup()
	for _, user := range users {
		if err := user.Create(); err != nil {
			t.Error(err, "Cannot create user.")
		}
	}
	u, err := Users()
	if err != nil {
		t.Error(err, "Cannot retrieve users.")
	}
	if len(u) != 2 {
		t.Error(err, "Wrong number of users retrieved.")
	}
	if u[0].Email != users[0].Email {
		t.Error(u[0], users[0], "Wrong user retrieved.")
	}
	setup()
}
