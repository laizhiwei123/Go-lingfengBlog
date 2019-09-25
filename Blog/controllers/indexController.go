package controllers

import (
	"Blog/controllers/BusinessPublicUtil"
	"Blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type IndexControllerStruct struct {
}

type TokenJson struct {
	token string `json:"token"`
}

func IndexContrller(w http.ResponseWriter, r *http.Request) {
	icst := IndexControllerStruct{}

	method := r.Method
	if method == "GET" {
		icst.Get(w, r)
	}
}

func (this *IndexControllerStruct) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header["Accept"], "Accept")
	if r.Header["Accept"][0] == "*VerificationUser*" {
		//Verification User Login
		fmt.Println("Token")
		fmt.Println(r.Cookie("Token"))
		verificationJwtUtil := BusinessPublicUtil.VerificationJwtUtil{}
		reqJsonData := verificationJwtUtil.VerificationJwtRequestJsonData(r)
		fmt.Println(string(reqJsonData), "WWWW")
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(reqJsonData)
	} else if r.Header["Accept"][0] == "*/*" {

		data := models.ContentAll()
		jsonData, err := json.Marshal(data)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(jsonData)
		return
	} else {
		t, err := template.ParseFiles("./views/index.html")
		checkErr(err)

		t.Execute(w, nil)
	}
}
