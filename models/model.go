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

type DataBus struct {
	Id           int         `json:"id"`
	Nama_bus     string      `json:"nama_bus"`
	Nomor_polisi string      `json:"nomor_polisi"`
	Jumlah_kursi int         `json:"jumlah_kursi"`
	Type         string      `json:"type"`
	Harga        int         `json:"harga"`
	Status       int         `json:"status"`
	Detail_bus   []DetailBus `json:"detail_bus"`
}

type DetailBus struct {
	Id       int    `json:"id"`
	Id_bus   int    `json:"id_bus"`
	No_kursi string `json:"no_kursi"`
	Status   string `json:"status"`
}

type Otw struct {
	Asal          string `json:"asal"`
	Tgl_berangkat int    `json:"tgl_berangkat"`
	Tgl_pulang    int    `json:"tgl_pulang"`
	Tujuan        string `json:"tujuan"`
}

type Tiket struct {
	Id_bus         int             `json:"id_bus"`
	Kota_asal      string          `json:"kota_awal"`
	Kota_tujuan    string          `json:"kota_tujuan"`
	Pembayaran     int             `json:"pembayaran"`
	Pergi          int64           `json:"pergi"`
	Pulang         int64           `json:"pulang"`
	Id_user        string          `json:"id_user"`
	Create_by      string          `json:"create_by"`
	Create_date    int64           `json:"create_date"`
	Data_penumpang []DataPenumpang `json:"data_penumpang"`
}

type PilihKursi struct {
	Id_kursi string `json:"kursi"`
}

type DataPenumpang struct {
	Nama     string `json:"name"`
	Nik      string `json:"nik"`
	Id_kursi string `json:"kursi"`
}

type DataOrder struct {
	Id             int             `json:"id"`
	Id_bus         int             `json:"id_bus"`
	Kota_asal      string          `json:"kota_awal"`
	Kota_tujuan    string          `json:"kota_tujuan"`
	Id_pembayaran  int             `json:"id_pembayaran"`
	Pergi          int64           `json:"pergi"`
	Pulang         int64           `json:"pulang"`
	Id_user        int             `json:"id_user"`
	Create_date    int64           `json:"create_date"`
	Create_by      string          `json:"create_by"`
	Data_penumpang []DataPenumpang `json:"data_penumpang"`
}
