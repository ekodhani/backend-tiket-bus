package controllers

import (
	"backend-pertama/database"
	"backend-pertama/models"
	"backend-pertama/query"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

	// Hash Password
	hashedPassword := md5.Sum([]byte(user.Password))
	encodePassword := hex.EncodeToString(hashedPassword[:])

	// Konek Database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// Cek Query Usernamenya ada ga ?
	var usernameScan string
	var passwordScan string
	cek_username := query.CekUsername(user.Username)
	errUsername := db.QueryRow(cek_username).Scan(&usernameScan)

	if errUsername != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Username tidak ditemukan",
		})
	}

	// Cek Query Passwordnya salah ga ?
	if len(usernameScan) > 0 {
		cek_password := query.CekPassword(encodePassword)
		errPassword := db.QueryRow(cek_password).Scan(&passwordScan)

		if errPassword != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"message": "Password Salah",
			})
		}
	}

	// Jika Semuanya ada maka baru bisa masuk ke dashboard
	if len(usernameScan) > 0 && len(passwordScan) > 0 {
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
				c.IndentedJSON(http.StatusNotFound, gin.H{
					"message": "data tidak di temukan",
				})
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
		// // Query ke database untuk mendapatkan data kursi berdasarkan id bus
		// getKursi, errorKursi := db.Query(`SELECT * FROM kursi WHERE id_bus = ` + strconv.Itoa(dataBus.Id) + ``)
		// if errorKursi != nil {
		// 	fmt.Println("Error querying kursi data:", errorKursi.Error())
		// 	return
		// }
		// defer getKursi.Close()
	}

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

func SaveTiket(c *gin.Context) {
	// Deklarasi struct
	var tiket models.Tiket

	// Cek Json Error ?
	if err := c.ShouldBindJSON(&tiket); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Koneksi Database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var query, query_insert_detail string
	tiket.Create_date = time.Now().Unix()

	// Insert Ke Pemesanan Tiket
	query = "INSERT INTO pemesanan_tiket (id_bus, kota_asal, kota_tujuan, id_pembayaran, pergi, pulang, create_date, create_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(query,
		tiket.Id_bus,
		tiket.Kota_asal,
		tiket.Kota_tujuan,
		tiket.Pembayaran,
		tiket.Pergi,
		tiket.Pulang,
		tiket.Create_date,
		tiket.Create_by,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		// GET ID tiket yang sudah di insert
		insertedID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		} else {
			status := 0
			// Insert Ke Detail Pemesanan Tiket
			for i := range tiket.Data_penumpang {
				query_insert_detail = "INSERT INTO detail_pemesanan_tiket (id_tiket, nama, nik, id_kursi, status) VALUES (?, ?, ?, ?, ?)"
				_, err = db.Exec(query_insert_detail,
					insertedID,
					tiket.Data_penumpang[i].Nama,
					tiket.Data_penumpang[i].Nik,
					tiket.Data_penumpang[i].Id_kursi,
					status,
				)

				if err != nil {
					c.JSON(http.StatusBadRequest, err.Error())
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
	}
}
