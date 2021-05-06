package registery

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type dataBaseMap struct {
	dbases map[string]dataBaseConfig
}

type dataBaseConfig struct {
	db         *sql.DB
	startalpha byte
	endalpha   byte
}

//Establish DB connection for each of the database
func GetConnections() *dataBaseMap {

	var dbmap dataBaseMap
	var dbconfig dataBaseConfig
	dbmap.dbases = make(map[string]dataBaseConfig)

	dbdetaillist := getdbDetails()
	for _, dbdetail := range dbdetaillist {
		db := newMySqlDataConn(dbdetail.dbname, dbdetail.userid, dbdetail.password)
		dbconfig.db = db
		dbconfig.startalpha = dbdetail.startalpha
		dbconfig.endalpha = dbdetail.endalpha
		dbmap.dbases[dbdetail.dbname] = dbconfig
	}
	return &dbmap
}

func newMySqlDataConn(dbname string, userid string, password string) *sql.DB {

	db, err := sql.Open("mysql", userid+":"+password+"@/"+dbname+"?parseTime=true")
	if err != nil {
		log.Fatal("Error while connecting to database ", err)
		return nil
	}
	return db
}

//Get database name base on first alphabet of name
func (dbmap *dataBaseMap) getDataBase(name string) string {
	var s string
	for k, t := range dbmap.dbases {
		if name[0] >= t.startalpha && name[0] <= t.endalpha {
			return k
		}
	}
	return s
}
