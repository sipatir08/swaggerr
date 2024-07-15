package course

type Course struct {
	CourseId int `json:"course_id"`
	CourseName string `json:"course_name"`
	Teacher string `json:"teacher"`
}