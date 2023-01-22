package repository

type CovidStatRecord struct {
	Age *int `json:"Age"`
	// Nationality string `json:"NationEn"`
	Province string `json:"ProvinceEn"`
}

type CovidStat struct {
	Data []CovidStatRecord `json:"Data"`
}
