package models

import (
	_"Blog/github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"log"
	"time"
)


type indexContenData struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	UserName string `json:"UserName"`
	Content string `json:"Content"`
	CreatedTime string `json:"CreatedTime"`
}

var db  *sql.DB
var err  error
func Init()  {
	var err  error
	db, err= sql.Open("sqlite3", "./data/Blog.db")
	checkErr(err)

	newAdminSqlTable := `CREATE TABLE IF NOT EXISTS Admin (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	USERNAME VALCHAR(20) NOT NULL,
	PASSWORD VALCHAR(30) NOT NULL,
	CREATEDTIME DATE NOT NULL);
	`

	newUserSqlTable := `CREATE TABLE IF NOT EXISTS Users(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	USERNAME VALCHAR(20) NOT NULL,
	PASSWORD VALCHAR(30) NOT NULL,
	CREATEDTIME DATE NOT NULL);
	`

	newContent := `CREATE TABLE IF NOT EXISTS ArticleContent (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	USERNAME VALCHAR(20) NOT NULL,
	TITLE VALCHAR(30) NOT NULL,
	CONTENT VALCHAR(100) NOT NULL,
	CREATEDTIME DATE NOT NULL);
	`
	db.Exec(newAdminSqlTable)
	db.Exec(newUserSqlTable)
	db.Exec(newContent)
}

func AddAdmin(userName, password string) bool {
	return false
}

func VerificationAdmin()  {

}

func AddUser(username, passwrod string) bool {
	//AddUserData
	_, find := VerificationUser(username, passwrod)
	if find {
		return false
	}else {
		stmt,err := db.Prepare("INSERT INTO Users(USERNAME, PASSWORD, CREATEDTIME) values(?,?,?)")
		checkErr(err)

		_,err = stmt.Exec(username, passwrod, time.Now())
		checkErr(err)
		stmt.Close()
		return true
	}
}

func AddContent(userName interface{}, title, content string) bool {
	_, find := QueryContentTitle(title)
	if find {
		return false
	}
	stmt, err := db.Prepare("INSERT INTO ArticleContent(USERNAME, TITLE, CONTENT, CREATEDTIME	) values(?,?,?,?)")
	checkErr(err)
	println(userName, title, content)
	_,err = stmt.Exec(userName, title, content, time.Now())
	checkErr(err)
	stmt.Close()
	fmt.Println("Add")
	return true
}

