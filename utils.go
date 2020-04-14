package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Db : ハンドラをグローバル変数として定義
var Db *sql.DB

// PsqlExec : DBハンドラの作成
func psqlExec() {
	// .envからDB関連の環境変数を取得
	if err := godotenv.Load("./.env", ) ; err != nil {
		panic(err)
	}
	DBHOST := "localhost"
	DBPORT := os.Getenv("DB_PORT")
	DBUSER := os.Getenv("DB_USER")
	DBPASS := os.Getenv("DB_PASS")
	DBNAME   := os.Getenv("DB_NAME")

	// ハンドラ作成
	var err error
	pgCoString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHOST, DBPORT, DBUSER, DBPASS, DBNAME)
	Db, err = sql.Open("postgres", pgCoString)
	if err != nil {
		panic(err)
	}
	return
}

func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	// 引数は ResponseWriter, data, filenames(可変長引数)
	var files []string  // templates
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	// filesにいれた複数のtemplateを解析し，Mustに渡してtemplatesとして定義
	templates := template.Must(template.ParseFiles(files...))
	// templatesのなかから layout.html で定義されているtemplate "layout" を実行
	// "layout"だけ実行しておけば，後のtemplateはlayout中でincludeされるように実装しているのでおｋ
	// 実行時にdataを渡す
	templates.ExecuteTemplate(writer, "basis", data)
}