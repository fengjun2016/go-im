package main

import (
	"fastIM/app/controller"
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

func registerView() {
	tpl, err := template.ParseGlob("./view/**/*")
	if err != nil {
		logrus.Fatal(err.Error())
	}

	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(w http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(w, tplName, nil)
		})
	}
}

func main() {

	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	http.HandleFunc("/contact/addfriend", controller.AddFriend)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/createcommunity", controller.CreateCommunity)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	http.HandleFunc("/chat", controller.Chat)
	http.HandleFunc("/attach/upload", controller.FileUpload)

	http.Handle("/asset/", http.FileServer(http.Dir("../")))
	http.Handle("/resource/", http.FileServer(http.Dir("../")))
	registerView()
	logrus.Fatal(http.ListenAndServe(":8081", nil))
}
