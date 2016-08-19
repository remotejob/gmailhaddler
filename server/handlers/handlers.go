package handlers

import (
	"log"

	"github.com/remotejob/gmail/domains"
	"gopkg.in/gcfg.v1"
)

var username string
var password string
var mechanism string

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		addrs = cfg.Dbmgo.Addrs
		database = cfg.Dbmgo.Database
		username = cfg.Dbmgo.Username
		password = cfg.Dbmgo.Password
		mechanism = cfg.Dbmgo.Mechanism

	}

}

// func AllRecords(w http.ResponseWriter, r *http.Request) {

// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:     addrs,
// 		Timeout:   60 * time.Second,
// 		Database:  database,
// 		Username:  username,
// 		Password:  password,
// 		Mechanism: mechanism,
// 	}

// 	//	dbsession, err := mgo.Dial("127.0.0.1")
// 	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer dbsession.Close()

// 	results := dbhandler.GetAll(*dbsession)

// 	output, err := json.Marshal(results)
// 	if err != nil {
// 		fmt.Println("Something went wrong!")
// 	} else {
// 		fmt.Fprintf(w, string(output))
// 	}

// }

// func OneRecords(w http.ResponseWriter, r *http.Request) {

// 	urlParams := mux.Vars(r)
// 	id := urlParams["id"]
// 	fmt.Println(id)

// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:     addrs,
// 		Timeout:   60 * time.Second,
// 		Database:  database,
// 		Username:  username,
// 		Password:  password,
// 		Mechanism: mechanism,
// 	}

// 	//	dbsession, err := mgo.Dial("127.0.0.1")
// 	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer dbsession.Close()

// 	if idint, err := strconv.Atoi(id); err != nil {

// 		http.Error(w, err.Error(), 500)

// 	} else {

// 		results := dbhandler.GetOne(*dbsession, idint)

// 		output, err := json.Marshal(results)
// 		if err != nil {
// 			fmt.Println("Something went wrong!")
// 		} else {
// 			fmt.Fprintf(w, string(output))
// 		}

// 	}

// }
