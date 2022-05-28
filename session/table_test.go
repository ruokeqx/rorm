package session

import (
	"database/sql"
	"testing"
	"rorm/dialect"
	_ "github.com/mattn/go-sqlite3"
)

var (
	testDb,_ = sql.Open("sqlite3","rorm.db")
	testDia,_ = dialect.GetDialect("sqlite3")
)

type User struct {
	Name string `rorm:"PRIMARY KEY"`
	Age  int
}

func NewSession() *Session{
	return New(testDb,testDia);
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	s.Raw("drop table if exists user;").Exec()
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}