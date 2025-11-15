package main


import(
	"log"
	"net/http"
	"github.com/gin-gonic/gin" 
)

func main(){
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