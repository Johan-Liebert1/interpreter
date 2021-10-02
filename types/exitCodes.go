package types

var ExitCodes = map[string]bool{
	"exit":   true,
	"\\q":    true,
	"exit()": true,
}
