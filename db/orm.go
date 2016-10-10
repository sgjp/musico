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
	AvgPrice     int
	Reviews      []Review
	Comments     []Comment
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

type Comment struct {
	Id      string
	Comment string
	Type    int
}
