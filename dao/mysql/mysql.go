package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"qishuBook/setting"
)

var db *sqlx.DB

func InitMysql(cfg *setting.MysqlConfig) (err error) {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err = sqlx.Connect("mysql", dsn) // = sql.open + sql.ping
	if err != nil {
		fmt.Printf("[error] find a err in connect DB:%v\n", err)
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenCons)
	db.SetMaxIdleConns(cfg.MaxIdleCons)
	return nil
}

func Close() {
	_ = db.Close()
}
