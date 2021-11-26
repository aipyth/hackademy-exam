package main

import (
	"errors"
)

var listsCounter int
var tasksCounter int

type Task struct {
    Id int
    Name string
    Description string
    Status string
}

type ToDoList struct {
    Id int
    Name string
    Tasks map[int]*Task
}

type UserToDo struct {
    Lists map[int]*ToDoList
}

type InMemoryToDoStorage struct {
    todo map[string]UserToDo
}

type CreateTaskParams struct {
    ListId int
    Name string
    Description string
}

type UpdateTaskParams struct {
    ListId int
    TaskId int
    Name string
    Status string
}

type DeleteTaskParams struct {
    ListId int
    TaskId int
}

type UserList struct {
    Id int
    Name string
}

func NewInMemoryToDoStorage() *InMemoryToDoStorage {
    return &InMemoryToDoStorage{
        todo: make(map[string]UserToDo),
    }
}

func (ts *InMemoryToDoStorage) CreateTask(user User, taskParams CreateTaskParams) (Task, error) {
    userTodos := ts.todo[user.Email]
    list, ok := userTodos.Lists[taskParams.ListId]
    if !ok {
        return Task{}, errors.New("no such list")
    }

    task := &Task{
        Id: tasksCounter,
        Name: taskParams.Name,
        Description: taskParams.Description,
        Status: "open",
    }

    tasksCounter++
    list.Tasks[task.Id] = task
    return *task, nil
}

func (ts *InMemoryToDoStorage) CreateList(user User, name string) (ToDoList, error) {
    userTodos, ok := ts.todo[user.Email]
    if !ok {
        userTodos = UserToDo{
            Lists: make(map[int]*ToDoList),
        }
        ts.todo[user.Email] = userTodos
    }
    list := &ToDoList{
        Id: listsCounter,
        Name: name,
        Tasks: make(map[int]*Task),
    }
    listsCounter++
    userTodos.Lists[list.Id] = list
    return *list, nil
}

func (ts *InMemoryToDoStorage) ChangeListName(user User, id int, newName string) (ToDoList, error) {
    userTodos := ts.todo[user.Email]
    list, ok := userTodos.Lists[id]
    if !ok {
        return ToDoList{}, errors.New("no such list")
    }

    
    // userTodos.Lists[id].Name = newName
    list.Name = newName
    return *list, nil
}

func (ts *InMemoryToDoStorage) DeleteList(user User, id int) error {
    userTodos := ts.todo[user.Email]
    _, ok := userTodos.Lists[id]
    if !ok {
        return errors.New("no such list")
    }
    delete(userTodos.Lists, id)
    return nil
}

func (ts *InMemoryToDoStorage) GetUserLists(user User) []UserList {
    userTodos := ts.todo[user.Email]
    lists := make([]UserList, 0)
    for _, v := range userTodos.Lists {
        lists = append(lists, UserList{
            Id: v.Id,
            Name: v.Name,
        })
    }
    return lists
}

func (ts *InMemoryToDoStorage) UpdateTask(user User, taskParams UpdateTaskParams) (Task, error) {
    userTodos := ts.todo[user.Email]
    list, ok := userTodos.Lists[taskParams.ListId]
    if !ok {
        return Task{}, errors.New("no such list")
    }

    task, ok := list.Tasks[taskParams.TaskId]
    if !ok {
        return Task{}, errors.New("no such task")
    }

    task.Name = taskParams.Name
    task.Status = taskParams.Status
    return *task, nil 
}

func (ts *InMemoryToDoStorage) DeleteTask(user User, taskParams DeleteTaskParams) error {
    userTodos := ts.todo[user.Email]
    list, ok := userTodos.Lists[taskParams.ListId]
    if !ok {
        return errors.New("no such list")
    }

    _, ok = list.Tasks[taskParams.TaskId]
    if !ok {
        return errors.New("no such task")
    }
    delete(list.Tasks, taskParams.TaskId)
    return nil
}

func (ts *InMemoryToDoStorage) GetListTasks(user User, listId int) ([]Task, error) {
    userTodos := ts.todo[user.Email]
    list, ok := userTodos.Lists[listId]
    if !ok {
        return []Task{}, errors.New("no such list")
    }
    tasks := make([]Task, 0)
    for _, task := range list.Tasks {
        tasks = append(tasks, *task)
    }
    return tasks, nil 
}
