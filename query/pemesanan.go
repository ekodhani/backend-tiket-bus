package query

func GetBelumBayar(id_user string) string {
	return `SELECT 
	A.id, 
	A.kota_asal, 
	A.kota_tujuan, 
	A.id_pembayaran, 
	A.pergi, 
	A.pulang, 
	A.status_pembayaran, 
	A.id_user, 
	A.create_date, 
	A.create_by,
	B.nama_bus, 
	B.nomor_polisi, 
	B.type, 
	B.harga, 
	B.jumlah_kursi 
	
	FROM pemesanan_tiket A 
	LEFT JOIN bus B on B.id = A.id_bus 
	
	WHERE A.id_user = ` + id_user
}
