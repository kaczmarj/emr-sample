package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Printf("received request to get index")
	_, err := io.WriteString(w, "Home page.")
	if err != nil {
		http.Error(w, "error writing content", http.StatusInternalServerError)
	}
}

func AddPatient(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("received request to add patient")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusInternalServerError)
		return
	}
	pt := NewPatient(r.FormValue("name"), r.FormValue("dob"))
	AddPatientToList(pt)
}

func GetPatient(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	log.Printf("received request to get patient")
	ids := ps.ByName("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		http.Error(w, "error converting ID to integer", http.StatusBadRequest)
		return
	}
	pt, err := Patients.Get(id)
	if err != nil {
		http.Error(w, "patient not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(pt)
	if err != nil {
		http.Error(w, "error encoding resource to json", http.StatusInternalServerError)
		return
	}
}

func PatchPatient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("received request to patch patient")
	ids := ps.ByName("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		http.Error(w, "error converting ID to integer", http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusInternalServerError)
		return
	}
	err = EditPatient(id, r.FormValue("name"), r.FormValue("dob"))
	if err != nil {
		http.Error(w, "patient not found", http.StatusNotFound)
		return
	}
}

func GetPatients(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Printf("received request to get patients")
	err := json.NewEncoder(w).Encode(Patients)
	if err != nil {
		http.Error(w, "error encoding resource to json", http.StatusInternalServerError)
		return
	}
}
