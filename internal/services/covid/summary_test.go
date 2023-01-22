package covid

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Natt-S10/LMWN-assesment/internal/models/covid"
	"github.com/Natt-S10/LMWN-assesment/internal/repository"
	"github.com/go-playground/assert/v2"
)

const testingJSON = `{
    "Data": [
        {
            "ConfirmDate": "2021-05-04",
            "No": null,
            "Gender": "หญิง",
            "GenderEn": "Female",
            "Nation": null,
            "NationEn": "China",
            "Province": "Phrae",
            "ProvinceId": 46,
            "District": null,
            "ProvinceEn": "Phrae",
            "StatQuarantine": 5
        },
        {
            "ConfirmDate": "2021-05-01",
            "No": null,
        	"Age": 51,
            "Gender": "ชาย",
            "GenderEn": "Male",
            "Nation": null,
            "NationEn": "India",
            "Province": "Suphan Buri",
            "ProvinceId": 65,
            "District": null,
            "StatQuarantine": 8
        },
        {
            "ConfirmDate": "2021-05-01",
            "No": null,
            "Age": 79,
            "Gender": null,
            "GenderEn": null,
            "Nation": null,
            "NationEn": "India",
            "Province": "Roi Et",
            "ProvinceId": 53,
            "District": null,
            "ProvinceEn": "Roi Et",
            "StatQuarantine": 1
        },
        {
            "ConfirmDate": "2021-05-04",
            "No": null,
            "Age": 52,
            "Gender": "หญิง",
            "GenderEn": "Female",
            "Nation": null,
            "NationEn": "Thailand",
            "Province": "Chumphon",
            "ProvinceId": 12,
            "District": null,
            "ProvinceEn": "Chumphon",
            "StatQuarantine": 8
        },
        {
            "ConfirmDate": "2021-05-04",
            "No": null,
            "Age": 99,
            "Gender": "หญิง",
            "GenderEn": "Female",
            "Nation": null,
            "NationEn": "China",
            "Province": "Bangkok",
            "ProvinceId": 0,
            "District": null,
            "ProvinceEn": "Bangkok",
            "StatQuarantine": 18
        }
	]
}`

func TestSummarizeStat(t *testing.T) {
	expected := covid.CovidSummary{
		AgeGroup: covid.AgeGroup{Young: 0, Middle: 2, Old: 2, NA: 1},
		Province: make(map[string]int),
	}
	expected.Province["Bangkok"] = 1
	expected.Province["Chumphon"] = 1
	expected.Province["Roi Et"] = 1
	expected.Province["N/A"] = 1
	expected.Province["Phrae"] = 1

	var raw repository.CovidStat
	if err := json.Unmarshal([]byte(testingJSON), &raw); err != nil {
		fmt.Println("Can't unmarshall testing json")
	}

	actual := summarizeStat(raw)
	assert.Equal(t, expected, actual)
}
