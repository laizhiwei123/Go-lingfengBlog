package controllers

import (
	_ "Blog/github.com/dgrijalva/jwt-go"
	"Blog/models"
	"Blog/util"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type bodyJsonStruct struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Remember bool `json:"remember"`
}

type Sign_inControllerStruct struct {
}

func Sign_inController(w http.ResponseWriter, r *http.Request) {
	lcst := Sign_inControllerStruct{}

	method := r.Method

	if method == "GET" {
		lcst.Get(w, r)
	} else if method == "POST" {
		lcst.Post(w, r)
	}
}

func (this *Sign_inControllerStruct) Get(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./views/sign_in.html")

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
	return
}

func (this *Sign_inControllerStruct) Post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	println(string(body))
	bodyJsonStruct := bodyJsonStruct{}
	err = json.Unmarshal(body, &bodyJsonStruct)
	checkErr(err)
	userName := bodyJsonStruct.UserName
	passWord := bodyJsonStruct.PassWord
	//remember := bodyJsonStruct.Remember
	id, find := models.VerificationUser(userName, passWord)
	reqData := make(map[string]string)
	if find{
		token, err := util.JwtUtil{}.New(id)
		if err != nil {
			println("Token:", err)
		}
		reqData["Token"] = token
		reqData["state"] = "Success"
		reqJsonData,err := json.Marshal(reqData)
		if err != nil {
			println("json", err)
		}
		fmt.Println("Tokens",string(reqJsonData))
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(reqJsonData)
	} else {
		reqData["state"] = "Fail"
		reqJsonData,err := json.Marshal(reqData)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(reqJsonData)
	}
	//handlerFunc := http.HandlerFunc(w, r)
	//data := make(map[string]string)
	//if find {
	//	println(bodyJsonStruct.Remember)
	//	if bodyJsonStruct.Remember{
	//		maxAge := 0
	//		userNameCookice := &http.Cookie{
	//			Name:   "userName",
	//			Value:  userName,
	//			Path:   "/",
	//			MaxAge: maxAge,
	//		}
	//		passWordCookice := &http.Cookie{
	//			Name:   "passWord",
	//			Value:  passWord,
	//			Path:   "/",
	//			MaxAge: maxAge,
	//		}
	//		http.SetCookie(w, userNameCookice)
	//		http.SetCookie(w, passWordCookice)
	//	}
	//	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//	data["state"] = "Success"
	//
	//	jsonData, err := json.Marshal(data)
	//	checkErr(err)
	//	println(string(jsonData))
	//	w.Write(jsonData)
	//} else {
	//	data["state"] = "Fail"
	//	jsonData, err := json.Marshal(data)
	//	checkErr(err)
	//	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//	w.Write(jsonData)
	//
	//}
	//return
}
