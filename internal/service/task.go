package service

import (
	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/Njrctr/DeNet_test/internal/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task models.TaskCreate) (int, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) CompleteTask(userId, taskId int) error {
	return s.repo.CompleteTask(userId, taskId)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) ReferrerCode(userId int, refCode string) error {
	return s.repo.ReferrerCode(userId, refCode)
}
