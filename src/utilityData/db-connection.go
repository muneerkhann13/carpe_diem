package utilityData

import (
	"common/configuration"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
	Author : Shivendra Pratap Singh
	CreatedDate : 16-Nov-2016
	Purpose : Starts Connection With DB
	Parameters : {
	"input":"",
	 "output":"error"
							}
	Module : Data
	LastUpdate : { "name": "Shivendra Pratap Singh", "date":"22 December 2016" }
*/

var DB []*sql.DB
var DBtravel, DBredus *sql.DB

//Starts DB Connection

func StartDB() error {
	fmt.Println("Reached start db")
	for _, value := range configuration.Config.MysqlDatabase {
		dbconnection := value.Username + ":" + value.Password + "@" + value.Host + "/" + value.Name + "?charset=utf8"
		db, err := sql.Open("mysql", dbconnection)
		if err != nil {
			return err
		}
		DB = append(DB, db)
	}
	DBtravel = DB[0]
	DBredus = DB[1]
	return nil
}
