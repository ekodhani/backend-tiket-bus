package query

func GetBelumBayar(id_user string) string {
	return `SELECT * FROM pemesanan_tiket WHERE id_user = ` + id_user
}
