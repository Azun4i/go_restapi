package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"restApi/internal/model"
	store2 "restApi/store"
	"strconv"
	"time"
)

//type APIserver struct {
//	config *Config
//	logger *logrus.Logger
//	router *mux.Router
//}
//
//func New(config *Config) *APIserver {
//	return &APIserver{
//		config: config,
//		logger: logrus.New(),
//		router: mux.NewRouter(),
//	}
//}

func Start(config *Config) error {

	db, err := newdbConnect(config.DatabaseUrl)
	if err != nil {
		log.Fatal("can't open db connect", err)
	}
	defer db.Close()

	store := store2.New(db)
	s := newserver(*store)
	return http.ListenAndServe(config.BindAddr, s)
}

// newdbConnect  set database connection
func newdbConnect(DatabaseUrl string) (db *sql.DB, err error) {

	fmt.Println(DatabaseUrl)
	db, err = sql.Open("postgres", DatabaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	u := &model.User{}
	u.ID = uuid.New()
	u.Created = time.Now()
	month, _ := strconv.Atoi(time.Now().Month().String())
	u.Created = store2.Date(time.Now().Year(), month, time.Now().Day())
	db.Query(
		"INSERT INTO users (uuid,firstname,lastname,email,age, created) VALUES ($1,$2,$3,$4,$5,$6)",
		u.ID,
		u.Firstname,
		u.Lastname,
		u.Email,
		u.Age,
		u.Created,
	)

	return db, err
}
