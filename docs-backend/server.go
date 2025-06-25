package main

import (
	"fmt"
	"log"
    "os"
	"net/http"
    "net/url"
	"github.com/go-sql-driver/mysql"
	"database/sql"
    "github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
var residentHandlerFuncMap = map[string]func(*sql.DB, url.Values)(Resident,error){
    "GET": getResident,
    }


func getResident(db *sql.DB,params url.Values)([]Resident,error){
    var resident Resident
    if(params.Has("id")){
        return getResidentByID(db, params.Get("id"))
    }
    if(param.Has("name")){
        return getResidentByName(db, params.Get("name"))
    }
    return resident, nil
}

func residentHandler(w http.ResponseWriter, r * http.Request){
    action := r.Method
    params := r.URL.Query()
    db := initDB()
    fmt.Println(residentHandlerFuncMap[action](db, params))
    
    
}
func loadEnv(){
    err := godotenv.Load()
    
    if err != nil {
        log.Fatal("Error loading .env file")
  }
}
func initDB()(*sql.DB){
    
    var db *sql.DB
	// Capture connection properties.
    cfg := mysql.NewConfig()
    cfg.User = os.Getenv("MYSQL_USER")
    cfg.Passwd = os.Getenv("MYSQL_PASSWD")
    cfg.Net = "tcp"
    cfg.Addr = os.Getenv("MYSQL_ADDR")
    cfg.DBName = os.Getenv("MYSQL_DBNAME")
    cfg.ParseTime = true
    

    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
    return db
}


func main() {
    //http.HandleFunc("/", handler)
    loadEnv()
	db := initDB()
    http.HandleFunc("/residents/",residentHandler)

    fmt.Println(getAllResidents(db))
    log.Fatal(http.ListenAndServe(":8080", nil))
}