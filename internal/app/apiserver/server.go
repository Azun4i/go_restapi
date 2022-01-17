package apiserver

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"restApi/internal/model"
	"restApi/store"
	"strconv"
	"time"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newserver(store store.Store) *server {
	s := &server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/user", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/user/{id}", s.handleGetUserbyid()).Methods("GET")
	s.router.HandleFunc("/user/{id}", s.handeleEditUser()).Methods("PUT")
	s.router.HandleFunc("/users", s.handelGetallUser()).Methods("GET")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type tmpu struct {
		ID        uuid.UUID `json:"id"`
		Firstname string    `json:"firstname"`
		Lastname  string    `json:"lastname"`
		Email     string    `json:"email"`
		Age       string    `json:"age"`
		Created   time.Time `json:"created"`
	}

	return func(w http.ResponseWriter, req *http.Request) {
		tmp := &tmpu{}
		if err := json.NewDecoder(req.Body).Decode(tmp); err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		age, _ := strconv.Atoi(tmp.Age)
		u := &model.User{
			ID:        tmp.ID,
			Firstname: tmp.Firstname,
			Lastname:  tmp.Lastname,
			Email:     tmp.Email,
			Age:       uint(age),
			Created:   tmp.Created,
		}

		if err := s.store.User().Create(u); err != nil {
			s.respond(w, req, http.StatusUnprocessableEntity, nil)
		}
		s.respond(w, req, http.StatusCreated, u)
	}
}

func (s *server) handelGetallUser() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		u, err := s.store.User().GetAll()
		if err != nil {
			s.error(w, req, http.StatusNoContent, nil)
		}
		s.respond(w, req, http.StatusOK, u)
	}
}

func (s *server) handleGetUserbyid() http.HandlerFunc {

	type tmp_req struct {
		ID string
	}
	var ok bool

	tmp := &tmp_req{}
	return func(w http.ResponseWriter, req *http.Request) {
		p := mux.Vars(req)
		tmp.ID, ok = p["id"]
		if !ok {
			s.error(w, req, http.StatusNoContent, nil)
		}
		u, err := s.store.User().FindById(tmp.ID)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, nil)
		}
		s.respond(w, req, http.StatusOK, u)
	}
}

func (s *server) handeleEditUser() http.HandlerFunc {
	type tmp_req struct {
		ID string
	}

	type User struct {
		//ID        uuid.UUID `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Age       string `json:"age"`
	}

	var ok bool

	tmp_id := &tmp_req{}
	return func(w http.ResponseWriter, req *http.Request) {
		tmp_user := &User{}
		user := &model.User{}
		p := mux.Vars(req)
		tmp_id.ID, ok = p["id"]
		if !ok {
			s.error(w, req, http.StatusNoContent, nil)
		}

		if err := json.NewDecoder(req.Body).Decode(tmp_user); err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		ui := uuid.MustParse(tmp_id.ID)
		age, _ := strconv.Atoi(tmp_user.Age)
		user.ID = ui
		user.Firstname = tmp_user.Firstname
		user.Lastname = tmp_user.Lastname
		user.Email = tmp_user.Email
		user.Age = uint(age)
		u, err := s.store.User().EditUser(user)
		if err != nil {
			s.error(w, req, http.StatusNoContent, err)
		}
		s.respond(w, req, http.StatusOK, u)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {

	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
