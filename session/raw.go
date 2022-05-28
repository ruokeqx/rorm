package session

import (
	"database/sql"
	"rorm/log"
	"strings"
	"rorm/dialect"
	"rorm/schema"
)

type Session struct {
	db      *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql     strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db: db,
		dialect: dialect,
	}
}

// session 复用
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

// 设置sql和args
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 设置好sql和args后执行
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow 获取一条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows 获取记录的列表
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
