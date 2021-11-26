package main

import "testing"

func TestInMemoryTodoStorage(t *testing.T) {
	t.Run("storage is created", func(t *testing.T) {
		storage := NewInMemoryToDoStorage()

		if storage == nil {
			t.Errorf("storage is nil")
		}
        if storage.todo == nil {
            t.Error("storage map is nil")
        }
	})


    t.Run("creates list", func(t *testing.T) {
		storage := NewInMemoryToDoStorage()

        user := User{
            Email: "test",
        }

        listName := "testList"
       
        list, err := storage.CreateList(user, listName)
        if err != nil {
            t.Error(err)
        }

        if list.Name != listName {
            t.Error("returned list has wrong name")
        }

        lss, ok := storage.todo[user.Email]
        if !ok {
            t.Fatal("could not get user todos")
        }

        l, ok := lss.Lists[list.Id]
        if !ok {
            t.Fatal("could not get list by listId")
        }
        
        switch {
        case list.Id != l.Id:
            t.Error("ids do not match")        
        case list.Name != l.Name:
            t.Error("names do not match")        
        }

    })

    t.Run("changes list name", func(t *testing.T) {
		storage := NewInMemoryToDoStorage()
        user := User{
            Email: "test",
        }
        listName := "testList"
        l, err := storage.CreateList(user, listName)
        if err != nil {
            t.Error(err)
        }

        toChangeTo := "renewedTestList"
        list, err := storage.ChangeListName(user, l.Id, toChangeTo)
        ls := storage.todo[user.Email].Lists[list.Id]
        if ls.Name != toChangeTo {
            t.Error("list name didn't change")
        }
    })

    t.Run("deletes list", func(t *testing.T) {
		storage := NewInMemoryToDoStorage()
        user := User{
            Email: "test",
        }
        listName := "testList"
        l, err := storage.CreateList(user, listName)
        if err != nil {
            t.Error(err)
        }
        
        err = storage.DeleteList(user, l.Id)
        if err != nil {
            t.Error(err)
        }

        _, ok := storage.todo[user.Email].Lists[l.Id]
        if ok {
            t.Error("list is not deleted")
        }
    })

    t.Run("creates task", func(t *testing.T) {
		storage := NewInMemoryToDoStorage()
        user := User{
            Email: "test",
        }
        listName := "testList"
        l, err := storage.CreateList(user, listName)
        if err != nil {
            t.Error(err)
        }

        taskParams := CreateTaskParams{
            ListId: l.Id, 
            Name: "testTask",
            Description: "desc task",
        }
        task, err := storage.CreateTask(user, taskParams)
        if err != nil {
            t.Error(err)
        }

        tsk, ok := storage.todo[user.Email].Lists[l.Id].Tasks[task.Id]
        if !ok {
            t.Error("task is not created")
        }

        switch {
        case task.Name != tsk.Name || tsk.Name != taskParams.Name:
            t.Error("names do not match")
        case tsk.Description != taskParams.Description:
            t.Error("task desc do not match")
        }
    })

    t.Run("updates task", func(t *testing.T) {
		storage := NewInMemoryToDoStorage()
        user := User{
            Email: "test",
        }
        listName := "testList"
        l, err := storage.CreateList(user, listName)
        if err != nil {
            t.Error(err)
        }

        taskParams := CreateTaskParams{
            ListId: l.Id, 
            Name: "testTask",
            Description: "desc task",
        }
        task, err := storage.CreateTask(user, taskParams)
        if err != nil {
            t.Error(err)
        }

        updateParams := UpdateTaskParams{
            ListId: l.Id,
            TaskId: task.Id,
            Name: "newTaskName",
            Status: "completed",
        }
        _, err = storage.UpdateTask(user, updateParams)
        if err != nil {
            t.Error(err)
        }

        tsk, ok := storage.todo[user.Email].Lists[l.Id].Tasks[task.Id]
        if !ok {
            t.Error("task is not created")
        }

        switch {
        case updateParams.Name != tsk.Name:
            t.Error("names is not updated")
        case updateParams.Status != tsk.Status:
            t.Error("status is not updated")
        }
    })

    t.Run("deletes task", func(t *testing.T) {
		storage := NewInMemoryToDoStorage()
        user := User{
            Email: "test",
        }
        listName := "testList"
        l, err := storage.CreateList(user, listName)
        if err != nil {
            t.Error(err)
        }

        taskParams := CreateTaskParams{
            ListId: l.Id, 
            Name: "testTask",
            Description: "desc task",
        }
        task, err := storage.CreateTask(user, taskParams)
        if err != nil {
            t.Error(err)
        }

        deleteParams := DeleteTaskParams{
            ListId: l.Id,
            TaskId: task.Id,
        }
        err = storage.DeleteTask(user, deleteParams)
        if err != nil {
            t.Error(err)
        }

        _, ok := storage.todo[user.Email].Lists[l.Id].Tasks[task.Id]
        if ok {
            t.Error("task is not created")
        }
    })
}
