package handlers

import (
	"context"

	"github.com/platinumscatter/simple_api/internal/taskService"
	"github.com/platinumscatter/simple_api/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasksUserUserId(ctx context.Context, request tasks.GetTasksUserUserIdRequestObject) (tasks.GetTasksUserUserIdResponseObject, error) {
	allTasks, err := h.Service.GetTasksByUserID(request.UserId)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksUserUserId200JSONResponse{}

	for _, tsk := range allTasks {
		isDone := tsk.IsDone
		task := tasks.Task{
			Id:        &tsk.ID,
			Task:      tsk.Task,
			IsDone:    &isDone,
			UserId:    tsk.UserID,
			CreatedAt: &tsk.CreatedAt,
			UpdatedAt: &tsk.UpdatedAt,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
    taskRequest := request.Body

    taskToCreate := taskService.Task{
        Task:   taskRequest.Task,
        IsDone: *taskRequest.IsDone,
        UserID: taskRequest.UserId,  // Add this line
    }

    createdTask, err := h.Service.CreateTask(taskToCreate)
    if err != nil {
        return nil, err
    }
    response := tasks.PostTasks201JSONResponse{
        Id:        &createdTask.ID,
        Task:      createdTask.Task,
        IsDone:    &createdTask.IsDone,
        UserId:    createdTask.UserID,  // Add this line
        CreatedAt: &createdTask.CreatedAt,
        UpdatedAt: &createdTask.UpdatedAt,
    }
    return response, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskToUpdate := taskService.Task{
		Task:   request.Body.Task,
		IsDone: *request.Body.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(request.Id, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskByID(request.Id)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}
