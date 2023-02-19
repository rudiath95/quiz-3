package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Segitiga(c *gin.Context) {

	hitung := c.DefaultQuery("luas", "keliling")
	alas := c.Query("alas")
	tinggi := c.Query("tinggi")
	alasINT, _ := strconv.Atoi(alas)
	tinggiINT, _ := strconv.Atoi(tinggi)
	var hasil int
	if hitung == "luas" {
		hasil = alasINT * tinggiINT / 2
	} else {
		hasil = alasINT*2 + tinggiINT
	}

	c.String(http.StatusOK, "Hitung %s Segitiga = %d", hitung, hasil)
}

func Persegi(c *gin.Context) {

	hitung := c.DefaultQuery("luas", "keliling")
	sisi := c.Query("sisi")
	sisiINT, _ := strconv.Atoi(sisi)
	var hasil int
	if hitung == "luas" {
		hasil = sisiINT * sisiINT
	} else {
		hasil = sisiINT * 4
	}

	c.String(http.StatusOK, "Hitung %s Persegi = %d", hitung, hasil)
}

func PersegiPanjang(c *gin.Context) {

	hitung := c.DefaultQuery("luas", "keliling")
	p := c.Query("panjang")
	l := c.Query("lebar")

	pINT, _ := strconv.Atoi(p)
	lINT, _ := strconv.Atoi(l)
	var hasil int
	if hitung == "luas" {
		hasil = pINT * lINT
	} else {
		hasil = pINT*2 + lINT*2
	}

	c.String(http.StatusOK, "Hitung %s PersegiPanjang = %d", hitung, hasil)
}

func Lingkaran(c *gin.Context) {
	hitung := c.DefaultQuery("luas", "keliling")
	r := c.Query("jariJari")
	jariJariINT, _ := strconv.Atoi(r)
	var hasil int
	if hitung == "luas" {
		hasil = 22 * jariJariINT * jariJariINT / 7
	} else {
		hasil = 2 * 22 * jariJariINT / 7
	}

	c.String(http.StatusOK, "Hitung %s Lingkaran = %d", hitung, hasil)
}
