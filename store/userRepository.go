package store

import (
	"fmt"
	"github.com/google/uuid"
	"restApi/internal/model"
	"strconv"
	"time"
)

type UserRepository struct {
	store *Store
}

var database_url = "host=localhost port=5432 user=usergo password=0000 dbname=restdb sslmode=disable"

type tmp struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
	Email     string
	Age       uint
	Created   time.Time
}

func (r *UserRepository) Create(u *model.User) error {

	u.ID = uuid.New()
	month, _ := strconv.Atoi(time.Now().Month().String())
	u.Created = Date(time.Now().Year(), month, time.Now().Day())
	if _, err := r.store.db.Query(
		"INSERT INTO users (uuid,firstname,lastname,email,age, created) VALUES ($1,$2,$3,$4,$5,$6)",
		u.ID.String(),
		u.Firstname,
		u.Lastname,
		u.Email,
		u.Age,
		u.Created,
	); err != nil {

		return err
	}
	return nil
}

func (r *UserRepository) EditUser(u *model.User) (*model.User, error) {
	//u := &model.User{}
	fmt.Println(u.ID, u.Firstname, u.Email, u.Age)
	_, err := r.store.db.Exec(
		"UPDATE users SET firstname=$1, lastname=$2, email=$3, age=$4 WHERE uuid = $5",
		u.Firstname, u.Lastname, u.Email, u.Age,
		u.ID.String(),
	)

	if err != nil {
		return nil, err
	}

	return u, err
}

func (r *UserRepository) FindById(id string) (*model.User, error) {

	u := &model.User{}
	fmt.Println(id)

	if err := r.store.db.QueryRow("SELECT * FROM users WHERE uuid = $1", id).Scan(
		&u.ID,
		&u.Firstname,
		&u.Lastname,
		&u.Email,
		&u.Age,
		&u.Created,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetAll() (*[]model.User, error) {

	data, err := r.store.db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	umass := make([]model.User, 0, 100)
	for data.Next() {
		tmp := model.User{}
		err := data.Scan(&tmp.ID, &tmp.Firstname, &tmp.Lastname, &tmp.Email, &tmp.Age, &tmp.Created)
		if err != nil {
			return nil, err
		}
		umass = append(umass, tmp)
	}
	return &umass, nil
}

//Date
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day,
		time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC)
}
