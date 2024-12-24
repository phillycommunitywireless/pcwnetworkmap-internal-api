package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/api/sheets/v4"
)

func process_internal(service *sheets.Service) []networkpoint {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	spreadsheetID := os.Getenv("SPREADSHEET_ID")
	readRange := os.Getenv("SPREADSHEET_READ_RANGE")

	resp, err := service.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	networkpoints := []networkpoint{}

	// Process the spreadsheet data
	for _, row := range resp.Values {
		// fmt.Println(row)

		networkpoint_properties := networkpoint_properties{
			Name:           row[1].(string),
			Id:             row[2].(string),
			Street_address: row[3].(string),
			Feature_type:   row[4].(string),
			Image:          row[5].(string),
			Latitude:       row[6].(string),
			Longitude:      row[7].(string),
		}

		// coordinates_inner := []int{row[9].(int), row[10].(int)}
		latitude, err := strconv.ParseFloat(row[9].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		longitude, err := strconv.ParseFloat(row[10].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		coordinates_inner := []float64{latitude, longitude}

		networkpoint_geometry := networkpoint_geometry{
			Feature_type: row[8].(string),
			Coordinates:  [][]float64{coordinates_inner},
		}

		networkpoint := networkpoint{
			Geojson_type: row[0].(string),
			Properties:   networkpoint_properties,
			Geometry:     networkpoint_geometry,
		}
		// fmt.Println(networkpoint)
		networkpoints = append(networkpoints, networkpoint)
	}

	return networkpoints
}

func prep_for_export(networkpoints []networkpoint) geojson_fmt {
	properties := crs_properties{
		Name: "urn:ogc:def:crs:OGC:1.3:CRS84",
	}

	crs := crs{
		Type:       "name",
		Properties: properties,
	}

	geojson := geojson_fmt{
		Type:     "FeatureCollection",
		Name:     "internal",
		Crs:      crs,
		Features: networkpoints,
	}

	return geojson
}
