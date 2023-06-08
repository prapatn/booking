package entities

type User struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	NamePrefix string    `json:"name_prefix"`
	FisrtName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Username   string    `json:"username"`
	SumGrade   string    `json:"sum_grade"`
	AGrade     string    `json:"a_grade"`
	BGrade     string    `json:"b_grade"`
	CGrade     string    `json:"c_grade"`
	Bookings   []Booking `gorm:"foreignKey:UsersID" json:"bookings"`
}

func (u *User) getFullName() string {
	if u.NamePrefix != "" {
		return u.NamePrefix + "." + u.FisrtName + " " + u.LastName
	}
	return u.FisrtName + " " + u.LastName
}
