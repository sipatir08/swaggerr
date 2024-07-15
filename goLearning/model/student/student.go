package student

type Student struct {
	StudentId int `json:"student_id"`
	CourseId int `json:"course_id"`
	Name string `json:"name"`
	Class string `json:"class"`
	Address string `json:"address"`
}