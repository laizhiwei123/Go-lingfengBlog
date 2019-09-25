package BusinessPublicUtil

import (
	"Blog/models"
	"Blog/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type VerificationJwtUtil struct {
	//Take HTTP Header Token Verification Business
}

func (VerificationJwtUtil) VerificationJwtRequestJsonData(r *http.Request) ([]byte) {
	fmt.Println(r.Header["Token"])
	reqData := make(map[string]string)
	token, err := r.Cookie("Token")
	if err != nil {
		fmt.Println("Cookice Token", err)
		reqData["state"] = "Fink"
		reqJsonData, err := json.Marshal(reqData)

		if err != nil {
			fmt.Println("Json:", err)
		}
		return reqJsonData
	}
	jwtUtil := util.JwtUtil{}
	fmt.Println(token.Value)
	jwtClaims, verificationjwtResult := jwtUtil.Verification(token.Value)
	if verificationjwtResult {
		userName, verificationIdResult := models.VerificationUserId(jwtClaims.Uid)
		if verificationIdResult {
			reqData["state"] = "Success"
			reqData["userName"] = userName
			reqJsonData, err := json.Marshal(reqData)
			if err != nil {
				fmt.Println("Json:", err)
			}
			return reqJsonData
		} else {
			reqData["state"] = "Fink"
			reqJsonData, err := json.Marshal(reqData)

			if err != nil {
				fmt.Println("Json:", err)
			}
			return reqJsonData
		}
	} else {
		reqData["state"] = "Fink"
		reqJsonData, err := json.Marshal(reqData)

		if err != nil {
			fmt.Println("Json:", err)
		}
		return reqJsonData
	}
}

func (VerificationJwtUtil) VerificationJwtRequestUserName(r *http.Request) (string, bool) {

	token, err := r.Cookie("Token")
	if err != nil {
		fmt.Println("Cookies:", err)
		return "", false
	}

	jwtUtil := util.JwtUtil{}
	JwtClaims, verificationJwtResult := jwtUtil.Verification(token.Value)

	if verificationJwtResult {
		userName, verificationIdResult := models.VerificationUserId(JwtClaims.Uid)

		if verificationIdResult {
			return userName, true
		} else {
			return "", false
		}
	} else {
		return "", false
	}
}
