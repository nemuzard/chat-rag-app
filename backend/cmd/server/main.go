package main


import(
	"log"
	"net/http"
	"github.com/gin-gonic/gin" 
	"github.com/nemuzard/chat-rag-backend/internal/store"
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
	// start on port 8080
	if err:=r.Run(":8000"); err!=nil{

		log.Fatalf("server failed to start: %v",err)
	}
}