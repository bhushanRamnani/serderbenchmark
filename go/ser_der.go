package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zip_code"`
}

type Movie struct {
	Name     string   `json:"name"`
	Year     int      `json:"year"`
	Director string   `json:"director"`
	Producer string   `json:"producer"`
	Actors   []string `json:"actors"`
	Genre    string   `json:"genre"`
}

type Person struct {
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Address        Address `json:"address"`
	PhoneNumber    int     `json:"phone_number"`
	FavoriteMovies []Movie `json:"favorite_movies"`
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteRune(rune(chars[rand.Intn(len(chars))]))
	}
	return builder.String()
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(90000) + 10000
}

func generateRandomMovie() Movie {
	return Movie{
		Name:     generateRandomString(10),
		Year:     rand.Intn(42) + 1980,
		Director: generateRandomString(10),
		Producer: generateRandomString(10),
		Actors:   []string{generateRandomString(10), generateRandomString(10), generateRandomString(10)},
		Genre:    generateRandomString(10),
	}
}

func generateRandomAddress() Address {
	return Address{
		Street:  generateRandomString(10),
		City:    generateRandomString(10),
		State:   generateRandomString(10),
		ZipCode: generateRandomNumber(),
	}
}

func generateRandomPerson() Person {
	return Person{
		FirstName:   generateRandomString(10),
		LastName:    generateRandomString(10),
		Address:     generateRandomAddress(),
		PhoneNumber: generateRandomNumber(),
		FavoriteMovies: []Movie{
			generateRandomMovie(),
			generateRandomMovie(),
			generateRandomMovie(),
		},
	}
}

func measureSerializationTime(persons []Person, iterations int) float64 {
	var totalSerializationTime float64

	for i := 0; i < iterations; i++ {
		startTime := time.Now()
		_, err := json.Marshal(persons)
		if err != nil {
			fmt.Println("Error during serialization:", err)
			os.Exit(1)
		}
		// fmt.Println(string(personsJSON))
		serializationTime := time.Since(startTime).Seconds() * 1000
		totalSerializationTime += serializationTime
	}

	return totalSerializationTime / float64(iterations)
}

func measureDeserializationTime(jsonString string, iterations int) float64 {
	var totalDeserializationTime float64

	for i := 0; i < iterations; i++ {
		startTime := time.Now()
		var persons []Person
		err := json.Unmarshal([]byte(jsonString), &persons)
		if err != nil {
			fmt.Println("Error during deserialization:", err)
			os.Exit(1)
		}
		deserializationTime := time.Since(startTime).Seconds() * 1000
		totalDeserializationTime += deserializationTime
	}

	return totalDeserializationTime / float64(iterations)
}

func main() {
	iterations := 10
	numPersons := 100

	if len(os.Args) > 1 {
		iterations, _ = strconv.Atoi(os.Args[1])
	}

	if len(os.Args) > 2 {
		numPersons, _ = strconv.Atoi(os.Args[2])
	}
	fmt.Printf("Number of Iterations: %d. Number of persons: %d\n", iterations, numPersons)

	persons := make([]Person, numPersons)
	for i := 0; i < numPersons; i++ {
		persons[i] = generateRandomPerson()
	}

	// Measure serialization time
	serializationTime := measureSerializationTime(persons, iterations)
	fmt.Printf("Avg Serialization time (%d iterations): %.6f milliseconds\n", iterations, serializationTime)

	// Serialize to JSON-formatted string
	personsJSON, err := json.Marshal(persons)
	if err != nil {
		fmt.Println("Error during serialization:", err)
		os.Exit(1)
	}

	// Print the size of the serialized JSON string
	serializedDataSize := len(personsJSON)
	fmt.Printf("Size of serialized data: %d bytes\n", serializedDataSize)

	// Measure deserialization time
	deserializationTime := measureDeserializationTime(string(personsJSON), iterations)
	fmt.Printf("Avg Deserialization time (%d iterations): %.6f milliseconds\n", iterations, deserializationTime)
}
