package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// Var User
type User struct {
	Username string
	Password string
}

var db *sql.DB

// Hbungin Postgres
func initDB() {
	var err error
	// Connection string PostgreSQL
	connStr := "user=postgres password=admin dbname=Login sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Check Ping ke Postgres
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	log.Println("Successfully connected to the database!")
}

// CORS buat svelte akses
func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// fungsi buat nyamain user dan pass
func authenticate(w http.ResponseWriter, r *http.Request) {
	log.Println("Received login request")

	// Parse data yang dari login page (DECODE Content-Type)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		log.Println("Error parsing form:", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println("Username:", username)

	var dbPassword string
	// akses database dan ambil password dari spesifik username
	query := "SELECT password FROM users WHERE username=$1" // $1 prevensi kecil Injection
	err = db.QueryRow(query, username).Scan(&dbPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			log.Println("Invalid username or password")
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
			log.Println("Database error:", err)
		}
		return
	}

	// Banding input Password dan Password yang di dB
	if password == dbPassword {
		fmt.Fprintln(w, "Login successful")
		log.Println("Login successful")
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		log.Println("Invalid username or password")
	}
}

// Fungsi Registe
func register(w http.ResponseWriter, r *http.Request) {
	log.Println("Received registration request")

	// Parse data dari fom Registrasi
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		log.Println("Error parsing form:", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println("Username:", username)

	// Check Username sudah ada yang punya atau belum jika sudah ada akan keluar error dan jika gk ada akan di input ke dB
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = db.Exec(query, username, password)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			http.Error(w, "Username already exists", http.StatusConflict)
			log.Println("Username already exists")
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
			log.Println("Database error:", err)
		}
		return
	}

	fmt.Fprintln(w, "Registration successful")
	log.Println("Registration successful")
}

func main() {
	// Paggil Fungsi Koneksi ke DB
	initDB()
	defer db.Close()

	// Handle Login dan Register
	http.HandleFunc("/login", authenticate)
	http.HandleFunc("/register", register)

	// CORS akses
	handler := enableCORS(http.DefaultServeMux)

	// Start Server
	port := ":8081"
	log.Printf("Starting server on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
