package routes

import (
	"fmt"
	"github.com/crud"
	"github.com/gofiber/fiber/v2"
	"github.com/models"
	"github.com/util"
)

func createTask(ctx *fiber.Ctx) error {
	var task *models.Task
	if err := ctx.BodyParser(&task); err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 400)
		return ctx.SendStatus(400)
	}
	if err := crud.CreateTask(task); err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 502)
		return ctx.SendString("Unable to create task")
	}
	util.LogRequest(ctx.Method(), ctx.Path(), 200)
	return ctx.SendStatus(200)
}

func listTasks(ctx *fiber.Ctx) error {
	tasks, err := crud.ListTasks()
	if err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 502)
		ctx.SendStatus(502)
		return ctx.SendString("Unable to list tasks")
	}
	util.LogRequest(ctx.Method(), ctx.Path(), 200)
	return ctx.JSON(tasks)
}

func deleteTasks(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 403)
		ctx.SendString("Unable to Parse ID")
		return ctx.SendStatus(403)
	}

	err = crud.DeleteTask(int32(id))
	if err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 502)
		ctx.SendString("Unable to delete task")
		return ctx.SendStatus(502)
	}
	util.LogRequest(ctx.Method(), ctx.Path(), 200)
	ctx.SendString("Successfully deleted the task")
	return ctx.SendStatus(200)
}

func updateTask(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 403)
		ctx.SendString("Unable to Parse ID")
		return ctx.SendStatus(403)
	}
	var task *models.Task
	if err := ctx.BodyParser(&task); err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 400)
		return ctx.SendStatus(400)
	}
	fmt.Println("status", task.Status)
	err = crud.UpdateTask(int32(id), task)
	if err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 502)
		ctx.SendString("Unable to update task")
		return ctx.SendStatus(502)
	}
	util.LogRequest(ctx.Method(), ctx.Path(), 200)
	ctx.SendString("Successfully updated the task")
	return ctx.SendStatus(200)
}

func searchTasks(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
	fmt.Println("this is key",key)
	tasks, err := crud.SearchTasks(key)
	if err != nil {
		util.LogRequest(ctx.Method(), ctx.Path(), 502)
		ctx.SendStatus(502)
		return ctx.SendString("Unable to list tasks")
	}
	util.LogRequest(ctx.Method(), ctx.Path(), 200)
	return ctx.JSON(tasks)
}

func add(a int, b int) int {
	a = a + b
}

