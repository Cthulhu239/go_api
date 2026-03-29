package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}


var books = []book{
    {ID:"1",Title:"1984",Author:"owell",Quantity:2},
}   

func getBooks(c *gin.Context){
      c.IndentedJSON(http.StatusOK,books)  
}

func createBook(c *gin.Context){
    var newBook book
    if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated,newBook)
}

func checkoutBook(c *gin.Context){
	id,ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Missing id query parameter"})
	}

	book,err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Book not available"})
		return
	}

	book.Quantity--
	c.IndentedJSON(http.StatusOK,book)
}


func bookById(c *gin.Context){
	id := c.Param("id")
	book,err := getBookById(id)
    if err != nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK,book)
}

func returnBook(c *gin.Context){
	id,ok := c.GetQuery("id")
	
	if !ok {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"query paramter not found"})
		return
	}

	book,err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"id does not exist"})
		return
	}
	
	book.Quantity++
    c.IndentedJSON(http.StatusOK,book)
}

func getBookById(id string) (*book,error){
	for i,b := range books{
       if b.ID == id{
		return &books[i],nil
	   }
	}

	return nil,errors.New("book not found")
}

func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.POST("/books",createBook)
	router.GET("/books/:id",bookById)
	router.PATCH("/checkout",checkoutBook)
	router.PATCH("/return",returnBook)
	router.Run("localhost:8080")
}