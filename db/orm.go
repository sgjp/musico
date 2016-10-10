package db

type User struct {
	Id       int
	Username string
	Password string
}

type Band struct {
	Id           int
	Name         string
	Genre        string
	Youtube      string
	Facebook     string
	Requirements string
	Location     string
	AvgPrice     string
	Reviews	     []Review
}

type Review struct {
	Id              string
	Comment         string
	RateQuality     string
	RatePunctuality string
	RateFlexibility string
	RateEnthusiasm  string
	RateSimilarity  string
	Rate            string
}
