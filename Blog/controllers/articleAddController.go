package controllers

import (
	"Blog/controllers/BusinessPublicUtil"
	"Blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type ArticleAddStruct struct {

}

func ArticleAddController(w http.ResponseWriter, r *http.Request)  {
	aacr := ArticleAddStruct{}
	method := r.Method
	if method == "GET" {
		aacr.Get(w, r)
	}else if method  == "POST"{
		aacr.Post(w, r)
	}
}

func (this *ArticleAddStruct) Get(w http.ResponseWriter, r *http.Request) {
	if r. Header["Accept"][0] == "*/*"{
	//	usernamemap := make(map[string]string)
	//	cookice := Cookice{}
	//	userName,err := cookice.Read(r,"userName")
	//	if err != nil {
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		usernamemap["state"] = "Fail"
	//		data,err := json.Marshal(&usernamemap)
	//		checkErr(err)
	//		w.Write(data)
	//	}
	//	passWord,err := cookice.Read(r,"passWord")
	//	if err != nil {
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		usernamemap["state"] = "Fail"
	//		data,err := json.Marshal(&usernamemap)
	//		checkErr(err)
	//		w.Write(data)
	//	}
	//
	//	_, find := models.VerificationUser(userName, passWord)
	//
	//	if find {
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		usernamemap["state"] = "Success"
	//		usernamemap["value"] = userName
	//		data,err := json.Marshal(&usernamemap)
	//		checkErr(err)
	//		w.Write(data)
	//	}else {
	//		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//		usernamemap["state"] = "Fail"
	//		data,err := json.Marshal(&usernamemap)
	//		checkErr(err)
	//		w.Write(data)
	//	}
		verificationJwtUtil := BusinessPublicUtil.VerificationJwtUtil{}
		reqJsonData := verificationJwtUtil.VerificationJwtRequestJsonData(r)

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(reqJsonData)

	}else {
		t, err := template.ParseFiles("./views/articleAdd.html")
		checkErr(err)

		t.Execute(w, nil)
	}
}

type formData struct {
	//From Data Json models
	Title string `json:"input"`
	Content string	`json:"text"`
}
func (this *ArticleAddStruct) Post(w http.ResponseWriter, r *http.Request)  {
	data := make(map[string]string)
	body, err := ioutil.ReadAll(r.Body)//http Body Data
	checkErr(err)
	var formDataStruct formData
	err = json.Unmarshal(body, &formDataStruct)// Json  conversion struct
	checkErr(err)

	verificationJwtUtil := BusinessPublicUtil.VerificationJwtUtil{}
	userName,  verificationJwtResult := verificationJwtUtil.VerificationJwtRequestUserName(r)
	fmt.Println(formDataStruct.Content, formDataStruct.Title, userName )
	if verificationJwtResult {
		println("name:", userName, "Title:",formDataStruct.Title, "conent:", formDataStruct.Content)
		find := models.AddContent(userName, formDataStruct.Title, formDataStruct.Content)

		if find {
			data["state"] = "Success"
			jsonData,err := json.Marshal(data)
			fmt.Println(string(jsonData))
			checkErr(err)
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			w.Write(jsonData)
		}else {
			data["state"] = "Fail"
			jsonData,err := json.Marshal(data)
			checkErr(err)
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			w.Write(jsonData)
		}

	} else {
		data["state"] = "Fail"
		jsonData,err := json.Marshal(data)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(jsonData)
	}
}


func checkErr(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}
