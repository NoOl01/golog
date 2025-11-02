package log_data

type LogData struct {
	Level     []byte
	Timestamp []byte
	Caller    []byte
	Literal   []byte
	Name      []byte
	Content   []byte
}
