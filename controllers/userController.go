package controllers

import (
    "context"
    "fmt"
    "log"
    "strconv"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    helper "cfw/helpers"
    "cfw/models"
    "cfw/helpers"
    "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        var user models.User

        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        validationErr := validate.Struct(user)
        if validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()}
            return
        }
        count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking email."})
        }
        count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.h{"error": "error while checking phone"})
        }

        if count > 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "this email/phone already exists"})
        }
    }
}

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        userId := c.Param("user_id")

        if err := helper.MatchUserTypeToUid(c, userId); err != nil {
            c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
            return
        }
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

        var user models.User
        err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
        defer cancel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOk, user)
    }
}


