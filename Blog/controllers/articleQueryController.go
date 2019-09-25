package controllers

import (
	"Blog/controllers/BusinessPublicUtil"
	"Blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type ArticieQueryControllerStruct struct {
}

type Title struct {
	Title string `json:"title"`
}

type Id struct {
	Id int32 `json:"Id"`
} 

func ArticieQueryController(w http.ResponseWriter, r *http.Request) {
	aqcs := ArticieQueryControllerStruct{}

	method := r.Method
	println(method)
	if method == "GET" {
		aqcs.Get(w, r)
	} else if method == "POST" {
		aqcs.Post(w, r)
	} else if method == "DELETE" {
		aqcs.Delete(w, r)
	}
}

func (ArticieQueryControllerStruct) Get(w http.ResponseWriter, r *http.Request) {
	//case "*VerificationUser*":
	//	data := make(map[string]string)
	//	cookice := Cookice{}
	//	userName, err :=  cookice.Read(r,"userName")
	//	if err != nil {
	//		data["state"] = "Fail"
	//		jsonData,err := json.Marshal(data)
	//		checkErr(err)
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		w.Write(jsonData)
	//		panic(err)
	//		return
	//	}
	//	passWord, err := cookice.Read(r, "passWord")
	//	if err != nil {
	//		data["state"] = "Fail"
	//		jsonData,err := json.Marshal(data)
	//		checkErr(err)
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		w.Write(jsonData)
	//		panic(err)
	//		return
	//	}
	//
	//	find := models.VerificationUser(userName, passWord)
	//
	//	if find {
	//		data["state"] = "Success"
	//		data["value"] = userName
	//		jsonData,err := json.Marshal(data)
	//		checkErr(err)
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		w.Write(jsonData)
	//	}else {
	//		data["state"] = "Fail"
	//		jsonData,err := json.Marshal(data)
	//		checkErr(err)
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		w.Write(jsonData)
	//	}
	//	return
	accept := r.Header["Accept"][0]
	fmt.Println(accept)
	switch accept {
	case "*/*":
		{
			requestData := models.ContentAll()
			jsonData, err := json.Marshal(requestData)
			checkErr(err)
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			w.Write(jsonData)
		}
	default:
		//Accept is text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
		verificationJwtUtil := BusinessPublicUtil.VerificationJwtUtil{}
		_, verificationUserResult  := verificationJwtUtil.VerificationJwtRequestUserName(r)
		fmt.Println()
		if verificationUserResult {
			fmt.Println("true")
			t, err := template.ParseFiles("./views/articleQuery.html")
			checkErr(err)
			t.Execute(w, nil)
		}else {
			t, err := template.ParseFiles("./views/errRedirectLogin.html")
			checkErr(err)
			t.Execute(w, nil)
		}

	}
}

func (ArticieQueryControllerStruct) Post(w http.ResponseWriter, r *http.Request) {
	//data:= make(map[string]string)

	//bodyData,err := ioutil.ReadAll(r.Body)
	//checkErr(err)
	//
	//err = json.Unmarshal(bodyData, &queryId)
	//checkErr(err)
	//if len(queryId.Id) > 0 {
	//	requestData,find := models.QueryContentId(queryId.Id)
	//	if find {
	//		jsonData, err := json.Marshal(requestData)
	//		checkErr(err)
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		w.Write(jsonData)
	//	}else {
	//		data["state"] = "Success"
	//		responseWriterJsonData, err := json.Marshal(data)
	//		checkErr(err)
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		w.Write(responseWriterJsonData)
	//	}
	//
	//} else {
	//	requestData := models.ContentAll()
	//	jsonData, err := json.Marshal(requestData)
	//	checkErr(err)
	//	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//	w.Write(jsonData)
	//}
	queryTitle := Title{}
	bodyData,err := ioutil.ReadAll(r.Body)
	checkErr(err)
	err = json.Unmarshal(bodyData, &queryTitle)
	checkErr(err)
	requestData, find := models.QueryContentTitle(queryTitle.Title)

	if find {
		requestJsonData, err := json.Marshal(requestData)
		checkErr(err)
		println(string(requestJsonData))
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(requestJsonData)
	}else {
		requestNotFound := make(map[string]string)
		requestNotFound["state"] = "Fail"
		requestNotFoundJson, err := json.Marshal(requestNotFound)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(requestNotFoundJson)
	}
}

func (ArticieQueryControllerStruct) Delete(w http.ResponseWriter, r *http.Request) {
	deleteId := Id{}

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = json.Unmarshal(body, &deleteId)
	checkErr(err)

	find := models.DeleteArticle(deleteId.Id)
	data := make(map[string]string)
	if find {
		data["state"] = "Success"
		responseWriterJsonData, err := json.Marshal(data)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(responseWriterJsonData)
		fmt.Println(responseWriterJsonData)
	} else {
		data["state"] = "Fail"
		responseWriterJsonData, err := json.Marshal(data)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(responseWriterJsonData)
	}
}
