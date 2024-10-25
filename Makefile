db.drop:
	rm -f db/dev.db

# < 運算子，將 init.sql 當作指令的標準輸入。
db.init: db.drop
	sqlite3 db/dev.db < db/init.sql
