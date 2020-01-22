package main

import (
	// Built-in Golang packages
	"context" // manage multiple requests
	"fmt"     // Println() function
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"reflect" // get an object type
	"time"

	// Official 'mongo-go-driver' packages
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type MongoFields struct {

	FirstName  string  `json:"string field"`
	LastName  string  `json:"string field"`
	Email  string     `json:"string field"`
	Password    string `json:"string field"`
	PhoneNumber string  `json:"string field"`

}

func main() {

	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://admin:Lmkt%40ptcl1234@192.168.0.224:27017/admin")
	fmt.Println("clientOptions TYPE:", reflect.TypeOf(clientOptions), "\n")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}

	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)

	// Access a MongoDB collection through a database
	col := client.Database("obaid").Collection("SignUp")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	// Declare a MongoDB struct instance for the document's fields and data


	oneDoc := MongoFields{

		FirstName:string(FName()) ,
		LastName:string(LName()),
		PhoneNumber:string(phone()),
		Password: string(hashAndSalt(getPwd())),
		Email:string(Email()),

	}
	fmt.Println("oneDoc TYPE:", reflect.TypeOf(oneDoc), "\n")

	// InsertOne() method Returns mongo.InsertOneResult
	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr !=

		nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		os.Exit(1) // safely exit script on error
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() API result:", result)

		// get the inserted ID string
		newID := result.InsertedID
		fmt.Println("InsertOne() newID:", newID)
		fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	}

	var result1 MongoFields
	err = col.FindOne(ctx, bson.D{{"email", string(Email1())}}).Decode(&result1)
	if err != nil {
		log.Fatal(err)
	}
	log1:= result1.Password

	fmt.Println("Encrypted password return from MongoDb: %+ v \n", log1)

		// Enter a password and generate a salted hash
		pwd := logPwd()
		hash := log1

		
		pwdMatch := comparePasswords(hash, pwd)
		fmt.Println("Passwords Match?", pwdMatch)

}
func logPwd() []byte {

	fmt.Println("Enter your  password to login")
	// We will use this to store the users input
	var pwd string
	// Read the users input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd)

}

	func FName() []byte{

	fmt.Println("Enter your first name")
	var f string
	_, _ = fmt.Scan(&f)
	return []byte(f)
}

func LName() []byte{

	fmt.Println("Enter your Last name")
	// We will use this to store the users input
	var l string
	// Read the users input
	_, _ = fmt.Scan(&l)
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(l)
}
func Email1() []byte{

	fmt.Println("Enter your Email ID to login")
	// We will use this to store the users input
	var e string
	// Read the users input
	_, _ = fmt.Scan(&e)
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(e)
}
func Email() []byte{

	fmt.Println("Enter your Email ID")
	// We will use this to store the users input
	var e string
	// Read the users input
	_, _ = fmt.Scan(&e)
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(e)
}

func phone() []byte{

	fmt.Println("Enter your phone number")
	// We will use this to store the users input
	var p string
	// Read the users input
	_, _ = fmt.Scan(&p)
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(p)
}
func getPwd() []byte {

	fmt.Println("Enter your  password")
	// We will use this to store the users input
	var pwd string
	// Read the users input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

