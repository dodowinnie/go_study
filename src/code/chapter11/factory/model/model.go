package model

type student struct{
	Name string
	score float64
}

func (stu *student) GetScore() float64{
	return stu.score
}

func NewStudent(name string, score float64) *student{
	return &student{
		Name:name,
		score:score,
	}
}