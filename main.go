package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
<<<<<<< HEAD
	"io/ioutil"
	"log"
	"net/http"
	"os"
=======
	"log"
	"net/http"
>>>>>>> ad1cf3a06435a8320297114d74023fd7f0685ca0
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/users", getUsersHandler)
	http.HandleFunc("/users/update", updateUserHandler)
	http.HandleFunc("/users/delete", deleteUserHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	// Чтение содержимого файла index.html
	file, err := os.ReadFile("index.html")
=======
	tmpl, err := template.ParseFiles("index.html")
>>>>>>> ad1cf3a06435a8320297114d74023fd7f0685ca0
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

<<<<<<< HEAD
	// Преобразование содержимого в строку
	htmlContent := string(file)

	// Создание шаблона из строки
	tmpl, err := template.New("index").Parse(htmlContent)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Выполнение шаблона
=======
>>>>>>> ad1cf3a06435a8320297114d74023fd7f0685ca0
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
<<<<<<< HEAD
		file, err := os.ReadFile("register.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Преобразование содержимого в строку
		htmlContent := string(file)

		// Создание шаблона из строки
		tmpl, err := template.New("register").Parse(htmlContent)
=======
		tmpl, err := template.ParseFiles("register.html")
>>>>>>> ad1cf3a06435a8320297114d74023fd7f0685ca0
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := User{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	collection := client.Database("housing").Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

<<<<<<< HEAD
	file, err := os.ReadFile("success.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Преобразование содержимого в строку
	htmlContent := string(file)

	// Создание шаблона из строки
	tmpl, err := template.New("success").Parse(htmlContent)
=======
	tmpl, err := template.ParseFiles("success.html")
>>>>>>> ad1cf3a06435a8320297114d74023fd7f0685ca0
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

<<<<<<< HEAD
=======

>>>>>>> ad1cf3a06435a8320297114d74023fd7f0685ca0
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	collection := client.Database("housing").Collection("users")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var users []User
	if err := cursor.All(context.Background(), &users); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

<<<<<<< HEAD
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		htmlContent, err := ioutil.ReadFile("delete.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.New("delete").Parse(string(htmlContent))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}
		userID := r.FormValue("userID")

		userIDObj, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}

		collection := client.Database("housing").Collection("users")

		filter := bson.M{"_id": userIDObj}

		_, err = collection.DeleteOne(context.Background(), filter)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		htmlContent, err := ioutil.ReadFile("update.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.New("update").Parse(string(htmlContent))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}
		userID := r.FormValue("userID")
		newUsername := r.FormValue("newUsername")
		newEmail := r.FormValue("newEmail")

		userIDObj, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}

		collection := client.Database("housing").Collection("users")

		filter := bson.M{"_id": userIDObj}

		update := bson.M{"$set": bson.M{"username": newUsername, "email": newEmail}}

		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
=======
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	collection := client.Database("housing").Collection("users")
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email, "password": user.Password, "updated_at": time.Now()}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	collection := client.Database("housing").Collection("users")
	filter := bson.M{"_id": user.ID}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
>>>>>>> ad1cf3a06435a8320297114d74023fd7f0685ca0
}