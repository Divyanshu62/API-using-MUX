/*
Routes application using MUX package
*/


package main

import (
  "fmt"
  "encoding/json"
//  "log"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux" //  need to install this package go -> get -u github.com/gorilla/mux
)

// create a book model

type Team struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Designation string `json:"designation"`
  Age string `json:"age"`
}

// slice of getBooks
var teams []Team

// get all books
func getTeams(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-type", "appplication/json")
    json.NewEncoder(w).Encode(teams)
}

// get one book
func getTeam(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-type", "appplication/json")
    params:= mux.Vars(r)
    fmt.Println(params)
    for _, item := range teams{
      if item.ID == params["id"]{
        json.NewEncoder(w).Encode(item)
        return
      }
    }
    json.NewEncoder(w).Encode(&Team{})
}

// create a book
func createTeam(w http.ResponseWriter, r *http.Request){
      w.Header().Set("Content-type", "appplication/json")
      var team Team
      _ = json.NewDecoder(r.Body).Decode(&team)
      team.ID = strconv.Itoa(rand.Intn(1000000))   // Itoa func type cast to string from int form strconv package
      team.Name = "My " + strconv.Itoa(rand.Intn(10)) + " Name"
      team.Designation = "My " + strconv.Itoa(rand.Intn(10)) + " Designation"
      team.Age = strconv.Itoa(rand.Intn(100))
      teams = append(teams, team)
      json.NewEncoder(w).Encode(teams)

}




// delete a book
func deleteTeam(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-type", "appplication/json")
   params:= mux.Vars(r)

   for index, item := range teams {
     if item.ID == params["id"]{
       teams = append(teams[:index],  teams[index+1:]...)
       break
     }
   }

   json.NewEncoder(w).Encode(teams)

}

func main(){
  //init router
  r := mux.NewRouter()



  // router handlers
  r.HandleFunc("/api/books", getTeams).Methods("GET")
  r.HandleFunc("/api/book/{id}", getTeam).Methods("GET")
  r.HandleFunc("/api/books", createTeam).Methods("POST")
  r.HandleFunc("/api/book/{id}", deleteTeam).Methods("DELETE")
  http.ListenAndServe(":8092", r)

}
