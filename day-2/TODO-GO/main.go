package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID			string	`json:"id"`
	Item		string	`json:"item"`
	Completede	bool	`json:"completed"`
}


var todos = []todo{
	{ID: "1", Item: "Clean Room", Completede: false},
	{ID: "2", Item: "Read Book", Completede: false},
	{ID: "3", Item: "Record Video", Completede: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id string ) (*todo, error) {
	for _ , t:= range todos {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, errors.New("todo not found")	
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo Not Found"})
		return 
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context){
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo Not Found"})
		return 
	}

	todo.Completede = !todo.Completede

	context.IndentedJSON(http.StatusOK, todo)

}


func main (){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}