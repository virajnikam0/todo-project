package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/virajnikam0/todo-project/web-todo-project/internal/response"
	myTypes "github.com/virajnikam0/todo-project/web-todo-project/internal/types"
)

// Register
func RegisterUser(w http.ResponseWriter,r *http.Request){
	var backendUser myTypes.User
	// check request 
	if r.Method == http.MethodPost{
		// get the data from body 
		err := json.NewDecoder(r.Body).Decode(&backendUser)
		if err == nil{
				myTypes.Users = append(myTypes.Users,backendUser)
				res := response.CorrectResponse{
					StatusCode: http.StatusOK,
					Data: map[string]string{
						"msg":"User added successfully",
						"code" : string(http.StatusOK),
					},
				}
				data,_ :=json.Marshal(res)
				w.Write(data)
				
			fmt.Println("inside the if condition ",myTypes.Users)
		}else {
			// data,_ :=json.Marshal(res)
			// errorRes := response.ErrorResponse{
			// 	StatusCode: http.StatusBadRequest,
			// 	Data: map[string]string {
			// 		"msg": "body data cannot gets decoded from frontend",
			// 		"data":string(backendUser.UserEmail + " " +  backendUser.UserName + " " +  backendUser.UserPassw ),
			// 	},
			// }
			// fmt.Println("inside the else condition ",myTypes.Users)

			// res,_ := json.Marshal(errorRes)
			w.Write([]byte(err.Error()))
			
		}
	}
}

// login in
func Login(w http.ResponseWriter,r *http.Request){
	// check and verify it just
	var backendUser myTypes.User



    json.NewDecoder(r.Body).Decode(&backendUser)

	for _,user := range myTypes.Users{
		if user.UserEmail == backendUser.UserEmail && user.UserPassw == backendUser.UserPassw {
			
			// make the JWT token creation process start
			claims := jwt.MapClaims{
				"username" : backendUser.UserName,
			}
			// create the token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
			// create token string
			secretKey := []byte("123")
			tokenString,_ := token.SignedString(secretKey)
			w.Header().Set("Content-Type","application/json")
			fmt.Println("my token string : ",tokenString)
			json.NewEncoder(w).Encode(&tokenString)
			w.Write([]byte(tokenString))
			}else{
			w.Write([]byte("no user cannot enter"))
		}
	}
}