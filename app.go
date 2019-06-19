package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"
	// "strconv"
	// "strings"

	. "trails-ms/config"
	. "trails-ms/dao"
	. "trails-ms/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var config = Config{}
var dao = TrailsDAO{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 8,
	WriteBufferSize: 1024 * 8,
}

const VIDEO_DIR = "."

const BUFSIZE = 1024 * 8

// ==============================================================================================

// INICIALIZACION Y CONFIGURACION DE LA BASE DE DATOS

func init() {
	config.Read()
	
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}


func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
// ==============================================================================================


// ==============================================================================================
// CREA UNA NUEVA RUTA
func CreateTrailEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	defer r.Body.Close()
	var trail Trail
	if err := json.NewDecoder(r.Body).Decode(&trail); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	trail.ID = bson.NewObjectId()
	
	if trail.Usertrail == 0 || trail.Nametrail == "" || trail.Origintrail == 0.0 ||trail.Destinytrail == 0.0{
		respondWithError(w, http.StatusInternalServerError, "Empty Values")
		return
	}
	
	if err := dao.InsertTrail(trail); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, trail)
}
// ==============================================================================================


// ==============================================================================================
// OBTIENE TODAS LAS RUTAS DE LA TABLA DE DATOS
func AllTrailsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	trails, err := dao.FindAllTrails()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, trails)
}
// ==============================================================================================


// ==============================================================================================
// OBTIENE TODAS LAS RUTAS DE UN USUARIO
func FindTrailsByUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	trails, err := dao.FindTrailsByUser(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, trails)
}
// ==============================================================================================


// ==============================================================================================
// OBTIENE UNA RUTA EN ESPECIFICO
func FindTrailByIdEnpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	trail, err := dao.FindTrailById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie Id")
		return
	}
	respondWithJson(w, http.StatusOK, trail)
}
// ==============================================================================================


// ==============================================================================================
// BORRA TODAS LAS RUTAS DE UN USUARIO
func DeleteTrailsEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	defer r.Body.Close()
	params := mux.Vars(r)
	info, err := dao.DeleteTrails(params["id"]);
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, info)
}
// ==============================================================================================


// ==============================================================================================
// BORRA UNA RUTA EN ESPECIFICO
func DeleteTrailByIdEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.DeleteTrailById(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
// ==============================================================================================


func main() {
	fmt.Println("Starting Server")
	r := mux.NewRouter()
	
	// CREA UNA NUEVA RUTA
	r.HandleFunc("/trails", CreateTrailEndpoint).Methods("POST")

	// OBTIENE TODAS LAS RUTAS DE LA TABLA DE DATOS
	r.HandleFunc("/trails", AllTrailsEndpoint).Methods("GET")

	// OBTIENE TODAS LAS RUTAS DE UN USUARIO
	r.HandleFunc("/trails/user/{id}", FindTrailsByUserEndpoint).Methods("GET")

	// OBTIENE UNA RUTA EN ESPECIFICO
	r.HandleFunc("/trails/{id}", FindTrailByIdEnpoint).Methods("GET")
	
	// BORRA TODAS LAS RUTAS DE UN USUARIO
	r.HandleFunc("/trails/user/{id}", DeleteTrailsEndPoint).Methods("DELETE")

	// BORRA UNA RUTA EN ESPECIFICO
	r.HandleFunc("/trails/{id}", DeleteTrailByIdEndPoint).Methods("DELETE")

	if err := http.ListenAndServe(":3002", r); err != nil {
		log.Fatal(err)
	}
}