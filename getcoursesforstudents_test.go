package registery

import (
	"database/sql"
	"log"
	"reflect"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func initialvalues3() (string, string, string, string, uint64) {
	name1 := "Dyn"
	name2 := "Ql w"
	coursecode := "Court"
	coursename := "Testing Sample courses"
	var mobilenum uint64
	mobilenum = 1212121278
	return name1, name2, coursecode, coursename, mobilenum
}

//Get courses for the students in the list
func TestGetCoursesForStudents1(t *testing.T) {
	index = 200
	name1, name2, coursecode, coursename, mobilenum := initialvalues3()

	db, _ := sql.Open("mysql", "root:Bond007@/db001?parseTime=true")
	dbmap := GetConnections()

	studentmap := make(map[Student][]Course)
	var courses []Course
	var students []Student

	sName := name1 + strconv.Itoa(index)
	var student Student
	student.id = index
	student.name = sName
	student.mobileNum = mobilenum
	students = append(students, student)
	db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)

	for k := 0; k < 3; k++ {
		ccode := coursecode + strconv.Itoa(k)
		cname := coursename + strconv.Itoa(k)
		var course Course
		course.courseCode = ccode
		course.courseName = cname
		courses = append(courses, course)
		db.Exec("Insert into courses(course_code,name)values(?,?)", course.courseCode, course.courseName)
		db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", student.id, course.courseCode, time.Now())
	}
	studentmap[student] = courses

	index += 1
	var ncourses []Course
	db, _ = sql.Open("mysql", "root:Bond007@/db002?parseTime=true")
	sName = name2 + strconv.Itoa(index)
	student.id = index
	student.name = sName
	student.mobileNum = mobilenum
	students = append(students, student)
	db.Exec("Insert into students(id,name,mobile_num) values(?,?,?)", student.id, student.name, student.mobileNum)
	for k := 0; k < 3; k++ {
		ccode := coursecode + strconv.Itoa(k)
		cname := coursename + strconv.Itoa(k)
		var course Course
		course.courseCode = ccode
		course.courseName = cname
		ncourses = append(ncourses, course)

		db.Exec("Insert into courses(course_code,name)values(?,?)", course.courseCode, course.courseName)
		db.Exec("Insert into enrollment(student_id,course_code,date_enrolled)values(?,?,?)", student.id, course.courseCode, time.Now())
	}
	studentmap[student] = ncourses
	o := dbmap.GetCoursesForStudents(students)
	if !reflect.DeepEqual(o, studentmap) {
		log.Println("Results does not match in Get courses for students list TC1")
		t.Fail()
	}

}

//Get courses for students who are not in system
func TestGetCoursesForStudents2(t *testing.T) {
	index += 1
	name1, name2, _, _, mobilenum := initialvalues3()

	//db := sql.Open("mysql", "root:Bond007@/db001?parseTime=true")
	dbmap := GetConnections()
	var students []Student
	studentmap := make(map[Student][]Course)
	sName := name1
	var student Student
	student.id = index
	student.name = sName
	student.mobileNum = mobilenum
	students = append(students, student)
	sName = name2
	student.id = index
	student.name = sName
	student.mobileNum = mobilenum
	students = append(students, student)
	o := dbmap.GetCoursesForStudents(students)
	if !reflect.DeepEqual(o, studentmap) {
		log.Println("Results does not match in Get courses for students list TC2")
		t.Fail()
	}
}
