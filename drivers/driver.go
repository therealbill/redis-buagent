package drivers

type Driver interface {
	Connect() bool
	Authenticate() bool
	Upload(rdb []byte) bool
}
