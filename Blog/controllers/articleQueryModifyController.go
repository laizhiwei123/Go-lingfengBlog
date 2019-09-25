package controllers

import (
	"Blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type ArticleQueryModifyStruct struct {
	//http.Request
}

type modifyid struct {
	Id string `json:"Id"`
}
type modifyContent struct {
	Id string `json:"Id"`
	Content string `json:"Content"`
}

func ArticleQueryModifyController(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	articleQueryModifyStruct := ArticleQueryModifyStruct{}
	if method == "GET" {
		articleQueryModifyStruct.Get(w, r)
	} else if method == "POST" {
		articleQueryModifyStruct.Post(w, r)
	}
}

func (ArticleQueryModifyStruct) Get(w http.ResponseWriter, r *http.Request) {
	accept := r.Header["Accept"][0]
	urlId := r.URL.RawQuery[3:]
	switch accept {
	case "*/*":
		{
			dataMap, verificationIdResult := models.QueryContentId(urlId)
			if verificationIdResult {
				requestData := make(map[string]string)
				requestData["title"] = dataMap["Title"]
				requestData["state"] = "Success"
				requestJsonData, err := json.Marshal(requestData)
				checkErr(err)
				fmt.Println(string(requestJsonData))
				w.Header().Set("Content-Type", "application/json;charset=utf-8")
				w.Write(requestJsonData)
			}
		}
	default:
		{

			_, verificationIdResult := models.QueryContentId(urlId)

			if verificationIdResult {
				t, err := template.ParseFiles("./views/articleQueryModify.html")
				checkErr(err)
				t.Execute(w, nil)
			} else {
				t, err := template.ParseFiles("./views/errLogin.html")
				checkErr(err)
				t.Execute(w, nil)
			}
		}
	}
}

func (ArticleQueryModifyStruct) Post(w http.ResponseWriter, r *http.Request) {
	bodyJsonData, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	modifyContent := modifyContent{}
	err = json.Unmarshal(bodyJsonData, &modifyContent)
	checkErr(err)
	fmt.Println(modifyContent)
	responseData := make(map[string]string)
	find := models.ModifyArticle(modifyContent.Id,  modifyContent.Content)
	if find {
		responseData["state"] = "Success"
		responseJsonData, err := json.Marshal(responseData)
		checkErr(err)
		fmt.Println(string(responseJsonData))
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(responseJsonData)
	} else {
		responseData["state"] = "Fail"
		responseJsonData, err := json.Marshal(responseData)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(responseJsonData)
	}
}
