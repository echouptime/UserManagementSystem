package controllers

import (
	"UserManagementSystem/models"
	"UserManagementSystem/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var Users []*models.User

//生成ID
func GetId() int {
	id := 0
	for _, user := range Users {
		if user.Id > id {
			id = user.Id
		}
	}
	return id + 1
}

func BaseInformation(w http.ResponseWriter, r *http.Request) {

	//默认初始化数据
	Users = []*models.User{
		{1, "杨旭", "运维部", "北京市", true, 3000, "13716977836"},
		{2, "张福权", "项目管理", "北京市", false, 6000, "10086"},
		{3, "张宝义", "研发部", "北京市", true, 5000, "12306"},
		{4, "陈国荣", "产品部", "北京市", true, 9000, "10086"},
		{5, "贾强军", "Siteops", "北京市", true, 8000, "10085"},
	}

	//读取持久化Json内容
	if utils.FileIsExists("userInfo.json") == true {
		jsontxt, _ := ioutil.ReadFile("userInfo.json")
		err := json.Unmarshal(jsontxt, &Users)
		if err != nil {
			log.Fatal(err)
		}
		tpl := template.Must(template.ParseFiles("templates/index.html"))
		if err := tpl.ExecuteTemplate(w, "index.html", Users); err != nil {
			fmt.Println(err)
		}
	} else {
		tpl := template.Must(template.ParseFiles("templates/index.html"))
		if err := tpl.ExecuteTemplate(w, "index.html", Users); err != nil {
			fmt.Println(err)
		}

	}
}

//创建用户
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("templates/create.html"))
		if err := tpl.ExecuteTemplate(w, "create.html", Users); err != nil {
			log.Fatalln(err)
		}
	} else if r.Method == "POST" {
		//添加用户信息
		sal, _ := strconv.Atoi(r.FormValue("salary"))

		//将数据存储到Json
		Users = append(Users, &models.User{
			Id:         GetId(),
			Name:       r.FormValue("name"),
			Department: r.FormValue("department"),
			Addr:       r.FormValue("addr"),
			Sex:        r.FormValue("sex") == "0",
			Salary:     sal,
			Phone:      r.FormValue("phone"),
		})

		//将数据存储到Json
		utils.SaveDb(Users)

	}
	http.Redirect(w, r, "/", 302)
}

//查询用户
func QueryUser(w http.ResponseWriter, r *http.Request) {
	TempUser := []*models.User{}

	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("templates/query.html"))
		tpl.ExecuteTemplate(w, "query.html", nil)
	} else {
		QueryInfo := r.FormValue("info")

		for _, users := range Users {
			if strings.Contains(QueryInfo, users.Name) ||
				strings.Contains(QueryInfo, users.Addr) ||
				strings.Contains(QueryInfo, users.Department) ||
				strings.Contains(QueryInfo, strconv.Itoa(users.Salary)) ||
				strings.Contains(QueryInfo, users.Phone) {
				TempUser = append(TempUser, users)
			}
		}
		if len(TempUser) == 0 || QueryInfo == "" {
			Messages := "User Is Null"
			tpl := template.Must(template.ParseFiles("templates/error.html"))
			tpl.ExecuteTemplate(w, "error.html", Messages)
		}

		tpl := template.Must(template.ParseFiles("templates/query.html"))
		tpl.ExecuteTemplate(w, "query.html", TempUser)
		//TempUser = TempUser[:0]

	}

}
