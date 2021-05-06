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

var index int

func initialvalues1() (string, string, string, string, uint64) {
	name1 := "Aab"
	name2 := "Opn"
	coursecode := "Crs"
	coursename := "Testing course"
	var mobilenum uint64
	mobilenum = 1212121212
	return name1, name2, coursecode, coursename, mobilenum
}

//Get student courses where student is in DB1
func TestGetCourses1(t *testing.T) {

	index = 50
	name1, _, coursecode, coursename, mobilenum := initialvalues1()

	db, _ := sql.Open("mysql", "root:Bond007@/db001?parseTime=true")
	dbmap := GetConnections()

	var courses []Course
	sName := name1 + strconv.Itoa(index)
	student := &Student{id: index, name: sName, mobileNum: mobilenum}
	db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)

	for k := 0; k < 3; k++ {
		ccode := coursecode + strconv.Itoa(k)
		cname := coursename + strconv.Itoa(k)
		var course Course
		course.courseCode = ccode
		course.courseName = cname
		courses = append(courses, course)

		db.Exec("Insert into courses(course_code,name)values(?,?)", course.courseCode, course.courseName)
		db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", index, course.courseCode, time.Now())
	}
	o := dbmap.GetCourses(sName)
	if !reflect.DeepEqual(o, courses) {
		log.Println("Results does not match in Get Courses TC1")
		t.Fail()
	}
}

//Get student courses where student is in DB2
func TestGetCourses2(t *testing.T) {
	index += 1
	_, name2, coursecode, coursename, mobilenum := initialvalues1()

	db, _ := sql.Open("mysql", "root:Bond007@/db002?parseTime=true")
	dbmap := GetConnections()

	var courses []Course
	sName := name2 + strconv.Itoa(index)
	student := &Student{id: index, name: sName, mobileNum: mobilenum}
	db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)
	for k := 0; k < 3; k++ {
		ccode := coursecode + strconv.Itoa(k)
		cname := coursename + strconv.Itoa(k)
		var course Course
		course.courseCode = ccode
		course.courseName = cname
		courses = append(courses, course)
		db.Exec("Insert into courses(course_code,name)values(?,?)", course.courseCode, course.courseName)
		db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", index, course.courseCode, time.Now())
	}
	o := dbmap.GetCourses(sName)
	if !reflect.DeepEqual(o, courses) {
		log.Println("Results does not match in Get Courses TC2")
		t.Fail()
	}
}

//Get student courses where student not in system
func TestGetCourses3(t *testing.T) {
	index += 1
	_, name2, coursecode, coursename, _ := initialvalues1()

	db, _ := sql.Open("mysql", "root:Bond007@/db002?parseTime=true")
	dbmap := GetConnections()

	var courses []Course
	for k := 0; k < 3; k++ {
		ccode := coursecode + strconv.Itoa(k)
		cname := coursename + strconv.Itoa(k)
		var course Course
		course.courseCode = ccode
		course.courseName = cname
		db.Exec("Insert into courses(course_code,name)values(?,?)", course.courseCode, course.courseName)
		_, err := db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", index, course.courseCode, time.Now())
		if err != nil {
			fmt.Println("Row not inserted", err)
		}
	}
	o := dbmap.GetCourses(name2)
	if !reflect.DeepEqual(o, courses) {
		log.Println("Results does not match in Get Courses TC3")
		t.Fail()
	}
}

//Get student courses where student is not enrolled to any course
func TestGetCourses4(t *testing.T) {
	index += 1
	_, name2, coursecode, coursename, mobilenum := initialvalues1()

	db, _ := sql.Open("mysql", "root:Bond007@/db002?parseTime=true")
	dbmap := GetConnections()

	var courses []Course
	sName := name2 + strconv.Itoa(index)
	student := &Student{id: index, name: sName, mobileNum: mobilenum}
	db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)
	for k := 0; k < 3; k++ {
		ccode := coursecode + strconv.Itoa(k)
		cname := coursename + strconv.Itoa(k)
		var course Course
		course.courseCode = ccode
		course.courseName = cname
		//courses = append(courses,course)
		db.Exec("Insert into courses(course_code,name)values(?,?)", course.courseCode, course.courseName)
		//db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(id,course.courseCode,time.Now()))
	}
	o := dbmap.GetCourses(sName)
	if !reflect.DeepEqual(o, courses) {
		log.Println("Results does not match in Get Courses TC4")
		t.Fail()
	}
}
