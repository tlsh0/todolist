package handlers

type successResponse struct {
	Success string `json:"success"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type userCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type newTaskRequest struct {
	Name    string `json:"name"`
	DueTime string `json:"due_time"`
}

type deleteTaskRequest struct {
	ID int64 `json:"id"`
}

type updateTaskRequest struct {
	ID   int64 `json:"id"`
	Done bool  `json:"status"`
}
