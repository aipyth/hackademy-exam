package main

import (
	"net/http"
    "log"
    "os"
	"os/signal"
    "context"
    "time"

	"github.com/gorilla/mux"
    // "github.com/rs/cors"
)

func wrapJwt(
	    jwt *JWTService,
	    f func(http.ResponseWriter, *http.Request, *JWTService),
) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		f(rw, r, jwt)
	}
}

func main() {
    router := mux.NewRouter()

    users := NewInMemoryStorage()
    userService := UserService{
        Repository: users,
    }

    todos := NewInMemoryToDoStorage()
    listsService := ListService{
        repository: todos,
    }

    jwtService, err := NewJWTService("public.rsa", "privkey.rsa")
	if err != nil {
		panic(err)
	}
   
    router.HandleFunc("/user/signup", logRequest(userService.SignUp)).
        Methods(http.MethodPost, http.MethodOptions)
    router.HandleFunc("/user/signin", logRequest(wrapJwt(jwtService, userService.SignIn))).
        Methods(http.MethodPost, http.MethodOptions)

    router.HandleFunc("/todo/lists", logRequest(jwtService.jwtAuth(users, listsService.CreateList))).
        Methods(http.MethodPost, http.MethodOptions)
    router.HandleFunc("/todo/lists/{list_id:[0-9]+}", logRequest(jwtService.jwtAuth(users, listsService.UpdateList))).
        Methods(http.MethodPut, http.MethodOptions)
    router.HandleFunc("/todo/lists/{list_id:[0-9]+}", logRequest(jwtService.jwtAuth(users, listsService.DeleteList))).
        Methods(http.MethodDelete, http.MethodOptions)
    router.HandleFunc("/todo/lists", logRequest(jwtService.jwtAuth(users, listsService.ReturnAllUserLists))).
        Methods(http.MethodGet, http.MethodOptions)

    router.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks", logRequest(jwtService.jwtAuth(users, listsService.CreateTask))).
        Methods(http.MethodPost, http.MethodOptions)
    router.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks/{task_id:[0-9]+}", logRequest(jwtService.jwtAuth(users, listsService.UpdateTask))).
        Methods(http.MethodPut, http.MethodOptions)
    router.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks/{task_id:[0-9]+}", logRequest(jwtService.jwtAuth(users, listsService.DeleteTask))).
        Methods(http.MethodDelete, http.MethodOptions)
    router.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks", logRequest(jwtService.jwtAuth(users, listsService.GetUserTasksInList))).
        Methods(http.MethodGet, http.MethodOptions)

    // c := cors.New(cors.Options{
    //     AllowedOrigins: []string{"http://localhost:8080"},
    //     AllowCredentials: true,
    // })

    // handler := c.Handler(router)

    srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		log.Println("Gracefully shutting down...")
		srv.Shutdown(ctx)
	}()

    log.Println("Server started, hit Ctrl+C to stop")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Server exited with error:", err)
	}
}
