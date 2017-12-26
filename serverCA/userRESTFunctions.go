package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Token    string        `json:"token"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	//ipFilter(w, r)

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Print("user signup: ")
	fmt.Println(user)

	//save the new project to mongodb
	rUser := User{}
	err = userCollection.Find(bson.M{"email": user.Email}).One(&rUser)
	if err != nil {
		//user not exists
		err = userCollection.Insert(user) //TODO find a way to get the object result when inserting in one line, without need of the two mgo petitions
		err = userCollection.Find(bson.M{"email": user.Email}).One(&user)
	} else {
		//user exists
		fmt.Fprintln(w, "User already registered")
		return
	}

	jResp, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}

func Login(w http.ResponseWriter, r *http.Request) {
	//ipFilter(w, r)

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Print("user login: ")
	fmt.Println(user)
	token, err := newToken()
	check(err)
	user.Token = token

	//save the new project to mongodb
	rUser := User{}
	err = userCollection.Find(bson.M{"email": user.Email}).One(&rUser)
	if err != nil {
	} else {
		//user exists, update with the token
		err = userCollection.Update(bson.M{"_id": rUser.Id}, user)
		check(err)
	}
	//generate the token
	//add the token to the user
	//save the user with the new token

	jResp, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}

type Sign struct {
	M string `json:"m"`
	C string `json:"c"`
}

func BlindSign(w http.ResponseWriter, r *http.Request) {
	//ipFilter(w, r)

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	jResp, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}
func VerifySign(w http.ResponseWriter, r *http.Request) {
	//ipFilter(w, r)

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	jResp, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}
