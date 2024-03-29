package peda

import (
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func GetAllBangunanLineString(mongoconn *mongo.Database, collection string) []GeoJson {
	lokasi := atdb.GetAllDoc[[]GeoJson](mongoconn, collection)
	return lokasi
}

func GetAllProduct(mongoconn *mongo.Database, collection string) []Product {
	product := atdb.GetAllDoc[[]Product](mongoconn, collection)
	return product
}

func GetNameAndPassowrd(mongoconn *mongo.Database, collection string) []User {
	user := atdb.GetAllDoc[[]User](mongoconn, collection)
	return user
}

func GetAllUser(mongoconn *mongo.Database, collection string) []User {
	user := atdb.GetAllDoc[[]User](mongoconn, collection)
	return user
}
func CreateNewUserRole(mongoconn *mongo.Database, collection string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Insert the user data into the database
	return atdb.InsertOneDoc(mongoconn, collection, userdata)
}
func CreateUserAndAddedToeken(PASETOPRIVATEKEYENV string, mongoconn *mongo.Database, collection string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Insert the user data into the database
	atdb.InsertOneDoc(mongoconn, collection, userdata)

	// Create a token for the user
	tokenstring, err := watoken.Encode(userdata.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		return err
	}
	userdata.Token = tokenstring

	// Update the user data in the database
	return atdb.ReplaceOneDoc(mongoconn, collection, bson.M{"username": userdata.Username}, userdata)
}

func DeleteUser(mongoconn *mongo.Database, collection string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}
func ReplaceOneDoc(mongoconn *mongo.Database, collection string, filter bson.M, userdata User) interface{} {
	return atdb.ReplaceOneDoc(mongoconn, collection, filter, userdata)
}
func FindUser(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}

func FindUserUser(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{
		"username": userdata.Username,
	}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}

func IsPasswordValid(mongoconn *mongo.Database, collection string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mongoconn, collection, filter)
	return CheckPasswordHash(userdata.Password, res.Password)
}

// product

func CreateNewProduct(mongoconn *mongo.Database, collection string, productdata Product) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, productdata)
}

func InsertUserdata(mongoenv *mongo.Database, collname, username, role, password string) (InsertedID interface{}) {
	req := new(User)
	req.Username = username
	req.Password = password
	req.Role = role
	return atdb.InsertOneDoc(mongoenv, collname, req)
}

func PostLineString(mongoconn *mongo.Database, collection string, commentdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, commentdata)
}

func PostLinestring(mongoconn *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, linestringdata)
}

func GetByCoordinate(mongoconn *mongo.Database, collection string, linestringdata GeoJsonLineString) GeoJsonLineString {
	filter := bson.M{"geometry.coordinates": linestringdata.Geometry.Coordinates}
	return atdb.GetOneDoc[GeoJsonLineString](mongoconn, collection, filter)
}

func DeleteLinestring(mongoconn *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	filter := bson.M{"geometry.coordinates": linestringdata.Geometry.Coordinates}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func UpdatedLinestring(mongoconn *mongo.Database, collection string, filter bson.M, linestringdata GeoJsonLineString) interface{} {
	filter = bson.M{"geometry.coordinates": linestringdata.Geometry.Coordinates}
	return atdb.ReplaceOneDoc(mongoconn, collection, filter, linestringdata)
}

func PostPolygone(mongoconn *mongo.Database, collection string, polygonedata GeoJsonPolygon) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, polygonedata)
}

func PostPoint(mongoconn *mongo.Database, collection string, pointdata GeometryPoint) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, pointdata)
}
