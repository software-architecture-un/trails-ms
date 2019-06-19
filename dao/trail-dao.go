package dao

import (
	"fmt"
	// "io/ioutil"
	"log"
	// "mime/multipart"
	"strconv"

	. "trails-ms/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ==============================================================================================
// INICIALIZACION Y CONFIGURACION DE LA BASE DE DATOS
type TrailsDAO struct {
	Server   string
	Database string
}

var DB *mgo.Database

// Establish a connection to database
func (m *TrailsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	DB = session.DB(m.Database)
}
// ==============================================================================================


// ==============================================================================================
// CREA UNA NUEVA RUTA
func (m *TrailsDAO) InsertTrail(trail Trail) error {
	err := DB.C("trails").Insert(&trail)
	return err
}
// ==============================================================================================


// ==============================================================================================
// OBTIENE TODAS LAS RUTAS DE LA TABLA DE DATOS
func (m *TrailsDAO) FindAllTrails() ([]Trail, error) {
	var trails []Trail
	err := DB.C("trails").Find(bson.M{}).All(&trails)
	return trails, err
}
// ==============================================================================================


// ==============================================================================================
// OBTIENE TODAS LAS RUTAS DE UN USUARIO
func (m *TrailsDAO) FindTrailsByUser(id string) ([]Trail, error) {

	usertrail, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(usertrail)
	}
	var trails []Trail
	err = DB.C("trails").Find(bson.M{"usertrail": usertrail}).All(&trails)
	return trails, err
}
// ==============================================================================================


// ==============================================================================================
// OBTIENE UNA RUTA EN ESPECIFICO
func (m *TrailsDAO) FindTrailById(id string) (Trail, error) {
	var trail Trail
	err := DB.C("trails").FindId(bson.ObjectIdHex(id)).One(&trail)
	return trail, err
}
// ==============================================================================================


// ==============================================================================================
// BORRA TODAS LAS RUTAS DE UN USUARIO
func (m *TrailsDAO) DeleteTrails(id string) (*mgo.ChangeInfo, error) {
	usertrail, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(usertrail)
	}
	info, err := DB.C("trails").RemoveAll(bson.M{"usertrail": usertrail})
	if err != nil {
		return nil, err
	}
	return info, nil
}
// ==============================================================================================


// ==============================================================================================
// BORRA UNA RUTA EN ESPECIFICO
func (m *TrailsDAO) DeleteTrailById(id string) error {
	var trail Trail
	err := DB.C("trails").FindId(bson.ObjectIdHex(id)).One(&trail)
	err = DB.C("trails").Remove(&trail)
	return err
}
// ==============================================================================================


// func (m *VideosDAO) DeleteCategory(category Category) error {
// 	err := DB.C("categories").Remove(&category)
// 	return err
// }

// // Update an existing video
// func (m *VideosDAO) UpdateCategory(category Category) error {
// 	err := DB.C("categories").UpdateId(category.ID, &category)
// 	return err
// }