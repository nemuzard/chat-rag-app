package main


import(
	"log"
	"net/http"
	"github.com/gin-gonic/gin" 
	"github.com/nemuzard/chat-rag-backend/internal/store"
	"github.com/nemuzard/chat-rag-backend/internal/handlers"

)

func main(){

	store.InitDB()

	if err := store.DB.AutoMigrate(
		&store.User{},
		&store.Conversation{},
		&store.ConversationMember{},
		&store.Message{},
	); err != nil{
		log.Fatalf("failed to migrate database: %v", err)
	}


	// create a gin router
	r:=gin.Default()
	// register a get route -/health
	r.GET("/health",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{

			"status":"ok",
		})
	})
	//When sends a POST request to `/users/register`, it invokes `handlers.RegisterUser`.
	r.POST("/users/register", handlers.RegisterUser)
	// start on port 8000
	if err:=r.Run(":8000"); err!=nil{
		log.Fatalf("server failed to start: %v",err)
	}
}