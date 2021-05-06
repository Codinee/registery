package main

import (
	"fmt"
	"registery"
)

func main() {
	dbmap := registery.GetConnections()

	s := dbmap.GetStudentsInCourse("")
	fmt.Println(s)
	c := dbmap.GetCourses("")
	fmt.Println(c)
	e := dbmap.EnrollStudent(" ", 89898989, c)
	fmt.Println(e)
	s1 := dbmap.GetCoursesForStudents(s)
	fmt.Println(s1)
}
