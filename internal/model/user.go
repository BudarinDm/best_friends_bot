package model

type UserBirthday struct {
	Month     int64
	MonthText string
	Day       int64
	FIO       string
}

var BDText = map[int64]string{
	1:  "Января",
	2:  "Февраля",
	3:  "Марта",
	4:  "Апреля",
	5:  "Мая",
	6:  "Июня",
	7:  "Июля",
	8:  "Августа",
	9:  "Сентября",
	10: "Октября",
	11: "Ноября",
	12: "Декабря",
}
