package main

import (
	"fmt"
	_ "errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func getAdd(c *gin.Context) {
	xStr := c.Param("x");
	yStr := c.Param("y");

	x, err := strconv.Atoi(xStr);
	if (err != nil) {
		fmt.Println("Error:", err);
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});	
		return;
	}

	y, err := strconv.Atoi(yStr);
	if (err != nil) {
		fmt.Println("Error:", err);
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});	
		return;
	}

	c.IndentedJSON(http.StatusOK, gin.H{"result": x + y});
}

func getMul(c *gin.Context) {
	xStr := c.Param("x");
	yStr := c.Param("y");

	x, err := strconv.Atoi(xStr);
	if (err != nil) {
		fmt.Println("Error:", err);
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});	
		return;
	}

	y, err := strconv.Atoi(yStr);
	if (err != nil) {
		fmt.Println("Error:", err);
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});	
		return;
	}

	c.IndentedJSON(http.StatusOK, gin.H{"result": x * y});
}


func getSub(c *gin.Context) {
	xStr := c.Param("x");
	yStr := c.Param("y");

	x, err := strconv.Atoi(xStr);
	if (err != nil) {
		fmt.Println("Error:", err);
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});	
		return;
	}

	y, err := strconv.Atoi(yStr);
	if (err != nil) {
		fmt.Println("Error:", err);
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});	
		return;
	}

	c.IndentedJSON(http.StatusOK, gin.H{"result": x - y});
}

func getDiv(c *gin.Context) {
	xStr := c.Param("x");
	yStr := c.Param("y");

	x, err := strconv.Atoi(xStr);
	if (err != nil) {
		fmt.Println("Error:", err);
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});	
		return;
	}

	y, err := strconv.Atoi(yStr);
	if (err != nil) {
		fmt.Println("Error:", err);
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});	
		return;
	}

	if (y == 0) {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"result": 0});
		return;
	}

	c.IndentedJSON(http.StatusOK, gin.H{"result": x / y});
}

func main() {
	router := gin.Default();
	router.GET("/add/:x/:y", getAdd);
	router.GET("/sub/:x/:y", getSub);
	router.GET("/mul/:x/:y", getMul);
	router.GET("/div/:x/:y", getDiv);

	router.Run("localhost:8080");
}
