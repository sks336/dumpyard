package main

import (
    "fmt"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "github.com/gin-gonic/gin"
    //"encoding/json"
)
type result_data struct{
	Id int
	First_Name string
	Last_Name string
}
type Test struct {
    result_data
    count int
}

func main() {
	router := gin.Default()

	// GET a person detail
	router.GET("/persons/", func(c *gin.Context) {
		apiUrl := "http://localhost:3000/"
		resource := "/persons/"


		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := u.String() // 'https://api.com/user/'

		client := &http.Client{}
		r, _ := http.NewRequest("GET", urlStr, nil) // URL-encoded payload

		resp, _ := client.Do(r)
		fmt.Println(resp)
		//println(test.result_data.First_Name)
	})

	// POST new person details
	router.POST("/person", func(c *gin.Context) {
		apiUrl := "http://localhost:3000/"
    	resource := "/person/"
    	data := url.Values{}
    	first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")
    	data.Set("first_name", first_name)
    	data.Add("last_name", last_name)

    	u, _ := url.ParseRequestURI(apiUrl)
    	u.Path = resource
    	urlStr := u.String() // 'https://api.com/user/'

    	client := &http.Client{}
    	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
    	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
    	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    	resp, _ := client.Do(r)
    	fmt.Println(resp.Status,resp)
    	})
	router.Run(":4000")
}