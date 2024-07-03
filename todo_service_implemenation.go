package main

import (
	context "context"
)

type TodoServiceServerImpelementation struct {
	UnimplementedTodoServiceServer
}

func (s *TodoServiceServerImpelementation) Get(ctx context.Context, in *GetTodoRequest) (*GetTodoResponse, error) {

	var res *GetTodoResponse = &GetTodoResponse{}
	res.Todo = &Todo{
		Id:          in.Id,
		Title:       "Todo Title",
		Description: "Todo Description",
		Done:        false,
	}
	return res, nil
}
