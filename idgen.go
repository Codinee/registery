package registery

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

//Generate new student id by reading config file
func getNewID() int {

	fi, err := os.OpenFile("static.txt", os.O_RDWR, 0777)
	if err != nil {
		log.Println("Error in file open", err)
	}

	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	var data string
	if scanner.Scan() {
		data = scanner.Text()
	}
	i, err := strconv.Atoi(data)
	if err != nil {
		log.Println("Error in converting file data", err)
	}
	i += 1
	_, err = fi.WriteAt([]byte(strconv.Itoa(i)), 0)
	if err != nil {
		log.Println("Error in Write file", err)
	}
	fi.Close()
	return i
}
