package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var connectionStr = "user=postgres password=asdyfe2rd dbname=webclinic host=localhost port=5432 sslmode=disable"

func main() {
	r := mux.NewRouter()
	// CORS middleware configuration
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Adjust this to your frontend's origin
	)
	// Use CORS middleware
	r.Use(cors)
	// Define your routes
	r.HandleFunc("/signin", signInHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/signup", signUpHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/create-slot", createSlotHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/view-avail-slot", getDoctorSchedulesHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/select-dr", selectDrHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/choose-slot", chooseSlotHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/view-all-doctors", getAllDoctorNamesHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/cancel-appointment/{scheduleid}", cancelAppointmentHandler).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/get-all-reservations/{patientID}", getAllReservationsHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/update-slot", updateSlotHandler).Methods("PUT", "OPTIONS")
	// Start the server
	port := 1234
	fmt.Printf("Server is running on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
