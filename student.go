package registery

import (
	"log"
	"time"
)

type Student struct {
	id        int
	name      string
	mobileNum uint64
}

//Get Student details enrolled for the given Course
func (dbmap *dataBaseMap) GetStudentsInCourse(courseCode string) []Student {
	var query string
	var students []Student
	var args []interface{}

	if len(courseCode) == 0 || courseCode == " " {
		return students
	}

	query = "SELECT s.name, s.id, s.mobile_num FROM students s, enrollment e WHERE e.course_code = ? AND id = e.student_id"
	args = append(args, courseCode)

	for dbName := range dbmap.dbases {
		slist := dbmap.executeGetStudentsInCourseSql(dbName, query, args)
		students = append(students, slist...)
	}
	return students
}

//Database operation to get Students enrolled for the course
func (dbmap *dataBaseMap) executeGetStudentsInCourseSql(dbName string, query string, args []interface{}) []Student {
	db := dbmap.dbases[dbName].db
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("Error while fetching student list for courses ", err)
		return nil
	}
	defer rows.Close()
	var students []Student
	for rows.Next() {
		var std Student
		err = rows.Scan(&std.name, &std.id, &std.mobileNum)
		if err != nil {
			log.Println("Error while fetching student list for courses ", err)
			return nil
		}
		students = append(students, std)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error while fetching student list for courses ", err)
	}
	return students
}

//Enroll new students by inserting new student to Student and Enrollment tables
func (dbmap *dataBaseMap) EnrollStudent(name string, number uint64, courses []Course) error {
	var args []interface{}

	if len(name) == 0 || name == " " {
		log.Println("Name given is empty, nothing is inserted")
		return nil
	}
	id := getNewID()
	dbName := dbmap.getDataBase(name)

	query := "INSERT INTO students(id,name,mobile_num)VALUES(?,?,?)"
	args = append(args, id, name, number)
	err := dbmap.executeEnrollStudentSql(dbName, query, args)

	for _, c := range courses {
		args = nil
		query = "INSERT INTO enrollment(student_id,course_code,date_enrolled)VALUES(?,?,?)"
		args = append(args, id, c.courseCode, time.Now())
		err = dbmap.executeEnrollStudentSql(dbName, query, args)
	}
	return err
}

//Database operation to enroll students
func (dbmap *dataBaseMap) executeEnrollStudentSql(dbName string, query string, args []interface{}) error {
	db := dbmap.dbases[dbName].db
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Println("Error while enrolling student for courses ", err)
		return err
	}
	return nil
}
