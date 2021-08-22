package txstore

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type TxDB struct {
	database sql.DB
}

func Create() TxDB {
	database, _ := sql.Open("sqlite3", "./.store.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS transactions (id INTEGER PRIMARY KEY, from_addr TEXT, to_addr TEXT, amount INTEGER)")
	statement.Exec()
	t := TxDB{database: *database}
	return t
}

func (t TxDB) Test() {
	statement, _ := t.database.Prepare(fmt.Sprintf("INSERT INTO transactions (from, to, amount) VALUES (%s, %s, %f)", "abc", "def", 0.001))
	statement.Exec()
	rows, _ := t.database.Query("SELECT id, from, to, amount FROM transactions")
	fmt.Println(rows)
}
