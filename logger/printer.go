package logger

func Print(msg string, level string) {
	if level == "TRACE" {
		Trace.Println(msg)
	}
	if level == "DEBUG" {
		Debug.Println(msg)
	}
	if level == "INFO" {
		Info.Println(msg)
	}
	if level == "WARN" {
		Warn.Println(msg)
	}
	if level == "ERROR" {
		Error.Println(msg)
	}
}
