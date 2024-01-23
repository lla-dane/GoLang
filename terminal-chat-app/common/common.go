package common

const (
	ServerAddress = "localhost:5000"
)

type Message struct {
	Sender    string
	Content   string
	Timestamp string
}
