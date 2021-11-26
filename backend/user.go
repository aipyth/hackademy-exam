package main

import (
    "crypto/md5"
    "net/http"
    "encoding/json"
    "errors"
)

type User struct {
    Email string
    PasswordDigest string
}

type UserRepository interface {
	Add(string, User) error
	Get(string) (User, error)
	Update(string, User) error
	Delete(string) (User, error)
}

type UserService struct {
	Repository UserRepository
}

type SignUpParams struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

type SignInParams struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

func (us *UserService) SignUp(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    params := &SignUpParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

    passwordDigest := md5.New().Sum([]byte(params.Password))
    user := User{
        Email: params.Email,
        PasswordDigest: string(passwordDigest),
    }

	err = us.Repository.Add(user.Email, user)
	if err != nil {
		handleError(err, w)
		return
	}

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("registered"))
}

func (us *UserService) SignIn(w http.ResponseWriter, r *http.Request, jwtService *JWTService) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    params := &SignInParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

	passwordDigest := md5.New().Sum([]byte(params.Password))
	user, err := us.Repository.Get(params.Email)
	if err != nil {
		handleError(err, w)
		return
	}

	if string(passwordDigest) != user.PasswordDigest {
		handleError(errors.New("invalid login params"), w)
		return
	}

	token, err := jwtService.GenerateJwt(user)
	if err != nil {
		handleError(err, w)
	}

	w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(token))
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnprocessableEntity)
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(err.Error()))
}
