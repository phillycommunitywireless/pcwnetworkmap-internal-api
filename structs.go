package main

type geojson_fmt struct {
	Type     string         `json:"type"`
	Name     string         `json:"name"`
	Crs      crs            `json:"crs"`
	Features []networkpoint `json:"features"`
}

type crs struct {
	Type       string         `json:"name"`
	Properties crs_properties `json:"properties"`
}

type crs_properties struct {
	Name string `json:"name"`
}

type networkpoint_properties struct {
	Name           string `json:"name"`
	Id             string `json:"id"`
	Street_address string `json:"street_address"`
	Image          string `json:"Image"`
	Feature_type   string `json:"type"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
}

type networkpoint_geometry struct {
	Feature_type string      `json:"type"`
	Coordinates  [][]float64 `json:"coordinates"`
}

type networkpoint struct {
	Geojson_type string                  `json:"type"`
	Properties   networkpoint_properties `json:"properties"`
	Geometry     networkpoint_geometry   `json:"geometry"`
}
