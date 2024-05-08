package controllers

import (
	"log"
	"net/http"

	"github.com/edwinhuish/go-rest-template/internal/api/gin2"
	models "github.com/edwinhuish/go-rest-template/internal/models/tasks"
	"github.com/edwinhuish/go-rest-template/internal/repos"
)

type TaskController struct {
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

// GetTaskById godoc
//
//	@Summary		Retrieves task based on given ID
//	@Description	get Task by ID
//	@Produce		json
//	@Param			id	path		integer	true	"Task ID"
//	@Success		200	{object}	tasks.Task
//	@Router			/api/tasks/{id} [get]
//	@Security		Authorization Token
func (ctrl *TaskController) Find(c *gin2.Context) {
	s := repos.GetTaskRepository()
	id := c.Param("id")
	if task, err := s.Get(id); err != nil {
		c.Status(http.StatusNotFound)
		c.Resp().Fail("task not found")
		log.Println(err)
	} else {
		c.Resp().Success(task)
	}
}

// GetTasks godoc
//
//	@Summary		Retrieves tasks based on query
//	@Description	Get Tasks
//	@Produce		json
//	@Param			taskname	query	string	false	"Taskname"
//	@Param			firstname	query	string	false	"Firstname"
//	@Param			lastname	query	string	false	"Lastname"
//	@Success		200			{array}	[]tasks.Task
//	@Router			/api/tasks [get]
//	@Security		Authorization Token
func (ctrl *TaskController) List(c *gin2.Context) {
	s := repos.GetTaskRepository()
	var q models.Task
	_ = c.Bind(&q)
	if tasks, err := s.Query(&q); err != nil {
		c.Status(http.StatusNotFound)
		c.Resp().Fail("tasks not found")
		log.Println(err)
	} else {
		c.Resp().Success(tasks)
	}
}

func (ctrl *TaskController) Create(c *gin2.Context) {
	s := repos.GetTaskRepository()
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if err := s.Add(&taskInput); err != nil {
		c.Status(http.StatusBadRequest)
		c.Resp().Fail(err)
		log.Println(err)
	} else {
		c.Resp().Success(taskInput)
	}
}

func (ctrl *TaskController) Update(c *gin2.Context) {
	s := repos.GetTaskRepository()
	id := c.Params.ByName("id")
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if _, err := s.Get(id); err != nil {
		c.Status(http.StatusNotFound)
		c.Resp().Fail("task not found")
		log.Println(err)
	} else {
		if err := s.Update(&taskInput); err != nil {
			c.Status(http.StatusNotFound)
			c.Resp().Fail(err)
			log.Println(err)
		} else {
			c.Resp().Success(taskInput)
		}
	}
}

func (ctrl *TaskController) Delete(c *gin2.Context) {
	s := repos.GetTaskRepository()
	id := c.Params.ByName("id")
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if task, err := s.Get(id); err != nil {
		c.Status(http.StatusNotFound)
		c.Resp().Fail("task not found")
		log.Println(err)
	} else {
		if err := s.Delete(task); err != nil {
			c.Status(http.StatusNotFound)
			c.Resp().Fail(err)
			log.Println(err)
		} else {
			c.Resp().Success()
		}
	}
}
