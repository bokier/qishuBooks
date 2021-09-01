package mysql

func InsertInfo(author, book_name, time string, size float64, level int64) (err error) {
	sqlStr := `insert into bookinfo(author,book_name,size,level,time) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, author, book_name, size, level, time)
	return
}
