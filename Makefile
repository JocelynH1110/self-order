db.drop:
	rm -f db/dev.db

db.init: db.drop
	sqlite3 db/dev.db < db/init.sql
