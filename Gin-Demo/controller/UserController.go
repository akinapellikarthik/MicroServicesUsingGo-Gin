package controller

import (
	"GIN-DEMO/database"
	"GIN-DEMO/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	errHash := user.Hashpassword(user.Password)
	if errHash != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"server error": errHash.Error()})
		context.Abort()
		return
	}
	createdRecord := database.Instance.Create(&user)
	if createdRecord.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error in creating/Saving DB": createdRecord.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"Created User Successfully with details UserId": user.ID,
		"UserName":  user.UserName,
		"UserEmail": user.Email,
	})

}
