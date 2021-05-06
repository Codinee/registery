package registery

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type dataBaseDetail struct {
	dbname     string
	userid     string
	password   string
	startalpha byte
	endalpha   byte
}

func getdbDetails() []dataBaseDetail {

	fi, err := os.Open("config.txt")
	if err != nil {
		log.Println("Error in Config file open", err)
	}
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	var fileData []string
	for scanner.Scan() {
		fileData = append(fileData, scanner.Text())
	}
	var dbDetailList []dataBaseDetail
	var dbdetail dataBaseDetail
	for _, data := range fileData {
		tmp := strings.Split(data, "/")
		dbdetail.dbname = tmp[0]
		dbdetail.userid = tmp[1]
		dbdetail.password = tmp[2]
		dbdetail.startalpha = byte(tmp[3][0])
		dbdetail.endalpha = byte(tmp[4][0])
		dbDetailList = append(dbDetailList, dbdetail)
	}
	return dbDetailList
}
