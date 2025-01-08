package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/api/sheets/v4"
)

func get_sheet_values(service *sheets.Service, sheet_name string) *sheets.ValueRange {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning: .env file not found. Falling back to environment variables.")
	}

	spreadsheetID := os.Getenv("SPREADSHEET_ID")

	sheet, err := service.Spreadsheets.Values.Get(
		spreadsheetID, sheet_name,
	).Do()

	if err != nil {
		return nil
	} else {
		return sheet
	}

}

// level1
func process_level1(values *sheets.ValueRange) []level1_point {
	level1_points := []level1_point{}

	sliced := values.Values[5:]

	for _, row := range sliced {
		level1point_properties := level1_properties{
			Fid:           row[1].(string),
			Qc_id:         row[2].(string),
			Name:          row[3].(string),
			Description:   row[4].(string),
			Hs_id:         row[5].(string),
			Rt_id:         row[6].(string),
			Fid_2:         row[7].(string),
			Qc_id_2:       row[8].(string),
			Name_2:        row[9].(string),
			Description_2: row[10].(string),
			Hs_id_2:       row[11].(string),
		}

		latitude, err := strconv.ParseFloat(row[13].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		longitude, err := strconv.ParseFloat(row[14].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		z_axis, err := strconv.ParseFloat(row[15].(string), 64)
		if err != nil {
			panic(err)
		}

		latitude_2, err := strconv.ParseFloat(row[16].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		longitude_2, err := strconv.ParseFloat(row[17].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		z_axis_2, err := strconv.ParseFloat(row[18].(string), 64)
		if err != nil {
			panic(err)
		}

		coordinates_inner := []float64{
			latitude, longitude, z_axis,
		}

		coordinates_inner_2 := []float64{
			latitude_2, longitude_2, z_axis_2,
		}

		level1_geometry := level1_geometry{
			Type:        row[12].(string),
			Coordinates: [][]float64{coordinates_inner, coordinates_inner_2},
		}

		level1_point := level1_point{
			Type:       row[0].(string),
			Properties: level1point_properties,
			Geometry:   level1_geometry,
		}

		level1_points = append(level1_points, level1_point)
	}

	return level1_points
}

// level2+level3
func process_level2_level3(values *sheets.ValueRange) []level2_3_point {
	points := []level2_3_point{}
	// Take a slice of everything below A6
	sliced := values.Values[5:]

	for _, row := range sliced {
		properties := level2_3_properties{
			Fid:           row[1].(string),
			Qc_id:         row[2].(string),
			Name:          row[3].(string),
			Description:   row[4].(string),
			Hs_id:         row[5].(string),
			Rt_id:         row[6].(string),
			Fid_2:         row[7].(string),
			Qc_id_2:       row[8].(string),
			Name_2:        row[9].(string),
			Description_2: row[10].(string),
			Rt_id_2:       row[11].(string),
		}

		latitude, err := strconv.ParseFloat(row[13].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		longitude, err := strconv.ParseFloat(row[14].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		z_axis, err := strconv.ParseFloat(row[15].(string), 64)
		if err != nil {
			panic(err)
		}

		coordinates_inner := []float64{
			latitude, longitude, z_axis,
		}

		latitude_2, err := strconv.ParseFloat(row[16].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		longitude_2, err := strconv.ParseFloat(row[17].(string), 64) // Parse as float64
		if err != nil {
			panic(err)
		}

		z_axis_2, err := strconv.ParseFloat(row[18].(string), 64)
		if err != nil {
			panic(err)
		}

		coordinates_inner_2 := []float64{latitude_2, longitude_2, z_axis_2}

		geometry := level2_3_geometry{
			Type:        row[12].(string),
			Coordinates: [][]float64{coordinates_inner, coordinates_inner_2},
		}

		point := level2_3_point{
			Type:       row[0].(string),
			Properties: properties,
			Geometry:   geometry,
		}

		points = append(points, point)

	}

	return points

}

// Take sheet values as input from get_sheet_values, convert to array of structs
// This uses the "networkpoint" struct, which works for networkpoints + internal
// need a different struct for level1, level2, and level3
func process_networkpoints_and_internal(values *sheets.ValueRange) []networkpoint {
	networkpoints := []networkpoint{}

	// Take a slice of everything below A6
	sliced := values.Values[5:]

	for _, row := range sliced {
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

func prep_for_export_level1(points []level1_point) geojson_fmt_level1 {
	properties := crs_properties{
		Name: "urn:ogc:def:crs:OGC:1.3:CRS84",
	}

	crs := crs{
		Type:       "name",
		Properties: properties,
	}

	geojson := geojson_fmt_level1{
		Type:     "FeatureCollection",
		Name:     "level1",
		Crs:      crs,
		Features: points,
	}

	return geojson

}

func prep_for_export_level2_3(points []level2_3_point) geojson_fmt_level2_3 {
	properties := crs_properties{
		Name: "urn:ogc:def:crs:OGC:1.3:CRS84",
	}

	crs := crs{
		Type:       "name",
		Properties: properties,
	}

	geojson := geojson_fmt_level2_3{
		Type:     "FeatureCollection",
		Name:     "internal",
		Crs:      crs,
		Features: points,
	}

	return geojson

}
