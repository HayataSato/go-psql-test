package main

// User :
type User struct {
	ID    int
	Name  string
	Major string
}

// retrieveA : DB上のuserテーブルから全レコードを取り出してスライスで出力する関数
func retrieveA() (users []User, err error) {
	// DBにクエリを投げる
	rows, err := Db.Query("SELECT id, name, major FROM users ORDER BY id")
	if err != nil { // エラーなら中止
		return
	}
	// 成功ならloopで行を走査する
	for rows.Next() {
		u := User{} // u を，空の構造体Userとして作成
		// 行をスキャンして u に入れる
		if err = rows.Scan(&u.ID, &u.Name, &u.Major); err != nil {
			return // 空じゃなければ走査終了を示すのでloop終了
		}
		// 構造体Userを要素に持つスライスである users について，行をスキャンした u を追加して定義し直す
		users = append(users, u)
	}
	rows.Close()
	return
}

// retrieveI : idで任意のユーザーに関するレコード(行)を取得するための関数
func retrieveI(id int) (u User, err error) {
	u = User{}
	err = Db.QueryRow("select id, name, major from users where id = $1", id).Scan(&u.ID, &u.Name, &u.Major)
	return
}

/*
	Userのメソッド
*/

// Create
func (u *User) create() (err error) {
	// DBで自動で割り振られたidを取得
	statement := "insert into users (name, major) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// クエリを実行しつつ，Scanで構造体のIDにもDBのidと同じ数値を反映．
	// よって，入力 u のIDはdefault値(0)のままで問題ない．
	err = stmt.QueryRow(u.Name, u.Major).Scan(&u.ID)
	return
}

// Update "by ID" (IDで指定して更新するので，入力 u のIDは更新したいレコードのidを設定しておく必要がある)
func (u *User) update() (err error) {
	// ScanしないならExecの方が高速らしいのでExecメソッドを使用
	_, err = Db.Exec("update users set name = $2, major = $3 where id = $1", u.ID, u.Name, u.Major)
	return
}

// Delete "by ID" (同上)
func (u *User) delete() (err error) {
	_, err = Db.Exec("delete from users where id = $1", u.ID)
	return
}
