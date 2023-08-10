package controllers

import (
	"backend-pertama/database"
	"backend-pertama/models"
	"backend-pertama/query"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Generate hash dari password
	hashedPassword := md5.Sum([]byte(user.Password))
	encodePassword := hex.EncodeToString(hashedPassword[:])

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.SignUp()
	rows, err := db.Query(query, &user.Name, &user.NoTelp, &user.Email, encodePassword, &user.Id_role, &user.Image)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	c.IndentedJSON(http.StatusOK, user)
}

func Signin(c *gin.Context) {
	var user models.UserLog

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword := md5.Sum([]byte(user.Password))
	encodePassword := hex.EncodeToString(hashedPassword[:])

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.SignIn()
	rows, err := db.Query(query, &user.Username, encodePassword)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	var results []models.Log

	// Loop melalui setiap baris hasil dan tambahkan data ke dalam slice
	for rows.Next() {

		var user models.Log

		// Scan data dari baris ke variabel-variabel
		err := rows.Scan(&user.IdUser, &user.Nama, &user.NoTelp, &user.Email, &user.Password, &user.IdRole, &user.Image)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}

		// Tambahkan peta ke dalam slice
		results = append(results, user)
	}

	// Cek apakah ada kesalahan saat mengiterasi melalui hasil
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}

func GetKursi(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.GetKursi()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	var results []models.DataKursi

	// Loop melalui setiap baris hasil dan tambahkan data ke dalam slice
	for rows.Next() {

		var dataKursi models.DataKursi

		// Scan data dari baris ke variabel-variabel
		err := rows.Scan(&dataKursi.IdKursi, &dataKursi.IdBis, &dataKursi.NoKursi, &dataKursi.Status)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}

		// Tambahkan peta ke dalam slice
		results = append(results, dataKursi)
	}
	fmt.Println(results)

	// Cek apakah ada kesalahan saat mengiterasi melalui hasil
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}

func GetBus(c *gin.Context) {
	// Inisialisasi Struct Data Json
	var otw models.Otw
	// Hook data Json yg dikirim dan cek jika error tampilkan kenapa
	if err := c.ShouldBindJSON(&otw); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("tgl berangkat : ", otw.Tgl_berangkat)
	// Koneksi Database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// Query
	query := query.GetBus(strconv.Itoa(otw.Tgl_berangkat), otw.Asal, otw.Tujuan)
	rows, err := db.Query(query)
	// Jika Error quernya
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	// Buat Struct
	var results []models.DataBus

	// Loop melalui setiap baris hasil dan tambahkan data ke dalam slice
	for rows.Next() {

		var dataBus models.DataBus

		// Scan data dari baris ke variabel-variabel
		err := rows.Scan(&dataBus.Id, &dataBus.Nama_bus, &dataBus.Nomor_polisi, &dataBus.Jumlah_kursi, &dataBus.Type, &dataBus.Harga, &dataBus.Status)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}

		// Tambahkan peta ke dalam slice
		results = append(results, dataBus)
	}
	// Debug
	fmt.Println(results)

	// Cek apakah ada kesalahan saat mengiterasi melalui hasil
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err.Error())
		return
	}

	// Jika ada Tgl Pulangnya
	// if len(otw.Tgl_pulang) != 0 {
	// 	fmt.Println("tgl pulang : ", otw.Tgl_pulang)
	// }

	if len(results) > 0 { // jika true
		c.IndentedJSON(http.StatusOK, results)
	} else { // jika data tidak ada
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "data tidak di temukan",
		})
	}
}
