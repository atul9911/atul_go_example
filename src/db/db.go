package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"encoding/json"
)


func main()  {
	db, err := sql.Open("mysql", "app:go123@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()  //check if db is available
	if err!= nil{
		log.Println("Data base connection is not available")
	}

	defer db.Close()  //close the connection outside the scope of a function
//  insert into db
//	stmt, err := db.Prepare("INSERT INTO table1(name,email_id) VALUES(?,?)")
//	if err != nil {
//		log.Fatal(err)
//	}
//	res, err := stmt.Exec("hellogo","hello@go.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//	lastId, err := res.LastInsertId()
//	if err != nil {
//		log.Fatal(err)
//	}
//	rowCnt, err := res.RowsAffected()
//	if err != nil {
//		log.Fatal(err)
//	}
//log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

// read from db

	type user  struct {
		name string
		email_id string
	}

	m := make(map[string]string)
	data:= make([]byte,)
	rows , err := db.Query("select name,email_id from table1")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err:= rows.Scan(&u.name,&u.email_id)
		if err!= nil{
			log.Fatal(err)
		}
		b,_ := json.Marshal(u)
		data=append(data,b[:])
	}
	err = rows.Err()
	if(err!= nil){
		log.Fatal(err)
	}

	for i :=range data{
		fmt.Print(data[i].name)
		fmt.Println("-",data[i].email_id)
	}
//
//	for k,v := range m{
//		fmt.Print(k)
//		fmt.Println("-",v)
//	}
////	fmt.Println(u)
	db.Close()
}

