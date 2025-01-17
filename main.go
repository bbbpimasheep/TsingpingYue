package main

import (
	"Serenesongserver/config"
	"Serenesongserver/controllers"
	"Serenesongserver/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	utils.SetupCronJobs()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Login related APIs
	router.POST("/login", controllers.Login)

	// Collection related APIs
	router.GET("/getAllCollections", controllers.GetAllColletions)
	router.GET("/getAllColletionItems", controllers.GetAllColletionItems)
	router.POST("/createCollection", controllers.CreateCollection)
	router.POST("/deleteCollection", controllers.DeleteCollection)
	router.POST("/addToCollection", controllers.AddToCollection)
	router.POST("/removeFromCollection", controllers.RemoveFromCollection)
	router.POST("/modifyCollectionComment", controllers.ModifyCollectionComment)
	router.GET("/getCollectionItemCount", controllers.GetCollectionItemCount)
	router.GET("/getCollectionItem", controllers.GetCollectionItem)

	// Recommendation related APIs
	router.GET("/recommendCi", controllers.RecommendCi)
	router.GET("/recommendPic", controllers.RecommendPic)

	// Searching related APIs
	router.GET("/search", controllers.SearchRouter)

	// Composing related APIs
	router.GET("/getRhymes", controllers.GetRhymes)
	router.GET("/getYunshu", controllers.GetYunshu)
	router.GET("/getFormat", controllers.GetFormat)
	router.POST("/putIntoDrafts", controllers.PutIntoDrafts)
	router.DELETE("/delDraft", controllers.DelDraft)
	router.POST("/turnToFormal", controllers.TurnToFormal)
	router.POST("/modifyDraft", controllers.ModifyDraft)
	router.POST("/modifyWork", controllers.ModifyWork)
	router.GET("/getMyWorks", controllers.GetMyWorks)
	router.GET("/getCiById", controllers.GetCiById)

	// User info related APIs
	router.POST("/saveUserInfo", controllers.SaveUserInfo)
	router.POST("/changePrivacy", controllers.ChangePrivacy)
	router.GET("/getDynamics", controllers.GetDynamics)
	router.GET("/getCollections", controllers.GetCollections)
	router.GET("/getSubscribers", controllers.GetSubscribers)
	router.GET("/getSubscribedTo", controllers.GetSubscribedTo)
	router.GET("/getPublicWorks", controllers.GetPublicWorks)
	router.GET("/getWorks", controllers.GetWorks)
	router.GET("/getUserInfo", controllers.GetUserInfo)
	router.GET("/getPersonalID", controllers.GetPersonalID)
	router.GET("/getUserInfoText", controllers.GetUserInfoText)

	// Community related APIs
	router.POST("/publishDynamic", controllers.PublishDynamic)
	router.POST("/commentPost", controllers.CommentPost)
	router.POST("/withdrawComment", controllers.WithdrawComment)
	router.POST("/withdrawPost", controllers.WithdrawPost)
	router.POST("/likePost", controllers.LikePost)
	router.POST("/withdrawLike", controllers.WithdrawLike)
	router.GET("/getRandomPosts", controllers.GetRandomPosts)
	router.GET("/getFollowingPosts", controllers.GetFollowingPosts)

	// Message related APIs
	router.GET("/getMessagesIGet", controllers.GetMessagesIGet)
	router.GET("/getMessagesISend", controllers.GetMessagesISend)
	router.POST("/sendMessage", controllers.SendMessage)
	router.POST("/subscribeOthers", controllers.SubscribeOthers)
	router.GET("/searchUserByName", controllers.SearchUserByName)
	router.GET("/getMessageById", controllers.GetMessageById)

	router.Run("0.0.0.0:8080")
}
