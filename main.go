package main

import (
	"database/sql"
	"fmt"
	"ginCLI/mysql"
	"ginCLI/uservalidation"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

const (
	conferenceTickets uint = 250
	conferenceName         = "Stand-up comedy"
	comedian               = "SUBBU"
)

var (
	remainingTickets uint
	count            uint
	dataBase         *sql.DB
	err              error
	userData         UserData
)

type UserData struct {
	FirstName     string `form:"FirstName" binding:"required"`
	LastName      string `form:"LastName" binding:"required"`
	Email         string `form:"Email" binding:"required"`
	TicketsBooked uint   `form:"TicketsBooked" binding:"required,min=1"` //the given min is minimum required as it binding it's mandatary to be more than 1
}

func init() {
	dataBase, err = sql.Open("mysql", mysql.DataBaseURL)
	if err != nil {
		fmt.Println("DataBase error", err)
	}
	err = dataBase.QueryRow("SELECT SUM(usertickets) FROM customer").Scan(&count)
	if err != nil {
		fmt.Println("Query Row error", err)
	}
	remainingTickets = conferenceTickets - count
}
func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", nil)
	})
	router.POST("/book", func(c *gin.Context) {
		go func() {
			bookTicketHandler(c)
		}()
		c.Redirect(http.StatusSeeOther, "/welcome")
	})
	router.Run(":8080")
	defer dataBase.Close()
}

func bookTicketHandler(c *gin.Context) {
	c.ShouldBind(&userData)
	c.JSON(200, gin.H{
		"FirstName":     userData.FirstName,
		"LastName":      userData.LastName,
		"Email":         userData.Email,
		"TicketsBooked": userData.TicketsBooked,
	})
	isValidName, isValidTicketNumber := uservalidation.ValidateUserInput(userData.FirstName, userData.LastName, userData.TicketsBooked, remainingTickets)
	if isValidName && isValidTicketNumber {
		if remainingTickets >= userData.TicketsBooked {
			remainingTickets -= userData.TicketsBooked
			statement, err := dataBase.Prepare("INSERT INTO customer(firstName, lastName, email, userTickets) values(?,?,?,?)")
			if err != nil {
				fmt.Println("error in statement object", err)
				return
			}
			defer statement.Close()
			_, err = statement.Exec(userData.FirstName, userData.LastName, userData.Email, userData.TicketsBooked)
			if err != nil {
				fmt.Println(err)
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Booking successful"})
			SendTicket(userData)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough tickets available"})
		}
	} else {
		errorMessages := make([]string, 0)
		if !isValidName {
			errorMessages = append(errorMessages, "First name or last name is too short")
		}
		if !isValidTicketNumber {
			errorMessages = append(errorMessages, "Number of tickets is invalid")
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": strings.Join(errorMessages, ", ")})
	}
}

func SendTicket(userData UserData) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userData.TicketsBooked, userData.FirstName, userData.LastName)
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, userData.Email)
}

//as i integrated usertickets must be more than 0. so i can avoid it here it with html page so i didn't have to verify it in middleware
