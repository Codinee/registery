package registery

import (
	"log"
	"strings"
)

type Course struct {
	courseCode string
	courseName string
}

//Get Course details for the given Student
func (dbmap *dataBaseMap) GetCourses(studentName string) []Course {
	var courses []Course

	if len(studentName) == 0 || studentName == " " {
		return courses
	}
	var args []interface{}
	query := "SELECT c.name,c.course_code FROM courses  c, students s, enrollment e WHERE c.course_code = e.course_code AND e.student_id = s.id AND s. name = ?"
	args = append(args, studentName)

	dbName := dbmap.getDataBase(studentName)
	courses = dbmap.executeGetCoursesSql(dbName, query, args)
	return courses
}

//Database operation to get Course details
func (dbmap *dataBaseMap) executeGetCoursesSql(dbName string, query string, args []interface{}) []Course {

	db := dbmap.dbases[dbName].db
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("Error while fetching students courses ", err)
		return nil
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var course Course
		err = rows.Scan(&course.courseName, &course.courseCode)
		if err != nil {
			log.Println("Error while fetching students courses ", err)
			return nil
		}
		courses = append(courses, course)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error while fetching students courses ", err)
		return nil
	}
	return courses
}

//Get Course details for the list of Students
func (dbmap *dataBaseMap) GetCoursesForStudents(students []Student) map[Student][]Course {

	studentCourses := make(map[Student][]Course)

	if len(students) == 0 {
		return studentCourses
	}

	studentids := make(map[string][]int)

	for _, student := range students {
		for dbname, dbdetail := range dbmap.dbases {
			if student.name[0] >= dbdetail.startalpha && student.name[0] <= dbdetail.endalpha {
				studentids[dbname] = append(studentids[dbname], student.id)
			}
		}
	}

	var args []interface{}
	for dbname, ids := range studentids {
		args = nil
		for _, id := range ids {
			args = append(args, id)
		}
		questions := strings.Repeat(",?", len(ids)-1)
		query := "SELECT s.id, c.name,c.course_code FROM courses  c, students s, enrollment e WHERE c.course_code = e.course_code AND e.student_id = s.id AND s.id IN (?" + questions + ")"
		studentmap := dbmap.executeGetCoursesForStudentsSql(dbname, query, args)

		for _, student := range students {
			_, found := studentmap[student.id]
			if found {
				studentCourses[student] = studentmap[student.id]
			}
		}

	}
	return studentCourses
}

func (dbmap *dataBaseMap) executeGetCoursesForStudentsSql(dbName string, query string, args []interface{}) map[int][]Course {
	studentmap := make(map[int][]Course)

	db := dbmap.dbases[dbName].db
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("Error while fetching students courses ", err)
		return nil
	}
	defer rows.Close()
	//var courses []Course
	for rows.Next() {
		var course Course
		var id int
		err = rows.Scan(&id, &course.courseName, &course.courseCode)
		if err != nil {
			log.Println("Error while fetching students courses ", err)
			return nil
		}
		studentmap[id] = append(studentmap[id], course)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error while fetching students courses ", err)
		return nil
	}
	return studentmap
}
