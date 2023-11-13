package peda

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllBangunanLineString(mconn, collectionname)
	return GCFReturnStruct(datagedung)
}

func Adduser(mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		hash, hashErr := HashPassword(datauser.Password)
		if hashErr != nil {
			response.Message = "Gagal Hash Password" + err.Error()
		}
		InsertUserdata(mconn, collname, datauser.Username, datauser.Role, hash)
		response.Message = "Berhasil Input data"
	}
	return GCFReturnStruct(response)
}

func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah"
		}
	}

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func GCFCreatePostLineStringg(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}

	// Mengambil nilai header PASETO dari permintaan HTTP
	pasetoValue := r.Header.Get("PASETOPRIVATEKEYENV")

	// Disini Anda dapat menggunakan nilai pasetoValue sesuai kebutuhan Anda
	// Misalnya, menggunakannya untuk otentikasi atau enkripsi.
	// Contoh sederhana menambahkan nilainya ke dalam pesan respons:
	response := GCFReturnStruct(geojsonline)
	response += " PASETO value: " + pasetoValue

	PostLinestring(mconn, collection, geojsonline)
	return response
}

func GCFCreatePostLineString(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}
	PostLinestring(mconn, collection, geojsonline)
	return GCFReturnStruct(geojsonline)
}

func GCFDeleteLineString(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var dataline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&dataline)
	if err != nil {
		return err.Error()
	}

	if err := DeleteLinestring(mconn, collectionname, dataline); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Delete LineString", dataline))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Delete LineString", dataline))
	}
}

func GCFUpdateLinestring(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var dataline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&dataline)
	if err != nil {
		return err.Error()
	}

	if err := UpdatedLinestring(mconn, collectionname, bson.M{"properties.coordinates": dataline.Geometry.Coordinates}, dataline); err != nil {
		return GCFReturnStruct(CreateResponse(true, "Success Update LineString", dataline))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Update LineString", dataline))
	}
}

func GCFCreateLineStringgg(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Parsing Request Body
	var dataline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&dataline)
	if err != nil {
		return err.Error()
	}

	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		// Handling Authorization
		err := PostLinestring(mconn, collectionname, dataline)
		if err != nil {
			// Success
			return GCFReturnStruct(CreateResponse(true, "Success: LineString created", dataline))
		} else {
			return GCFReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// This part is unreachable, so you might want to remove it
	// return GCFReturnStruct(CreateResponse(false, "Success to create LineString", nil))
}

func GCFCreatePolygone(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Parsing Request Body
	var datapolygone GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&datapolygone)
	if err != nil {
		return err.Error()
	}

	// Handling Authorization
	if err := PostPolygone(mconn, collectionname, datapolygone); err != nil {
		// Success
		return GCFReturnStruct(CreateResponse(true, "Success Create Polygone", datapolygone))
	} else {
		// Failure
		return GCFReturnStruct(CreateResponse(false, "Failed Create Polygone", datapolygone))
	}
}

func GCFPoint(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datapoint GeometryPoint

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&datapoint); err != nil {
		log.Printf("Error decoding request body: %v", err)
		return GCFReturnStruct(CreateResponse(false, "Bad Request: Invalid JSON", nil))
	}

	// Check for the "Secret" header
	secretHeader := r.Header.Get("Secret")
	expectedSecret := os.Getenv("SECRET")

	if secretHeader != expectedSecret {
		log.Printf("Unauthorized: Secret header does not match. Expected: %s, Actual: %s", expectedSecret, secretHeader)
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// Attempt to post the data point to MongoDB
	if err := PostPoint(mconn, collectionname, datapoint); err != nil {
		log.Printf("Error posting data point to MongoDB: %v", err)
		return GCFReturnStruct(CreateResponse(false, "Failed to create point", nil))
	}

	log.Println("Success: Point created")
	return GCFReturnStruct(CreateResponse(true, "Success: Point created", datapoint))
}

func GCFlineStingCreate(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}
	PostLinestring(mconn, collection, geojsonline)
	return GCFReturnStruct(geojsonline)
}

func GCFlineStingCreatea(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Parsing Request Body
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return GCFReturnStruct(CreateResponse(false, "Bad Request: Invalid JSON", nil))
	}

	// Checking Secret Header
	secretHeader := r.Header.Get("Secret")
	expectedSecret := os.Getenv("SECRET")

	if secretHeader != expectedSecret {
		log.Printf("Unauthorized: Secret header does not match. Expected: %s, Actual: %s", expectedSecret, secretHeader)
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// Handling Authorization
	PostLinestring(mconn, collection, geojsonline)

	return GCFReturnStruct(CreateResponse(true, "Success: LineString created", geojsonline))
}

func GCFCreatePolygonee(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Parsing Request Body
	var datapolygone GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&datapolygone)
	if err != nil {
		return GCFReturnStruct(CreateResponse(false, "Bad Request: Invalid JSON", nil))
	}

	// Checking Secret Header
	secretHeader := r.Header.Get("Secret")
	expectedSecret := os.Getenv("SECRET")

	if secretHeader != expectedSecret {
		log.Printf("Unauthorized: Secret header does not match. Expected: %s, Actual: %s", expectedSecret, secretHeader)
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// Handling Authorization
	if err := PostPolygone(mconn, collectionname, datapolygone); err != nil {
		log.Printf("Error creating polygon: %v", err)
		return GCFReturnStruct(CreateResponse(false, "Failed Create Polygone", nil))
	}

	log.Println("Success: Polygon created")
	return GCFReturnStruct(CreateResponse(true, "Success Create Polygone", datapolygone))
}
