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

type ContenControllerStruct struct {
}
type bodyDataStruct struct {
	Id string `json:"id"`
}

type contenStruct struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	UserName    string `json:"userName"`
	Conten      string `json:"content"`
	CreatedTime string `json:"CreatedTime"`
}

func ContenController(w http.ResponseWriter, r *http.Request) {
	contenControllerStruct := ContenControllerStruct{}

	method := r.Method

	if method == "GET" {
		contenControllerStruct.Get(w, r)
	} else if method == "POST" {
		contenControllerStruct.Post(w, r)
	}
}

func (ContenControllerStruct) Get(w http.ResponseWriter, r *http.Request) {
	accept := r.Header["Accept"][0]
	switch accept {
	case "*/*":
		{
			verificationJwtUtil := BusinessPublicUtil.VerificationJwtUtil{}
			responseJsonData := verificationJwtUtil.VerificationJwtRequestJsonData(r)
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			w.Write(responseJsonData)
		}
	default:
		{
			url := r.URL
			fmt.Println(strings.LastIndex(url.Path, "/"))
			subscript := strings.LastIndex(url.Path, "/")

			if subscript != -1 {
				id := url.Path[subscript + 1:]//subscript + 1 Remove itself
				fmt.Println(id)
				_, verificationIdResult := models.QueryContentId(id)
				fmt.Println(verificationIdResult)
				if verificationIdResult {
					t, err := template.ParseFiles("./views/conten.html")
					checkErr(err)
					t.Execute(w, nil)
				} else {
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
}

func (ContenControllerStruct) Post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	bodyDataStruct := bodyDataStruct{}
	err = json.Unmarshal(body, &bodyDataStruct)
	checkErr(err)
	data, find := models.QueryContentId(bodyDataStruct.Id)
	if find {
		contenStruct := contenStruct{
			Id:          data["Id"],
			Title:       data["Title"],
			UserName:    data["UserName"],
			Conten:      data["Content"],
			CreatedTime: data["CreatedTime"],}
		jsonData, err := json.Marshal(contenStruct)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(jsonData)
	}

}
