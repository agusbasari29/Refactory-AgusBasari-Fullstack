package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DataA struct {
	Posts []Posts `json:"posts"`
}

type Posts struct {
	UserId int     `json:"userId"`
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Users  []Users `json:"user"`
}

type DataB struct {
	Users []Users `json:"users"`
}
type Users struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Address  []Address `json:"address"`
}
type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     []Geo  `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func main() {
	posts := `{"posts":` + fetchData(`https://jsonplaceholder.typicode.com/posts`) + `}`
	users := `{"users":` + fetchData(`https://jsonplaceholder.typicode.com/users`) + `}`

	postsObject := DataA{}
	err := json.Unmarshal([]byte(posts), &postsObject)
	if err != nil {
		fmt.Println(err)
	}

	usersObject := DataB{}
	errr := json.Unmarshal([]byte(users), &usersObject)
	if errr != nil {
		fmt.Println(err)
	}

	fmt.Println(postsObject)
}

func fetchData(uri string) string {
	response, err := http.Get(uri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responsData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responsData)
}
