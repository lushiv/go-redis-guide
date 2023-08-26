package main

import (
	"encoding/json"
	"fmt"

	redisutil "go-redis-guide/utils"
)

var redisHost = "redis-16171.c239.us-east-1-2.ec2.cloud.redislabs.com:16171"
var redisPassword = "KHKVvAFMGMS4S9tISZA0GAEm5WhTRASa"
var redisDB = 0

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Initialize Redis connection
	err := redisutil.InitializeConnection(redisHost, redisPassword, redisDB)
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}

	// Example JSON data
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	// Convert the person struct to JSON
	personJSON, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Set JSON data in Redis
	err = redisutil.SetKey("person:1", string(personJSON))
	if err != nil {
		fmt.Println("Error setting data in Redis:", err)
		return
	}

	// Get JSON data from Redis
	retrievedJSON, err := redisutil.GetKey("person:1")
	if err != nil {
		fmt.Println("Error getting data from Redis:", err)
		return
	}

	// Convert the JSON data back to struct
	var retrievedPerson Person
	err = json.Unmarshal([]byte(retrievedJSON), &retrievedPerson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print retrieved person data
	fmt.Println("Retrieved Person:")
	fmt.Println("Name:", retrievedPerson.Name)
	fmt.Println("Age:", retrievedPerson.Age)
	fmt.Println("Email:", retrievedPerson.Email)

	// Delete the key from Redis
	err = redisutil.DeleteKey("person:1")
	if err != nil {
		fmt.Println("Error deleting key from Redis:", err)
		return
	}

	fmt.Println("Deleted")

	// Clear the cache (flush all keys)
	err = redisutil.ClearCache()
	if err != nil {
		fmt.Println("Error clearing cache:", err)
		return
	}

	fmt.Println("Cache Cleared")
}
