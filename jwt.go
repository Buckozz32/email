package jwt

import (
	"net/http"

	  "github.com/gorilla/mux"

	  "github.com/golang-jwt/jwt/v5"
	
)

type email struct {
	Name string
	Password string
	Id int
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
   w.Write([]byte("API is up and running"))
})


var EmailHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)){

	email, _ = json.Marshall(email)
w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(email))
}

 var mySigningKey = []byte("secret")

    var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, 
                                            r *http.Request)) {
     token := jwt.New(jwt.SigningMethodHS256)
											
    token.Claims[""] = true
    token.Claims["name"] = ""
    token.Claims[""] = time.Now().Add(time.Hour * 24).Unix()
											 tokenString, _ := token.SignedString(mySigningKey)
											 w.Write([]byte(tokenString))
    }
											
											




var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Not Implemented"))
})

func main()  {
	 r := mux.NewRouter()
	   r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", 
                                http.FileServer(http.Dir("./static/"))))

         http.ListenAndServe(":8080", r)
   
	
		 
    
    r.Handle("/status", NotImplemented).Methods("GET")
    r.Handle("/email", NotImplemented).Methods("GET")
    

 r.Handle("/get-token", GetTokenHandler).Methods("GET")


}





