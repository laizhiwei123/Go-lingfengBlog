package controllers

import (
	"Blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Sign_upStruct struct {

}
type userBodyDataStruct struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

func Sign_upController(w http.ResponseWriter, r *http.Request)  {
	sust := Sign_upStruct{}
	method := r.Method

	if method  == "GET"{
		sust.Get(w, r)
	}else if method == "POST"{
		sust.Post(w, r)
	}
}

func (this Sign_upStruct)Get (w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./views/sign_up.html")
	checkErr(err)

	t.Execute(w, nil)
}

func (this Sign_upStruct) Post(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	println(string(body))
	bodyJsonStruct := userBodyDataStruct{}
	err = json.Unmarshal(body, &bodyJsonStruct)
	checkErr(err)
	fmt.Println(bodyJsonStruct)
	find := models.AddUser(bodyJsonStruct.UserName, bodyJsonStruct.PassWord)
	data := make(map[string]string)
	if find {
		data["state"] = "Success"
		jsonData,err := json.Marshal(data)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(jsonData)	
	}else {
		data["state"] = "Fail"
		jsonData,err := json.Marshal(data)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(jsonData)
	}
}

