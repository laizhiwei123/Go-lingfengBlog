package controllers

import  "net/http"

type Cookice struct {

}

func (*Cookice) Read(r *http.Request, cookiceName string)  (string, error){
	cookice, err := r.Cookie(cookiceName)
	if cookice  != nil{
		return cookice.Value, nil
	}
	return "", err
}