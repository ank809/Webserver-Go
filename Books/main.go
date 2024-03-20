package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Price  string `json:"price"`
}

var books = []Book{
	{
		ID:     "1",
		Name:   "To Kill a Mockingbird",
		Author: "Harper Lee",
		Price:  "$10",
	},
	{
		ID:     "2",
		Name:   "1984",
		Author: "George Orwell",
		Price:  "$15",
	},
	{
		ID:     "3",
		Name:   "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  "$20",
	},
}

// Function to get all the books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "	Method Not Allowed", http.StatusNotFound)
		log.Println("Error:	Method Not Allowed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Function t get book by its id
func getBookById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println("Error: Method not allowed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/book/")

	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			json.NewEncoder(w).Encode(map[string]string{"message": "book found"})
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}
	http.Error(w, "Book Not Found", http.StatusNotFound)
	log.Println("Error: Book Not Found")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println("Error:Method Not Allowed")
		return
	}
	var newbook Book
	if err := json.NewDecoder(r.Body).Decode(&newbook); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot decode data"})
		return
	}
	books = append(books, newbook)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book appended successfully"})
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println("Error: Method Not Allowed")
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/delete/")

	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
			json.NewEncoder(w).Encode(books)
			return

		}
	}
	http.Error(w, "Book Not Found", http.StatusNotFound)
	log.Println("Error:Book Not Found")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/update/")
	if r.Method != "PUT" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println("Error:Method Not Allowed")
		return
	}
	var updateBook Book
	if err := json.NewDecoder(r.Body).Decode(&updateBook); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "error in decoding data"})
		return
	}
	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			books[i] = updateBook
			json.NewEncoder(w).Encode(map[string]string{"success": "book details updated"})
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	http.Error(w, "Book Not Found", http.StatusNotFound)
	log.Println("Error:Book Not Found")
}

func main() {
	http.HandleFunc("/books", getAllBooks)
	http.HandleFunc("/book/", getBookById)
	http.HandleFunc("/addbook", addBook)
	http.HandleFunc("/delete/", deleteBook)
	http.HandleFunc("/update/", updateBook)
	fmt.Println("Server running at port 8081")
	http.ListenAndServe(":8081", nil)

}
