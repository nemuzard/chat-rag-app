package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
	"github.com/nemuzard/chat-rag-backend/internal/store"
)

// expected json body for user registration
type RegisterRequest struct{
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required, min=6,max=64"`
}

//json return after successful registration
type RegisterResponse struct{
	ID	unit `json:"id"`
	Username string `json:"username"`
}

func RegisterUser(c *gin.Context){
	var req RegisterRequest

	// 1. bind and validate json body 
	if err:=c.ShouldBindJson(&req); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid request body"
		})
		return
	}

	// 2. hash password using bcrypt
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"failed to hash password"
		})
		return
	}

	// 3. create a user model instance 
	user:=store.User{
		Username:req.Username,
		Password:string(hashedPassword),
	}

	// 4. save user in database
	if err:=store.DB.Create(&user).Error; err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"username might already be taken",
		})
		return 
	}

	// 5. return created user w/o password
	c.JSON(http.StatusCzreated,RegisterResponse{
		ID: user.ID,
		Username: user.Username,
	})

}