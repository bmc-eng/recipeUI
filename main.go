package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []string     `json:"steps"`
	Picture     string       `json:"imageURL"`
}

type Ingredient struct {
	Quantity string `json:"quantity"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

var recipes []Recipe

func IndexHandler(c *gin.Context) {
	recipes := make([]Recipe, 0)
	recipes = append(recipes, Recipe{
		Name:    "Burger",
		Picture: "/assets/images/burger.jpg",
	})
	recipes = append(recipes, Recipe{
		Name:    "Pizza",
		Picture: "/assets/images/pizza.jpg",
	})
	recipes = append(recipes, Recipe{
		Name:    "Tacos",
		Picture: "/assets/images/tacos.jpg",
	})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"recipes": recipes,
	})
}

func init() {
	recipes = make([]Recipe, 0)
	file, _ := os.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob(("templates/*"))
	router.GET("/", IndexHandler)
	router.Run()
}
