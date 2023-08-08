package controllers

import (
	"backend-pertama/database"
	"backend-pertama/models"
	"backend-pertama/query"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

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
