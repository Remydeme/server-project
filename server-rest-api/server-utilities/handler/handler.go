package handler

import (
	"os"
	"server-project/server-rest-api/server-utilities/service"

	"log"

	"github.com/kataras/iris"
)

/* Set:
Response : the Token
Infos : This function is used for the JWT logic
it init a session on the user login
*/
func Set(ctx iris.Context) {
	name, err := os.Hostname()
	claims, err := service.GetClaims(&ctx)
	if claims == nil {
		log.Println(err)
	}
	response := make(map[string]string)
	response["host"] = name
	ctx.JSON(response)
}

//Method : Get
//Response : 200 status
//Infos : This function is called to end a session on user logout
func Delete(ctx iris.Context) {

}

/*Method : Get
Repsonse : 200 status code
Infos : THis function is called to logout the user if the user is login
*/
func Logout(ctx iris.Context) {

}

/*Login : Post
Repsonse : 200 status code
Header : {
	Content-Type: apllication/json,
}
body: {
		"username" : "yourusername",
		"password" : "yourpassword",
     }
Infos : This function start sign in the user and init a session
*/
func Login(ctx iris.Context) {
	ctx.ContentType("application/json")
	response, err := service.Login(&ctx)
	if response == nil {
		log.Println(err)
		return
	}
	ctx.JSON(response)
}

// Method : Get
// Reponse : 404
// Infos : Bad paramters format
func NotFound(ctx iris.Context) {
	// when 404 then render the template $views_dir/errors/404.html
	ctx.JSON(iris.Map{"response": "Not found"})
}

// Method : Get
// Reponse : 500
// infos  : Error internal server error
func InternalServerError(ctx iris.Context) {
	ctx.JSON(iris.Map{"response": "Internal server error"})
}

// Method : Get
// Response : 200 status code
// return value when user acces to root /
func Index(ctx iris.Context) {
	ctx.JSON(iris.Map{"response": "Index request"})
}
