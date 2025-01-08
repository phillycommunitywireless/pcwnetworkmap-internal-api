package main

type geojson_fmt struct {
	Type     string         `json:"type"`
	Name     string         `json:"name"`
	Crs      crs            `json:"crs"`
	Features []networkpoint `json:"features"`
}

type geojson_fmt_level1 struct {
	Type     string         `json:"type"`
	Name     string         `json:"name"`
	Crs      crs            `json:"crs"`
	Features []level1_point `json:"features"`
}

type geojson_fmt_level2_3 struct {
	Type     string           `json:"type"`
	Name     string           `json:"name"`
	Crs      crs              `json:"crs"`
	Features []level2_3_point `json:"features"`
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

type level1_point struct {
	Type       string            `json:"type"`
	Properties level1_properties `json:"properties"`
	Geometry   level1_geometry   `json:"geometry"`
}

type level1_properties struct {
	Fid           string `json:"fid"`
	Qc_id         string `json:"qc_id`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Hs_id         string `json:"hs_id"`
	Rt_id         string `json:"rt_id"`
	Fid_2         string `json:"fid_2"`
	Qc_id_2       string `json:"qc_id_2"`
	Name_2        string `json:"name_2"`
	Description_2 string `json:"description_2"`
	Hs_id_2       string `json:"hs_id_2"`
}

type level1_geometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type level2_3_point struct {
	Type       string              `json:"type"`
	Properties level2_3_properties `json:"properties"`
	Geometry   level2_3_geometry   `json:"geometry"`
}

type level2_3_properties struct {
	Fid           string `json:"fid"`
	Qc_id         string `json:"qc_id`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Hs_id         string `json:"hs_id"`
	Rt_id         string `json:"rt_id"`
	Fid_2         string `json:"fid_2"`
	Qc_id_2       string `json:"qc_id_2"`
	Name_2        string `json:"name_2"`
	Description_2 string `json:"description_2"`
	Rt_id_2       string `json:"rt_id_2"`
}

type level2_3_geometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}
