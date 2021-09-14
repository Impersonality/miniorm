package schema

import (
	"miniorm/geeorm/dialect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDialect, _ = dialect.GetDialect("mysql")

func TestParse(t *testing.T) {

	schema := Parse(&User{}, TestDialect)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
