package common

type Message interface {
	GetType() string
	GetHash() string
	Serialize() ([]byte, error)
	String() string
	IsWellFormed([]byte) error
	Equal(Message) bool
	Source() string
	// Validate(storage.LevelDBBackend) error
}
