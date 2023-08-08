package query

func SignUp() string {
	return `INSERT INTO user(nama, no_telp, email, password, id_role, image) VALUES (?, ?, ?, ?, ?, ?)`
}

func SignIn() string {
	return `SELECT * FROM user where nama = ? AND password = ?`
}
