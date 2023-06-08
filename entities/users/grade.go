package users

import "github.com/go-playground/validator/v10"

type Grade struct {
	ID       string `gorm:"primarykey" json:"id" form:"id" validate:"required"`
	SumGrade string `json:"sum_grade" form:"sum_grade" `
	AGrade   string `json:"a_grade" form:"a_grade" validate:"required" `
	BGrade   string `json:"b_grade" form:"b_grade" validate:"required"`
	CGrade   string `json:"c_grade" form:"c_grade" validate:"required"`
}

func (g *Grade) Sum() {
	a := checkGrade(g.AGrade)
	b := checkGrade(g.BGrade)
	c := checkGrade(g.CGrade)

	g.SumGrade = calculate(a, b, c)
}

func (g *Grade) Validate() error {
	validate := validator.New()
	return validate.Struct(g)
}

func checkGrade(grade string) int {
	if grade == "A" {
		return 80
	} else if grade == "B" {
		return 70
	} else if grade == "C" {
		return 60
	} else {
		return 0
	}
}

func calculate(value ...int) string {
	sum := 0
	for _, v := range value {
		sum += v
	}
	sum = sum / len(value)
	if sum >= 80 {
		return "A"
	} else if sum >= 70 {
		return "B"
	} else if sum >= 60 {
		return "C"
	} else {
		return ""
	}
}
