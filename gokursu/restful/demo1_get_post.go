package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Hata bulundu", err)

	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var Todo todo
	json.Unmarshal(bodyBytes, &Todo)
	fmt.Println(Todo)
}

func Demo2() {

	todo := todo{1, 2, "Alışveriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos/", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	bytes.NewBuffer(jsonTodo)

	if err != nil {
		fmt.Println("Hata bulundu", err)

	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}
