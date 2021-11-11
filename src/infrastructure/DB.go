package infrastructure

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

type DB struct {
		DBms string
    DBuser string
    DBpass string
    DBprotocol string
    DBname string
    Connection *gorm.DB
}

func NewDB() *DB {
    c := NewConfig()
    return newDB(&DB{
        DBms: c.DB.Production.DBms,
        DBuser: c.DB.Production.DBuser,
        DBpass: c.DB.Production.DBpass,
        DBprotocol: c.DB.Production.DBprotocol,
        DBname: c.DB.Production.DBname,
    })
}

func newDB(d *DB) *DB {
    //DBへ接続
    c := d.DBuser + ":" + d.DBpass + "@" + d.DBprotocol + "/" + d.DBname
    fmt.Println(c)
    db, err := gorm.Open(d.DBms, c)
    if err != nil {
     fmt.Println("db unconnected ")
     return d
    }

    d.Connection = db
    // DBエンジンを「InnoDB」に設定
    d.Connection.Set("gorm:table_options", "ENGINE=InnoDB")
    // 詳細なログを表示
    d.Connection.LogMode(true)
    // 登録するテーブル名を単数形にする（デフォルトは複数形）
    d.Connection.SingularTable(true)
    fmt.Println("db connected: ", &d.Connection)

    return d
}

// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
    return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
    return db.Connection
}
