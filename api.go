package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

/*
	jsonでDBのUserTabelを操作するAPIを作成
	POSTメソッド 以外は IPAdress/api/id のように，idを元に処理を実行させる

	Ex.)
		curl -i -X GET http://127.0.0.1:8080/api/1
		curl -i -X POST -H "Content-Type: application/json" -d '{"Name": "John", "Major": "philosophy"}' http://127.0.0.1:8080/api/
*/

// api : APIを実装するためのハンドラ関数
func api(w http.ResponseWriter, r *http.Request) {
	var err error
	// リクエストメソッドに応じた関数を呼び出す分岐
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	// CRUDを行う関数の実行中にエラーが発生している場合は500Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
	CRUDを行う関数群を定義
*/

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// path.Baseはpathの最後の要素を取得する．これをidと想定して，Atoiで整数値に変換
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 指定されたidのレコードを取得
	u, err := retrieveI(id)
	if err != nil {
		return
	}
	// 構造体をjsonとして出力
	o, err := json.MarshalIndent(&u, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(o)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	// リクエストbody取得のため，バイト型スライスの入れ物を作成
	body := make([]byte, len)
	// 取得
	r.Body.Read(body)
	// リクエストのjsonに基づいてレコードをcreate
	var u User
	json.Unmarshal(body, &u)
	err = u.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 更新するレコードをid指定で取得
	u, err := retrieveI(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	// json用いて更新
	json.Unmarshal(body, &u)
	err = u.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 削除するレコードをid指定で取得
	u, err := retrieveI(id)
	if err != nil {
		return
	}
	err = u.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
