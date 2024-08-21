package helper

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func RandomNumber(n int) int {

	return rand.Intn(n)

}

type Person struct {
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    string `json:"has_dog"`
}

func CheckErr(err error) {
	if err != nil {

		fmt.Println("There is an error eccurd", err)

	}

}

func Devide(x, y float32) (float32, error) {

	var d float32

	if y == 0 {
		return d, errors.New("cannot devide by 0")
	}

	d = x / y

	return d, nil

}

func Devidering(w http.ResponseWriter, r *http.Request) {

	// res,err := Devide(100,11)
	// CheckErr(err)
	// fmt.Fprintf(w,fmt.Sprintf("this is the result number %f",res))

	xStr := r.URL.Query().Get("x")
	yStr := r.URL.Query().Get("y")

	x, err := strconv.ParseFloat(xStr, 64)
	if err != nil {
		http.Error(w, "Invalid value for x", http.StatusBadRequest)
		return
	}

	y, err := strconv.ParseFloat(yStr, 64)
	if err != nil {
		http.Error(w, "Invalid value for y", http.StatusBadRequest)
	}

	if y == 0 {
		http.Error(w, "Cannot devide by 0! it's not possible!", http.StatusBadRequest)
	}

	res, err := Devide(float32(x), float32(y))
	CheckErr(err)

	fmt.Fprintf(w, "this is the result: %f", res)

}
