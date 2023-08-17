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
	Id_bus      int      `json:"id_bus"`
	Kota_asal   string   `json:"kota_awal"`
	Kota_tujuan string   `json:"kota_tujuan"`
	Pembayaran  int      `json:"pembayan"`
	Pergi       int      `json:"pergi"`
	Pulang      int      `json:"pulang"`
	Id_kursi    []string `json:"kursi"`
}

type PilihKursi struct {
	Id_kursi string `json:"kursi"`
}
