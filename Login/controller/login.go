package controller

import (
	"encoding/json"
	"fmt"

	//"C:/Users/Dell/Desktop/GOProject/model"
	"GOProject/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"io/ioutil"

	//"github.com/mongodb/mongo-go-driver/bson"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) UserRegister(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res model.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	//collection, err := db.GetDBCollection()

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var foundUser model.User
	//err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	// if err := uc.session.DB("go-web-dev-db").C("users").Find(bson.M{"uname": user.Username}).One(&foundUser); err != nil {
	// 	fmt.Println("Error we arer in 111")
	// 	fmt.Println(err)
	// 	// w.WriteHeader(404)
	// 	// return
	// }

	err = uc.session.DB("go-web-dev-db").C("users").Find(bson.M{"uname": user.Username}).One(&foundUser)

	if err != nil {
		fmt.Println("I am inside post", err)
		if err.Error() == "not found" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			user.Password = string(hash)

			// _, err = collection.InsertOne(context.TODO(), user)
			// if err != nil {
			// 	res.Error = "Error While Creating User, Try Again"
			// 	json.NewEncoder(w).Encode(res)
			// 	return
			// }
			user.Id = bson.NewObjectId()
			err2 := uc.session.DB("go-web-dev-db").C("users").Insert(user)
			fmt.Println("error if there", err2)
			res.Result = "Registration Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Username already Exists!!"
	json.NewEncoder(w).Encode(res)
	return

}

func (uc UserController) UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	var foundUser model.User
	var res model.ResponseResult

	//err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	err = uc.session.DB("go-web-dev-db").C("users").Find(bson.M{"username": user.Username}).One(&foundUser)

	if err != nil {
		res.Error = "Invalid username"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))

	if err != nil {
		res.Error = "Invalid password"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  foundUser.Username,
		"firstname": foundUser.FirstName,
		"lastname":  foundUser.LastName,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		json.NewEncoder(w).Encode(res)
		return
	}

	foundUser.Token = tokenString
	foundUser.Password = ""

	json.NewEncoder(w).Encode(foundUser)

}
