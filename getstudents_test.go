package registery

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func initialvalues2() (string, string, string, string, uint64) {
	name1 := "Cab"
	name2 := "Zl w"
	coursecode := "Cour"
	coursename := "Testing Sample course"
	var mobilenum uint64
	mobilenum = 1212121267
	return name1, name2, coursecode, coursename, mobilenum
}

//Get students in Course where course is in db1
func TestGetStudentsInCourse1(t *testing.T) {
	index = 100
	name1, _, coursecode, coursename, mobilenum := initialvalues2()

	db, _ := sql.Open("mysql", "root:Bond007@/db001?parseTime=true")
	dbmap := GetConnections()

	ccode := coursecode + strconv.Itoa(index)
	cname := coursename + strconv.Itoa(index)
	var students []Student

	db.Exec("Insert into courses(course_code,name)values(?,?)", ccode, cname)
	for k := 0; k < 3; k++ {
		sName := name1 + strconv.Itoa(k)
		var student Student
		student.id = index
		student.name = sName
		student.mobileNum = mobilenum
		db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)
		students = append(students, student)
		db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", student.id, ccode, time.Now())
		index += 1
	}
	o := dbmap.GetStudentsInCourse(ccode)
	if !reflect.DeepEqual(o, students) {
		log.Println("Result does not match in Get students TC1")
		t.Fail()
	}
}

//Get students in Course where course is in db2
func TestGetStudentsInCourse2(t *testing.T) {
	index += 1
	_, name2, coursecode, coursename, mobilenum := initialvalues2()

	db, _ := sql.Open("mysql", "root:Bond007@/db002?parseTime=true")
	dbmap := GetConnections()

	ccode := coursecode + strconv.Itoa(index)
	cname := coursename + strconv.Itoa(index)
	var students []Student
	db.Exec("Insert into courses(course_code,name)values(?,?)", ccode, cname)
	for k := 0; k < 3; k++ {
		sName := name2 + strconv.Itoa(k)
		var student Student
		student.id = index
		student.name = sName
		student.mobileNum = mobilenum
		db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)
		students = append(students, student)
		db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", student.id, ccode, time.Now())
		index += 1
	}
	o := dbmap.GetStudentsInCourse(ccode)
	if !reflect.DeepEqual(o, students) {
		log.Println("Result does not match in Get students TC2")
		t.Fail()
	}
}

//Get students in Course where course is in db1 and db2
func TestGetStudentsInCourse3(t *testing.T) {
	index += 1
	name1, name2, coursecode, coursename, mobilenum := initialvalues2()

	db, _ := sql.Open("mysql", "root:Bond007@/db001?parseTime=true")
	dbmap := GetConnections()

	ccode := coursecode + strconv.Itoa(index)
	cname := coursename + strconv.Itoa(index)
	var students []Student
	db.Exec("Insert into courses(course_code,name)values(?,?)", ccode, cname)
	for k := 0; k < 3; k++ {
		sName := name1 + strconv.Itoa(k)
		var student Student
		student.id = index
		student.name = sName
		student.mobileNum = mobilenum
		db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)
		students = append(students, student)
		db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", student.id, ccode, time.Now())
		index += 1
	}

	db, _ = sql.Open("mysql", "root:Bond007@/db002?parseTime=true")

	db.Exec("Insert into courses(course_code,name)values(?,?)", ccode, cname)
	for k := 0; k < 3; k++ {
		sName := name2 + strconv.Itoa(index)
		var student Student
		student.id = index
		student.name = sName
		student.mobileNum = mobilenum
		db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)
		students = append(students, student)
		db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", student.id, ccode, time.Now())
		index += 1
	}
	o := dbmap.GetStudentsInCourse(ccode)
	if !reflect.DeepEqual(o, students) {
		log.Println("Result does not match in Get students TC3")
		t.Fail()
	}
}

//Get students in Course where course is not in system
func TestGetStudentsInCourse4(t *testing.T) {
	index += 1
	_, name2, coursecode, _, mobilenum := initialvalues2()

	db, _ := sql.Open("mysql", "root:Bond007@/db002?parseTime=true")
	dbmap := GetConnections()

	ccode := coursecode + strconv.Itoa(index)
	//db.Exec("Insert into courses(course_code,name)values(?,?),ccode,cname)

	var students []Student
	sName := name2 + strconv.Itoa(index)
	var student Student
	student.id = index
	student.name = sName
	student.mobileNum = mobilenum
	db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)
	_, err := db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", student.id, coursecode, time.Now())
	if err != nil {
		fmt.Println("error while inserting into Students table", err)
	}
	index += 1
	o := dbmap.GetStudentsInCourse(ccode)
	if !reflect.DeepEqual(o, students) {
		log.Println("Result does not match in Get students TC4")
		t.Fail()
	}
}
