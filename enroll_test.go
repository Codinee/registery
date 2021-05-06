package registery

import (
	"database/sql"
	"log"
	"reflect"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func initialvalues4() (string, string, string, string, uint64) {
	name1 := "Hal"
	name2 := "Ran"
	coursecode := "AGG"
	coursename := "Testing courses"
	var mobilenum uint64
	mobilenum = 1212121212
	return name1, name2, coursecode, coursename, mobilenum
}

//Enroll student in db1
func TestEnrollement1(t *testing.T) {
	index = 300
	name1, _, coursecode, coursename, mobilenum := initialvalues4()
	var courses []Course
	db, _ := sql.Open("mysql", "root:Bond007@/db001?parseTime=true")
	for k := 0; k < 3; k++ {
		ccode := coursecode + strconv.Itoa(k)
		cname := coursename + strconv.Itoa(k)
		var course Course
		course.courseCode = ccode
		course.courseName = cname
		courses = append(courses, course)
		db.Exec("Insert into courses(course_code,name)values(?,?)", course.courseCode, course.courseName)
	}

	dbmap := GetConnections()
	sname := name1 + strconv.Itoa(index)
	dbmap.EnrollStudent(sname, mobilenum, courses)

	o := dbmap.GetCourses(sname)
	if !reflect.DeepEqual(o, courses) {
		log.Println("Results does not match in Enroll students TC1")
		t.Fail()
	}
}

//Enroll student in db2
func TestEnrollement2(t *testing.T) {
	index += 1
	_, name2, coursecode, coursename, mobilenum := initialvalues4()
	var courses []Course
	db, _ := sql.Open("mysql", "root:Bond007@/db002?parseTime=true")
	for k := 0; k < 3; k++ {
		ccode := coursecode + strconv.Itoa(k)
		cname := coursename + strconv.Itoa(k)
		var course Course
		course.courseCode = ccode
		course.courseName = cname
		courses = append(courses, course)
		db.Exec("Insert into courses(course_code,name)values(?,?)", course.courseCode, course.courseName)
	}

	dbmap := GetConnections()
	sname := name2 + strconv.Itoa(index)
	dbmap.EnrollStudent(sname, mobilenum, courses)

	o := dbmap.GetCourses(sname)
	if !reflect.DeepEqual(o, courses) {
		log.Println("Results does not match in Enroll students TC2")
		t.Fail()
	}
}
