package entities

type User struct {
	ID         uint   `gorm:"primarykey"`
	NamePrefix string `json:"name_prefix"`
	FisrtName  string `json:"fisrt_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	SumGrade   string `json:"sum_grade"`
	AGrade     string `json:"a_grade"`
	BGrade     string `json:"b_grade"`
	CGrade     string `json:"c_grade"`
	Booking    []Booking
}
