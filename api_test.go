package main

import (
	"fmt"
	"testing"
)

// -------------------------- Process sheet values tests -------------------------- //

func Test_process_level3(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()
	result_level3 := get_sheet_values(sheetsService, "level3")
	if result_level3 == nil {
		t.Fatal("Test_process_level3 failed")
	}

	level3 := process_level2_level3(result_level3)
	fmt.Println(level3)
}

func Test_process_level2(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()
	result_level2 := get_sheet_values(sheetsService, "level2")
	if result_level2 == nil {
		t.Fatal("Test_process_level2 failed")
	}

	level2 := process_level2_level3(result_level2)
	fmt.Println(level2)
}

func Test_process_level1(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()
	result_level1 := get_sheet_values(sheetsService, "level1")
	if result_level1 == nil {
		t.Fatal("Test_process_level1 failed")
	}

	level1 := process_level1(result_level1)
	fmt.Println(level1)
}

func Test_process_networkpoints(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()
	result_networkpoints := get_sheet_values(sheetsService, "networkpoints")
	if result_networkpoints == nil {
		t.Fatal("Networkpoints failed")
	}

	networkpoints := process_networkpoints_and_internal(result_networkpoints)
	fmt.Println(networkpoints)
}

func Test_process_internal(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()
	result_internal := get_sheet_values(sheetsService, "internal")
	if result_internal == nil {
		t.Fatal("internal failed")
	}

	internal := process_networkpoints_and_internal(result_internal)
	fmt.Println(internal)

}

// -------------------------- Get sheet values tests -------------------------- //
func Test_get_sheet_values_networkpoints(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()
	result_networkpoints := get_sheet_values(sheetsService, "networkpoints")
	if result_networkpoints == nil {
		t.Fatal("Networkpoints failed")
	}
}

func Test_get_sheet_values_level1(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()

	result_level1 := get_sheet_values(sheetsService, "level1")
	if result_level1 == nil {
		t.Fatal("Level1 failed")
	}
}

func Test_get_sheet_values_level2(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()

	result_level2 := get_sheet_values(sheetsService, "level2")
	if result_level2 == nil {
		t.Fatal("Level2 failed")
	}

}

func Test_get_sheet_values_level3(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()

	result_level3 := get_sheet_values(sheetsService, "level3")
	if result_level3 == nil {
		t.Fatal("Level3 failed")
	}
}

func Test_get_sheet_values_internal(t *testing.T) {
	var sheetsService = setUpGoogleSheetsAPI()

	result_internal := get_sheet_values(sheetsService, "internal")
	if result_internal == nil {
		t.Fatal("internal failed")
	}
}

// -------------------------- End -------------------------- //