func QueryContentTitle(title string) (map[string]string, bool) {
	fmt.Println(title, "QueryContentTitle")
	stmt, err := db.Prepare("select  ID, TITLE, USERNAME, CONTENT, CREATEDTIME from ArticleContent where TITLE = ?")
	checkErr(err)
	rows, err := stmt.Query(title)
	checkErr(err)
	for rows.Next() {
		println(" Get into Query")
		var  ids, titles, userNames,  contents,  createdTimes string
		err := rows.Scan(&ids,&titles, &userNames, &contents, &createdTimes)
		checkErr(err)
		if title == titles {
			println(ids, titles, userNames, "     ", contents, "      ", createdTimes)
			data := make(map[string]string)
			data["Id"] = ids
			data["Title"] = titles
			data["UserName"] = userNames
			data["Content"] = contents
			data["CreatedTime"] = createdTimes
			fmt.Println("QueryContentTitle")
			rows.Close()
			return data, true
		}
	}
	return nil, false
}
func QueryContentTitleAll(state bool, title string) ( []indexContenData, bool) {
	stmt, err := db.Prepare("select ID, TITLE, USERNAME, CONTENT, CREATEDTIME from ArticleContent where TITLE LIKE  ?")
	if err != nil {
		fmt.Println("sql", err)
	}
	rows, err := stmt.Query("%" +title +"%")

	if err != nil {
		fmt.Println("rows sql ", err)
	}
	data := make([]indexContenData, 0, 10)
	var ids, titles, userNames, contents, createdTimes string
	if state {
		for rows.Next() {
			err := rows.Scan(&ids, &titles, &userNames, &contents, &createdTimes)
			if err != nil {
				fmt.Println("err", err)
			}
			contenData := indexContenData{
				Id:          ids,
				Title:       titles,
				UserName:    userNames,
				Content:     contents,
				CreatedTime: createdTimes,
			}
			data = append(data, contenData)
			fmt.Println("dadad", data)
		}
		if data != nil {
			return data, true
		} else {
			return data, false
		}
	} else {
		if rows.Next() {
			err := rows.Scan(&ids, &titles, &userNames, &contents, &createdTimes)
			checkErr(err)
			contenData := indexContenData{
				Id:          ids,
				Title:       titles,
				UserName:    userNames,
				Content:     contents,
				CreatedTime: createdTimes,
			}
			data = append(data, contenData)
			fmt.Println("AAAA", data)
			rows.Close()
			return data, true
		}
		return nil, false
	}

}
func QueryContentId(id string) (map[string]string, bool) {
	stmt, err := db.Prepare("select ID, TITLE, USERNAME, CONTENT, CREATEDTIME from ArticleContent where ID = ? ")
	checkErr(err)
	rows,err := stmt.Query(id)
	if rows.Next() {
		var ids, titles, userNames, contents, createdTimes string
		err := rows.Scan(&ids, &titles, &userNames, &contents,&createdTimes)
		checkErr(err)

		if id == ids{
			println(ids, titles, userNames, "     ", contents, "      ", createdTimes)
			data := make(map[string]string)
			data["Id"] = ids
			data["Title"] = titles
			data["UserName"] = userNames
			data["Content"] = contents
			data["CreatedTime"] = createdTimes
			fmt.Println("QueryContentTitle")
			rows.Close()
			return data, true
		}
	}
	return nil, false
}

func ContentAll() ([]indexContenData) {
	rows, err := db.Query("select  ID, TITLE, USERNAME, CONTENT, CREATEDTIME from ArticleContent")
	checkErr(err)
	array := make([]indexContenData, 0, 2)
	for i := 0;rows.Next(); i++{
		println(" Get into Query")
		var  ids, titles, users,  contents,  createdTimes string
		err := rows.Scan(&ids,&titles, &users, &contents, &createdTimes)
		checkErr(err)
		println(ids, titles, users, "     ", contents, "      ", createdTimes)
		indexContenData := indexContenData{
			Id:ids,
			Title:titles,
			UserName:users,
			Content:contents,
			CreatedTime:createdTimes,
		}
		array = append(array, indexContenData)
	}
	return array
}

func VerificationUser(userName, passwrod string) (string, bool) {
	println(userName, passwrod)
	rows, err := db.Query("select  ID, USERNAME, PASSWORD from Users")
	checkErr(err)
	for rows.Next() {
		var id, pwd, user string
		//var date time.Time
		//var id  int
		err := rows.Scan(&id, &user, &pwd)
		checkErr(err)

		if (pwd == passwrod) && (user == userName) {
			println("Data", pwd, user)
			rows.Close()
			return id, true
		}
	}
	return "", false
}
func VerificationUserId(uid string) (string, bool) {
	rows, err := db.Query("SELECT  ID, USERNAME FROM Users")
	checkErr(err)

	for rows.Next() {
		var id, user string
		err := rows.Scan(&id, &user)
		checkErr(err)

		if id == uid {
			rows.Close()
			return user, true

		}
	}
	return "", false
}

func ModifyArticle(id,content string) bool {
	stmt,err := db.Prepare("UPDATE ArticleContent SET CONTENT = ?  WHERE ID = ?")
	checkErr(err)
	_, err = stmt.Exec(content, id)
	if err != nil {
		stmt.Close()
		return false
	}
	return true
}

func DeleteArticle(id int32) bool {
	stmt, err := db.Prepare("DELETE from ArticleContent where ID = ?")
	checkErr(err)
	_, err = stmt.Exec(id)

	if err != nil {
		stmt.Close()
		return false
	}
	return true
}

func checkErr(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}
