package main

import (
	"go-im/app/appInit"
	"go-im/app/handler"
	"go-im/app/migrate"
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

	//读取配置文件
	appInit.InitConfig()

	//初始化数据库连接
	appInit.InitDB()

	//自动建表
	migrate.CreateTable()

	http.HandleFunc("/user/login", handler.UserLogin)
	http.HandleFunc("/user/register", handler.UserRegister)
	http.HandleFunc("/contact/addfriend", handler.AddFriend)
	http.HandleFunc("/contact/loadfriend", handler.LoadFriend)
	http.HandleFunc("/contact/loadcommunity", handler.LoadCommunity)
	http.HandleFunc("/contact/createcommunity", handler.CreateCommunity)
	http.HandleFunc("/contact/joincommunity", handler.JoinCommunity)
	http.HandleFunc("/chat", handler.Chat)
	http.HandleFunc("/message/loadmsg", handler.LoadPersonalMessage)
	http.HandleFunc("/attach/upload", handler.FileUpload)

	http.Handle("/asset/", http.FileServer(http.Dir("../")))
	http.Handle("/resource/", http.FileServer(http.Dir("../")))
	registerView()
	logrus.Fatal(http.ListenAndServe(":8081", nil))
}
