package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rama-kairi/blog-api-golang-gin/db"
	"github.com/rama-kairi/blog-api-golang-gin/models"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type userController struct {
	db   *mongo.Client
	coll *mongo.Collection
}

func NewUserController() *userController {
	return &userController{
		db:   db.MongoClient,
		coll: db.MongoClient.Database("ecomm").Collection("users"),
	}
}

// Get all Users
func (u userController) GetAll(c *gin.Context) {
	var users []models.User
	ctx := context.TODO()
	cursor, err := u.coll.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		utils.Response(c, http.StatusBadRequest, nil, "Error getting users")
		return
	}

	if err = cursor.All(ctx, &users); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting users")
		return
	}

	utils.Response(c, http.StatusOK, users, "users found")
}

// Get a user
func (u userController) Get(c *gin.Context) {
	// Get user id from url
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting user")
		return
	}

	// Get the user from the database
	var user models.User
	ctx := context.TODO()
	err = u.coll.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting user")
		return
	}

	utils.Response(c, http.StatusNotFound, user, "user found")
}

// Create a User
func (u userController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error creating user")
		return
	}

	// Set the user id
	// user.ID = primitive.NewObjectID()

	// Create the user in the database
	ctx := context.TODO()
	res, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	// Marshal the user into json
	utils.Response(c, http.StatusCreated, res.InsertedID, "user created successfully")
}

// Delete a user
func (u userController) Delete(c *gin.Context) {
	// Get user id from url
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting user")
		return
	}
	// Delete the user from the database
	ctx := context.TODO()
	res, err := u.coll.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting user")
		return
	}

	if res.DeletedCount == 0 {
		utils.Response(c, http.StatusNotFound, nil, "user not found")
		return
	}

	utils.Response(c, http.StatusNoContent, nil, "user Deleted")
}

// Update a user
func (u userController) Update(c *gin.Context) {
	// Get user id from url
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting user")
		return
	}
	// Get the user payload from the request
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error updating user")
		return
	}

	// Update the user in the database
	ctx := context.TODO()
	res, err := u.coll.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": user})
	log.Println(res)
	if err != nil {
		log.Println(err)
		utils.Response(c, http.StatusInternalServerError, nil, "Error updating user")
		return
	}

	if res.MatchedCount == 0 {
		utils.Response(c, http.StatusNotFound, nil, "user not found")
		return
	}

	// return the updated user
	utils.Response(c, http.StatusOK, nil, "user updated successfully")
}
