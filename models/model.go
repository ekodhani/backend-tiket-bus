package models

type User struct {
	Name     string `json:"name"`
	NoTelp   string `json:"no_telp"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Id_role  string `json:"id_role"`
	Image    string `json:"image"`
}

type UserLog struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type Log struct {
	IdUser   string `json:"id_log"`
	Nama     string `json:"nama"`
	NoTelp   string `json:"no_telp"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IdRole   string `json:"id_role"`
	Image    string `json:"image"`
}

type DataKursi struct {
	IdKursi string `json:"id"`
	IdBis   string `json:"id_bis"`
	NoKursi string `json:"no_kursi"`
	Status  string `json:"status"`
}
