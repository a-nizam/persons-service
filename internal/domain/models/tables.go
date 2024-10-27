package models

import "time"

type Car struct {
	Id      int64  `db:"id"`
	MainId  int64  `db:"main_id"`
	Type    string `db:"type"`
	Model   string `db:"model"`
	Number  string `db:"number"`
	Driver  string `db:"driver"`
	Technic string `db:"technic"`
}

type Doc struct {
	Id     int64     `db:"id"`
	MainId int64     `db:"main_id"`
	Number string    `db:"number"` // Номер документа
	Date   time.Time `db:"date"`   // Дата выдачи документа
	Place  string    `db:"place"`  // Место выдачи документа
}

type Family struct {
	RId       int64     `db:"id"`
	MainId    int64     `db:"main_id"`
	Rodstvo   string    `db:"rodstvo"`
	Fio       string    `db:"fio"`
	Birthdate time.Time `db:"birth_date"`
	Workplace string    `db:"work_place"`
	Union     string    `db:"union"`
}

type Housing struct {
	Id         int64  `db:"id"`
	MainId     int64  `db:"main_id"`
	PlaceBirth string `db:"place_birth"`
	PlaceProp  string `db:"place_prop"`
	PlaceJitel string `db:"place_jitel"`
	Obespechen string `db:"obespechen"`
}

type Person struct {
	Id        int64     `db:"id"`
	Name      string    `db:"full_name"`
	Birthdate time.Time `db:"birth_date"`
	Doljnost  string    `db:"doljnost"`
	Zvanie    string    `db:"zvanie"`
	Phone     string    `db:"personal_number"`
	Cars      []Car     `db:"auto"`
	Docs      []Doc     `db:"docs"`
	Family    []Family  `db:"family"`
	Housing   Housing   `db:"housing"`
	Otpusk    []Otpusk  `db:"otpusk"`
	Phones    []Phone   `db:"phones"`
	Slujba    Slujba    `db:"slujba"`
	Study     []Study   `db:"study"`
}

type Otpusk struct {
	Id     int64  `db:"id"`
	MainId int64  `db:"main_id"`
	Type   string `db:"type"`
	DayOf  string `db:"day_of"`
	DaySt  string `db:"day_st"`
	Ride   string `db:"ride"`
}

type Phone struct {
	Id     int64  `db:"id"`
	MainId int64  `db:"main_id"`
	Phone  string `db:"phone_number"`
}

type Slujba struct {
	Id       int64  `db:"id"`
	MainId   int64  `db:"main_id"`
	DateSt   string `db:"date_st"`
	Kontrakt string `db:"kontrakt"`
	Cont     string `db:"cont"`
	Naznach  string `db:"naznach"`
	Zvan     string `db:"zvan"`
	Attestac string `db:"attestac"`
	Klassn   string `db:"klassn"`
}

type Study struct {
	Id          int64  `db:"id"`
	MainId      int64  `db:"main_id"`
	Obrazovanie string `db:"obrazovanie"`
	Type        string `db:"type"`
	UchebZaved  string `db:"ucheb_zaved"`
	Napravlenie string `db:"napravlenie"`
	Form        string `db:"form"`
	DateSt      string `db:"date_st"`
	DateEnd     string `db:"date_end"`
}
