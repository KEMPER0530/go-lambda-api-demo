package infrastructure

import (
  // フォーマットI/O
  "os"
)

type Config struct {
    DB struct {
        Production struct {
            DBms string
            DBuser string
            DBpass string
            DBprotocol string
            DBname string
        }
    }
    Routing struct {
        Port string
    }
}

func NewConfig() *Config {

    c := new(Config)

    // 環境変数から値を取得
    c.DB.Production.DBms = os.Getenv("DBMS")
    c.DB.Production.DBuser = os.Getenv("DB_USER")
    c.DB.Production.DBpass = os.Getenv("DB_PASS")
    c.DB.Production.DBprotocol = os.Getenv("DB_PROTOCOL")
    c.DB.Production.DBname = os.Getenv("DB_NAME")

    return c
}
