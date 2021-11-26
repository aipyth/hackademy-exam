package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ToDoRepository interface {
    CreateList(User, string) (ToDoList, error)
    ChangeListName(User, int, string) (ToDoList, error)
    DeleteList(User, int) error
    GetUserLists(User) []UserList
    CreateTask(User, CreateTaskParams) (Task, error)
    UpdateTask(User, UpdateTaskParams) (Task, error)
    DeleteTask(User, DeleteTaskParams) error
    GetListTasks(User, int) ([]Task, error)
}

type ListService struct {
    repository ToDoRepository
}

type CreateListRequestParams struct {
    Name string `json:"name"`
}

type CreateListResponse struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

type UpdateListRequestParams struct {
    Name string `json:"name"`
}

type UpdateListResponse struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

type CreateTaskRequestParams struct {
    TaskName string `json:"task_name"`
    Description string `json:"description"`
}

type UpdateTaskRequestParams struct {
    TaskName string `json:"task_name"`
    Status string `json:"status"` 
}

func (ls *ListService) CreateList(w http.ResponseWriter, r *http.Request, user User) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    params := &CreateListRequestParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

    if params.Name == "" {
		handleError(errors.New("empty name"), w)
		return
    }
    
    list, err := ls.repository.CreateList(user, params.Name)
    if err != nil {
		handleError(err, w)
		return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    encoder := json.NewEncoder(w)
    err = encoder.Encode(CreateListResponse{
        Id: list.Id,
        Name: list.Name,
    })
    if err != nil {
        handleError(err, w)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (ls *ListService) UpdateList(w http.ResponseWriter, r *http.Request, user User) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    vars := mux.Vars(r)
    listId, err := strconv.Atoi(vars["list_id"])
    if err != nil {
        handleError(errors.New("wrong list id"), w)
        return
    }

    params := &CreateListRequestParams{}
	err = json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

    if params.Name == "" {
		handleError(errors.New("empty name"), w)
		return
    }

    list, err := ls.repository.ChangeListName(user, listId, params.Name)
    if err != nil {
        handleError(err, w)
        return
    }

    err = json.NewEncoder(w).Encode(UpdateListResponse{
        Id: list.Id,
        Name: list.Name,
    })
    if err != nil {
        handleError(err, w)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func (ls *ListService) DeleteList(w http.ResponseWriter, r *http.Request, user User) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    vars := mux.Vars(r)
    listId, err := strconv.Atoi(vars["list_id"])
    if err != nil {
        handleError(errors.New("wrong list id"), w)
        return
    }

    err = ls.repository.DeleteList(user, listId)
    if err != nil {
        handleError(err, w)
        return
    }
    w.WriteHeader(http.StatusNoContent)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
}

func (ls *ListService) ReturnAllUserLists(w http.ResponseWriter, r *http.Request, user User) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    lists := ls.repository.GetUserLists(user)
    err := json.NewEncoder(w).Encode(lists)
    if err != nil {
       handleError(err, w) 
       return
    }
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func (ls *ListService) CreateTask(w http.ResponseWriter, r *http.Request, user User) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    vars := mux.Vars(r)
    listId, err := strconv.Atoi(vars["list_id"])
    if err != nil {
        handleError(errors.New("wrong list id"), w)
        return
    }
   
    params := &CreateTaskRequestParams{}
	err = json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

    if params.TaskName == "" {
		handleError(errors.New("empty task name"), w)
		return
    }

    task, err := ls.repository.CreateTask(user, CreateTaskParams{
        ListId: listId,
        Name: params.TaskName,
        Description: params.Description,
    })
    if err != nil {
        handleError(err, w)
        return
    }

    err = json.NewEncoder(w).Encode(task)
    if err != nil {
        handleError(err, w)
        return
    }
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func (ls *ListService) UpdateTask(w http.ResponseWriter, r *http.Request, user User) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    vars := mux.Vars(r)
    listId, err := strconv.Atoi(vars["list_id"])
    if err != nil {
        handleError(errors.New("wrong list id"), w)
        return
    }
    taskId, err := strconv.Atoi(vars["task_id"])
    if err != nil {
        handleError(errors.New("wrong task id"), w)
        return
    }
    
    params := &UpdateTaskRequestParams{}
	err = json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

    if params.TaskName == "" {
		handleError(errors.New("empty task name"), w)
		return
    }

    task, err := ls.repository.UpdateTask(user, UpdateTaskParams{
        ListId: listId,
        TaskId: taskId,
        Name: params.TaskName,
        Status: params.Status,
    })
    if err != nil {
        handleError(err, w)
        return
    }
    err = json.NewEncoder(w).Encode(task)
    if err != nil {
        handleError(err, w)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
} 

func (ls *ListService) DeleteTask(w http.ResponseWriter, r *http.Request, user User) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    vars := mux.Vars(r)
    listId, err := strconv.Atoi(vars["list_id"])
    if err != nil {
        handleError(errors.New("wrong list id"), w)
        return
    }
    taskId, err := strconv.Atoi(vars["task_id"])
    if err != nil {
        handleError(errors.New("wrong task id"), w)
        return
    }

    err = ls.repository.DeleteTask(user, DeleteTaskParams{
        ListId: listId,
        TaskId: taskId,
    })
    if err != nil {
        handleError(err, w)
        return
    }
    w.WriteHeader(http.StatusNoContent)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func (ls *ListService) GetUserTasksInList(w http.ResponseWriter, r *http.Request, user User) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
    if (r.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Headers", "Authorization")
        w.WriteHeader(204)
        return
    }
    vars := mux.Vars(r)
    listId, err := strconv.Atoi(vars["list_id"])
    if err != nil {
        handleError(errors.New("wrong list id"), w)
        return
    }
   
    tasks, err := ls.repository.GetListTasks(user, listId)
    if err != nil {
        handleError(err, w)
        return
    }

    err = json.NewEncoder(w).Encode(tasks)
    if err != nil {
        handleError(err, w)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
