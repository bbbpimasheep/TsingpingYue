package controllers

import (
	// "Serenesongserver/models"
	"Serenesongserver/services"
	"Serenesongserver/utils"
	// "go/token"

	// "encoding/json"
	"net/http"
	"strconv"
	// "encoding/json"
	// "fmt"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func GetDynamics(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	if user_id == "" || token == "" {
		utils.HandleError(c, http.StatusBadRequest, utils.ErrMsgInvalidParams, nil)
		return
	}
	services.ReturnDynamics(c, user_id, token)
}

func GetCollections(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	if user_id == "" || token == "" {
		utils.HandleError(c, http.StatusBadRequest, utils.ErrMsgInvalidParams, nil)
		return
	}
	services.ReturnCollections(c, user_id, token)
}

func GetSubscribers(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		utils.HandleError(c, http.StatusBadRequest, utils.ErrMsgInvalidParams, nil)
		return
	}
	services.ReturnSubscribers(c, token)
}

func GetSubscribedTo(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		utils.HandleError(c, http.StatusBadRequest, utils.ErrMsgInvalidParams, nil)
		return
	}
	services.ReturnSubscribedTo(c, token)
}

func GetPublicWorks(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	if user_id == "" || token == "" {
		utils.HandleError(c, http.StatusBadRequest, utils.ErrMsgInvalidParams, nil)
		return
	}
	services.ReturnPublicWorks(c, user_id, token)
}

func GetWorks(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	if user_id == "" || token == "" {
		utils.HandleError(c, http.StatusBadRequest, utils.ErrMsgInvalidParams, nil)
		return
	}
	services.ReturnWorks(c, user_id, token)
}

func GetAvatar(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	if user_id == "" || token == "" {
		utils.HandleError(c, http.StatusBadRequest, utils.ErrMsgInvalidParams, nil)
		return
	}
	services.ReturnAvatar(c, user_id, token)
}

func ChangePrivacy(c *gin.Context) {
	work_id := c.Query("work_id")
	token := c.Query("token")
	privacy := c.Query("privacy")
	if work_id == "" || token == "" || privacy == "" {
		utils.HandleError(c, http.StatusBadRequest, utils.ErrMsgInvalidParams, nil)
		return
	}
	// Parse the 'privacy' parameter to a boolean
    is_public, err := strconv.ParseBool(privacy)
    if err != nil {
        utils.HandleError(c, http.StatusBadRequest, "Invalid privacy value. Must be a boolean.", nil)
        return
    }
	services.ChangePrivacy(c, work_id, token, is_public)
}