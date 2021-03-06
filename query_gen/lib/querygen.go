/* WIP Under Construction */
package qgen

import (
	"database/sql"
	"errors"
)

var DB_Registry []DB_Adapter
var ErrNoAdapter = errors.New("This adapter doesn't exist")

type DB_Table_Column struct {
	Name           string
	Type           string
	Size           int
	Null           bool
	Auto_Increment bool
	Default        string
}

type DB_Table_Key struct {
	Columns string
	Type    string
}

type DB_Select struct {
	Table   string
	Columns string
	Where   string
	Orderby string
	Limit   string
}

type DB_Join struct {
	Table1  string
	Table2  string
	Columns string
	Joiners string
	Where   string
	Orderby string
	Limit   string
}

type DB_Insert struct {
	Table   string
	Columns string
	Fields  string
}

type DB_Column struct {
	Table string
	Left  string // Could be a function or a column, so I'm naming this Left
	Alias string // aka AS Blah, if it's present
	Type  string // function or column
}

type DB_Field struct {
	Name string
	Type string
}

type DB_Where struct {
	Expr []DB_Token // Simple expressions, the innards of functions are opaque for now.
}

type DB_Joiner struct {
	LeftTable   string
	LeftColumn  string
	RightTable  string
	RightColumn string
	Operator    string
}

type DB_Order struct {
	Column string
	Order  string
}

type DB_Token struct {
	Contents string
	Type     string // function, operator, column, number, string, substitute
}

type DB_Setter struct {
	Column string
	Expr   []DB_Token // Simple expressions, the innards of functions are opaque for now.
}

type DB_Limit struct {
	Offset   string // ? or int
	MaxCount string // ? or int
}

type DB_Stmt struct {
	Contents string
	Type     string // create-table, insert, update, delete
}

type DB_Adapter interface {
	GetName() string
	CreateTable(name string, table string, charset string, collation string, columns []DB_Table_Column, keys []DB_Table_Key) (string, error)
	SimpleInsert(name string, table string, columns string, fields string) (string, error)

	// ! DEPRECATED
	//SimpleReplace(name string, table string, columns string, fields string) (string, error)
	// ! NOTE: MySQL doesn't support upserts properly, so I'm removing this from the interface until we find a way to patch it in
	//SimpleUpsert(name string, table string, columns string, fields string, where string) (string, error)
	SimpleUpdate(name string, table string, set string, where string) (string, error)
	SimpleDelete(name string, table string, where string) (string, error)
	Purge(name string, table string) (string, error)
	SimpleSelect(name string, table string, columns string, where string, orderby string, limit string) (string, error)
	SimpleLeftJoin(name string, table1 string, table2 string, columns string, joiners string, where string, orderby string, limit string) (string, error)
	SimpleInnerJoin(string, string, string, string, string, string, string, string) (string, error)
	SimpleInsertSelect(string, DB_Insert, DB_Select) (string, error)
	SimpleInsertLeftJoin(string, DB_Insert, DB_Join) (string, error)
	SimpleInsertInnerJoin(string, DB_Insert, DB_Join) (string, error)
	SimpleCount(string, string, string, string) (string, error)
	Write() error

	// TODO: Add a simple query builder
}

func GetAdapter(name string) (adap DB_Adapter, err error) {
	for _, adapter := range DB_Registry {
		if adapter.GetName() == name {
			return adapter, nil
		}
	}
	return adap, ErrNoAdapter
}

type QueryPlugin interface {
	Hook(name string, args ...interface{}) error
	Write() error
}

type MySQLUpsertCallback struct {
	stmt *sql.Stmt
}

func (double *MySQLUpsertCallback) Exec(args ...interface{}) (res sql.Result, err error) {
	if len(args) < 2 {
		return res, errors.New("Need two or more arguments")
	}
	args = args[:len(args)-1]
	return double.stmt.Exec(append(args, args...)...)
}

func PrepareMySQLUpsertCallback(db *sql.DB, query string) (*MySQLUpsertCallback, error) {
	stmt, err := db.Prepare(query)
	return &MySQLUpsertCallback{stmt}, err
}
