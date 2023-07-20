package covid

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Natt-S10/LMWN-assesment/config"
	"github.com/Natt-S10/LMWN-assesment/internal/models/covid"
	"github.com/Natt-S10/LMWN-assesment/internal/repository"
)

func FetchCovidStats() (repository.CovidStat, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	res, err := client.Get(config.COVID_STAT_API)

	if err != nil {
		fmt.Println("Couldn't fetch covid status")
		return repository.CovidStat{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var covidStat repository.CovidStat
	if err := json.Unmarshal(body, &covidStat); err != nil {
		fmt.Println("Can't unmarshall covid stat")
		return repository.CovidStat{}, err
	}

	return covidStat, nil
}

func summarizeStat(covidStat repository.CovidStat) covid.CovidSummary {
	covidSummary := covid.CovidSummary{
		AgeGroup: covid.AgeGroup{},
		Province: make(map[string]int),
	}
	for _, record := range covidStat.Data {
		// count by age
		if record.Age != nil { // ensure that Age is not absent
			age := *record.Age
			if age <= 30 {
				covidSummary.AgeGroup.Young++
			} else if age <= 60 {
				covidSummary.AgeGroup.Middle++
			} else if age >= 61 {
				covidSummary.AgeGroup.Old++
			}
		} else { // if it is nil
			covidSummary.AgeGroup.NA++
		}

		// count by province
		if record.Province != "" {
			covidSummary.Province[record.Province]++
		} else {
			covidSummary.Province["N/A"]++
		}
	}
	// fmt.Println(covidSummary)
	return covidSummary
}

func GetCovidSummary() covid.CovidSummary {
	covidStat, err := FetchCovidStats()
	if err != nil {
		fmt.Println("Can't fetch covid stat")
		fmt.Println(err)
		return covid.CovidSummary{}
	}

	covidSummary := summarizeStat(covidStat)
	// fmt.Println(covidSummary)
	return covidSummary
}

// adwdawda
