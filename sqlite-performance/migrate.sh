export DATABASE_URL="sqlite:test.db"
dbmate create
dbmate --migrations-dir ./schema up
