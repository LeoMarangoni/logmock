package logger

var version = "0.0.1"

func Version() {
	Print(version, "INFO")
}
