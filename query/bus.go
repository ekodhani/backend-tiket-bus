package query

func GetKursi() string {
	return `SELECT * FROM kursi`
}

func GetBus(tgl, asal, tujuan string) string {
	return `SELECT 
	A.id,
	A.nama_bus,
	A.nomor_polisi,
	A.jumlah_kursi,
	A.type,
	A.status,
	B.id,
	B.asal,
	B.tujuan,
	B.tgl_berangkat,
	C.id,
	C.no_kursi,
	C.status
	FROM bus A
	LEFT JOIN rute B on B.id_bus = A.id
	LEFT JOIN kursi C on C.id_bis = A.id
	
	WHERE B.tgl_berangkat <=` + tgl + ` and B.asal = '` + asal + `' and B.tujuan = '` + tujuan + `'`
}
