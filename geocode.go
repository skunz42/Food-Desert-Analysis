package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "time"
    "io/ioutil"
    "googlemaps.github.io/maps"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Coords struct {
    lat, lng float64
}

func geocode(city string) []maps.GeocodingResult {
    api_key := os.Getenv("GMAPS_GO_KEY")
    if api_key == "" {
        log.Fatalf("error getting env var")
    }

    c, err := maps.NewClient(maps.WithAPIKey(api_key))
    if err != nil {
        log.Fatalf("error when reading in api key: ", err)
    }

    r := &maps.GeocodingRequest {
        Address: city,
    }

    resp, err := c.Geocode(context.Background(), r)

    return resp
}

func parse_response(raw_data []maps.GeocodingResult) {
    km := 0.015
    coords := make([]Coords, 0)
    northeast := raw_data[0].Geometry.Bounds.NorthEast
    southwest := raw_data[0].Geometry.Bounds.SouthWest

    nelat := float64(northeast.Lat)
    nelng := float64(northeast.Lng)
    swlat := float64(southwest.Lat)
    swlng := float64(southwest.Lng)
    templat := swlat
    templng := swlng

    for templat <= nelat {
        for templng <= nelng {
            coords = append(coords, Coords{templat, templng})
            templng += km
        }
        templng = swlng
        templat += km
    }

    fmt.Println(len(coords))
    //write to db
}

func main() {
    /*raw_resp := geocode("Binghamton, NY")
    parse_response(raw_resp)*/

    temp_uri, err := ioutil.ReadFile("../Geo-Credentials/atlas2.txt")
    uri := string(temp_uri)
    if err != nil {
        log.Fatalf("i'm going to off myself")
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil { log.Fatal(err) }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)


    err = client.Ping(ctx, readpref.Primary())
	if err != nil {
        fmt.Println("couldnt ping")
		log.Fatal(err)
	}


    databases, err := client.ListDatabaseNames(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(databases)
}
