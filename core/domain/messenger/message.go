package messenger

type Message struct {
	Topic string
	Key []byte
	Value []byte
}
