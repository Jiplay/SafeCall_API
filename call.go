package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProfiler(c *gin.Context) {
	url := "http://profiler:8081/testZMQServer"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		// return "false"
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		// return "false"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		// return "false"
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	// return string(body)
	c.JSON(200, gin.H{
		"success": "resp",
	})
}

func sendCall(c *gin.Context) {
	a := c.Param("userID")
	b := c.Param("dest")

	resp := sendCallService(a, b)

	c.JSON(200, gin.H{
		"success": resp,
	})
}
