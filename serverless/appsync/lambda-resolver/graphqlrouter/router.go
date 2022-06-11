package graphqlrouter

import (
	"context"
	"fmt"
	"myAppSyncApp-API/serverless/appsync"

	"github.com/mitchellh/mapstructure"
)

type TodoInterface interface {
	CreateTodoDB(in appsync.NewTodo) (out appsync.Todo, err error)
	//UpdateTodo(pk string, sk string, done bool) (out appsync.Todo, err error)
}

type GraphQLRouter struct {
	TodoInterface TodoInterface
}

// Event for an incoming GraphQL request
// Matches the request_template in aws_appsync_resolver
type Event struct {
	Field     string                 `json:"field"`
	Arguments map[string]interface{} `json:"arguments"`
	Identity  string                 `json:"identity"`
}

func NewGraphQLRouter(
	todointerface TodoInterface,
) GraphQLRouter {
	return GraphQLRouter{
		TodoInterface: todointerface,
	}
}

// Route routes graphql requests to golang code.
func (g GraphQLRouter) Route(ctx context.Context, req Event) (res interface{}, err error) {

	fmt.Printf(" args for req %+v", req)
	appError := fmt.Errorf("Default error")
	switch req.Field {
	// You can add multiple switch statements to handle each field
	/*case "postTodoUpdate":
	var args appsync.UpdateTodoInput

	fmt.Printf(" Identiy %+v", req.Identity)

	mapstructure.Decode(req.Arguments["input"], &args)
	fmt.Printf(" args %+v", args)

	res, appError = g.TodoInterface.UpdateTodo(req.Identity, *args.ID, *args.Done)

	if appError != nil {
		err = fmt.Errorf("error while handling request: %s, err: %s", req.Field, appError.Error())
	}
	break*/
	case "createtodo":
		var args appsync.NewTodo

		fmt.Println(" -- ")
		fmt.Printf(" Input is %+v", req.Arguments["input"])
		fmt.Println("--")

		mapstructure.Decode(req.Arguments["input"], &args)
		//fmt.Printf(" Error in decoding %+v", err)
		fmt.Printf(" args %+v", args)

		fmt.Printf(" Identify %+v", req.Identity)

		args.UserID = req.Identity
		res, appError = g.TodoInterface.CreateTodoDB(args)

		if appError != nil {
			err = fmt.Errorf("error while handling request: %s, err: %s", req.Field, appError.Error())
		}
		break

	}

	return res, nil

}
