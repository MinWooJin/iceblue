package storage // Item is struct of iceblue data
import "time"

type Item struct {
	Key    string
	KeyLen uint32
	Value  string
	ValLen uint32
	Hvalue uint32
	Time   time.Time
	Next   *Item
}
