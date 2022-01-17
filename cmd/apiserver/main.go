package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"restApi/internal/app/apiserver"
	"time"
)

//сделать REST API на Go для создания/удаления/редактирования юзеров. Любой framework (или без него). Запушить код на github. В идеале с unit тестами. БД - PostgreSQL. Запросы:
//POST /users - create user
//GET /users/<id> - get user
//PUT /users/<id> - edit user

type User struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
	Email     string
	Age       uint
	Created   time.Time
}

var (
	configpath string
)

func init() {
	flag.StringVar(&configpath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	conf := apiserver.NewConfig()
	_, err := toml.DecodeFile(configpath, conf)
	if err != nil {
		log.Fatal("can't parse toml file ", err)
	}

	if err := apiserver.Start(conf); err != nil {
		log.Fatal(err)
	}
}
