package query

func GetBelumBayar() string {
	return `SELECT * FROM pemesanan_tiket WHERE id_user = ?`
}
