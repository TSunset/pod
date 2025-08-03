package messageserver

type Message struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Text string `json:"text"`
}
