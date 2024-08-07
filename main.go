// A simple API for desert tortoise facts.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/mux"
)

type TortoiseFact struct {
	ID		 string `json:"id"`
	Fact	 string `json:"fact"`
}

var facts []TortoiseFact

func loadTortoiseFacts(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	err = json.Unmarshal(data, &facts)
	if err != nil {
		return fmt.Errorf("Error unmarshaling JSON: %v", err)
	}

	return nil
}

func getFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	randomFact := facts[rand.Intn(len(facts))]
	json.NewEncoder(w).Encode(randomFact)
}

func getAllFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(facts)
}

func addFact(w http.ResponseWriter, r *http.Request) {
	var newFact TortoiseFact
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &newFact)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if newFact.Fact == "" {
		http.Error(w, "Fact cannot be empty", http.StatusBadRequest)
		return
	}

	newID := generateUniqueID()
	newFact.ID = newID

	facts = append(facts, newFact)

	fmt.Printf("New fact added: %v\n", newFact)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newFact)
}

func generateUniqueID() string {
	maxID := 0
	for _, fact := range facts {
		id, err := strconv.Atoi(fact.ID[2:]) // Assuming IDs are in the format "DTxxx"
		if err == nil && id > maxID {
			maxID = id
		}
	}
	return fmt.Sprintf("DT%03d", maxID+1)
}

func main() {
	err := loadTortoiseFacts("data/facts.json")
	if err != nil {
		fmt.Printf("Error loading tortoise facts: %v\n", err)
	}

	// Start the Mux router.
	router := mux.NewRouter()
	router.HandleFunc("/fact", getFact).Methods("GET")
	router.HandleFunc("/facts", getAllFacts).Methods("GET")
	router.HandleFunc("/add", addFact).Methods("POST")

	fmt.Println("API is running on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
