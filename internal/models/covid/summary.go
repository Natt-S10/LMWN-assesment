package covid

type AgeGroup struct {
	Young  int `json:"0-30"`
	Middle int `json:"31-60"`
	Old    int `json:"61+"`
	NA     int `json:"N/A"`
}

type CovidSummary struct {
	AgeGroup AgeGroup       `json:"AgeGroup"`
	Province map[string]int `json:"Province"`
}

// just generate the struct then use ctx.status(...).JSON(theStruct)
