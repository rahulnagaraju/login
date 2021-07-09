package main

import (
	"GOProject/controller"
	"fmt"
	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
)

func init() {
	rnd = renderer.New(
		renderer.Options{
			ParseGlobPattern: "views/*.html",
		},
	)
}

var rnd *renderer.Render

func main() {

	r := httprouter.New()
	uc := controller.NewUserController(getSession())

	// r.GET("/users", uc.GetAllUsers)
	// r.GET("/user/:id", uc.GetUser)
	// r.POST("/user", uc.CreateUser)
	// r.DELETE("/user/:id", uc.DeleteUser)
	http.HandleFunc("/loginn", showLogin)
	r.POST("/login", uc.UserLogin)
	r.POST("/register", uc.UserRegister)
	http.HandleFunc("/products", showProducts)
	// r.GET("/products", uc.GetAllProducts)
	// r.GET("/product/:id", uc.GetProduct)
	// r.POST("/product", uc.CreateProduct)
	// // r.DELETE("/product/:id", uc.DeleteProduct)

	// r.GET("/carts", uc.GetAllCarts)
	// r.POST("/carts", uc.CreateCart)
	// //r.DELETE("/user/:id/cart", uc.DeleteCart)

	// r.GET("/user/:id/cart", uc.GetCartUser)
	// r.PUT("/user/:id/cart", uc.AddToCart)
	// r.DELETE("/user/:id/cart", uc.DeleteItemInCart)

	// r.POST("/user/:id/cart2", uc.AddToCart2)

	// // r.GET("/user/:id/payment", uc.GetPayment)
	// // r.POST("/user/:id/payment", uc.PostPayment)

	// // r.POST("/user/:id/order", uc.PlaceOrder)

	http.ListenAndServe("localhost:8012", nil)
	http.ListenAndServe("localhost:8014", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

func showLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("17271")

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		// t, _ := template.ParseFiles("login.gtpl")
		// t.Execute(w, nil)
		err := rnd.HTML(w, http.StatusOK, "login", nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["name"][0])
		fmt.Println("password:", r.Form["mob"])
	}

	// tmpl := template.Must(template.ParseFiles("views/login.html")) //dint work
	// tmpl.Execute(w, "data goes here")
}
func showProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("17271")

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		products := [2]string{"dhahdaydadbad", "rahdual"}
		fmt.Println(products)
		err := rnd.HTML(w, http.StatusOK, "display", products)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["name"][0])
		fmt.Println("password:", r.Form["mob"])
	}

}
