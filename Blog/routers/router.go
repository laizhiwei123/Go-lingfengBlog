package routers

import (
	"Blog/controllers"
	"log"
	"net/http"
)

func Router()  {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", controllers.IndexContrller)
	http.HandleFunc("/sign_in", controllers.Sign_inController)
	http.HandleFunc("/sign_up", controllers.Sign_upController)
	http.HandleFunc("/conten/",controllers.ContenController)
	http.HandleFunc("/queryContentTitle", controllers.QueryContenController)
	http.HandleFunc("/articleAdd", controllers.ArticleAddController)
	http.HandleFunc("/articleQuery", controllers.ArticieQueryController)
	http.HandleFunc("/articleQuery/modify",controllers.ArticleQueryModifyController)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
