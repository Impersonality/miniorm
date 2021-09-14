package session

import (
	"fmt"
	"miniorm/geeorm/log"
	"miniorm/geeorm/schema"
	"reflect"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	if value == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("create table %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("drop table if exists %s", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExitedSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var temp string
	_ = row.Scan(&temp)
	return temp == s.RefTable().Name
}
