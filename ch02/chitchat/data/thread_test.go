package data

import (
	"database/sql"
	"fmt"
	"testing"
	"log"
	"os"
)

func TestMain(m *testing.M) {
	fmt.Println("start test")
        var err error
        Db, err = sql.Open("postgres", "user=gwp dbname=chitchat password=gwp sslmode=disable")
        if err != nil {
                log.Fatal(err)
        }
        
	//m.Run()
	code := m.Run()

	Db.Close()
	fmt.Println("end test")
	os.Exit(code)
}

func ThreadDeleteAll() (err error) {
	//db := db()
	statement := "delete from threads"
	_, err = Db.Exec(statement)
	if err != nil {
		return
	}
	return
}

func Test_CreateThread(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}
	if conv.UserId != users[0].Id {
		t.Error("User not linked with thread")
	}
	setup()
}
