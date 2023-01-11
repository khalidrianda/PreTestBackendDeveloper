package entity

// 2. Buatkan struct di Go dengan data
type Data struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       struct {
		InfluencedBy []string `json:"influenced-by"`
		Influences   []string `json:"influences"`
	} `json:"relation"`
}

type Services interface {
	Palindrome(n string) string
}
