package controllers

import (
	"Blog/controllers/BusinessPublicUtil"
	"Blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type QueryContenControllerStruct struct {
	
}

type queryResults struct {
	Id string	`json:"id"`
	Title string `json:"title"`
	UserName string `json:"userName"`
	Conten string `json:"conten"`
	CreatedTime  string`json:"CreatedTime"`

}

func QueryContenController(w http.ResponseWriter, r *http.Request)  {
	queryContenControllerStruct := QueryContenControllerStruct{}
	method := r.Method

	if method == "GET" {
		queryContenControllerStruct.Get(w, r)
	}else if method == "POST" {
		queryContenControllerStruct.Post(w, r)
	}
}

func (QueryContenControllerStruct) Get(w http.ResponseWriter, r *http.Request)  {
	accept := r.Header["Accept"][0]
	if accept == "*/*" {
		VerificationJwtUtil := BusinessPublicUtil.VerificationJwtUtil{}
		reqJsonData := VerificationJwtUtil.VerificationJwtRequestJsonData(r)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(reqJsonData)
	}else {
		url := r.URL.RawQuery
		subscript := strings.LastIndex(url, "=")
		if subscript != -1 {
			title := url[subscript + 1:]
			_, verificationTitleResult := models.QueryContentTitleAll(false, title)

			if verificationTitleResult {
				t, err := template.ParseFiles("./views/queryConnten.html")
				checkErr(err)
				t.Execute(w, nil)
			}else {
				t, err := template.ParseFiles("./views/err.html")
				checkErr(err)
				t.Execute(w, nil)
			}
		} else {
			t, err := template.ParseFiles("./views/err.html")
			checkErr(err)
			t.Execute(w, nil)
		}
	}
}

func (QueryContenControllerStruct) Post(w http.ResponseWriter, r *http.Request)  {
	body,err := ioutil.ReadAll(r.Body)
	bodydata := string(body)
	println(bodydata)
	checkErr(err)

	data, find := models.QueryContentTitleAll(true ,bodydata)
	//data := &data1
	if find {
		jsonDta, err := json.Marshal(data)
		checkErr(err)
		fmt.Println("json	",data)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(jsonDta)
	}
}