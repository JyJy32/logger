package logger

import (
	"fmt"
    "runtime"
    "time"
)

//prefix for logger level
var http string = ""
var mosquitto string = ""
var debug string = ""
var warning string = ""
var error string = ""

func Init() {
    if runtime.GOOS == "windows" {
        http = "[HTTP]"
        mosquitto = "[MQTT]"
        debug = "[DEBG]"
        warning = "[WARN]"
        error = "[ERR ]"
        Debug("no colours on windows")
    } else {
        http = "[\033[32;1mHTTP\033[0m]"
        mosquitto = "[\033[34;1mMQTT\033[0m]"
        debug = "[\033[36mDEBG\033[0m]"
        warning = "[\033[33mWARN\033[0m]"
        error = "[\033[31mERR\033[0m ]"
    }
}

func Http(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, http, msg)
}
func Debug(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, debug, msg)
}
func Warning(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, warning, msg)
}
func Error(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, error, msg)
	panic(msg)
}
func Mqtt(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, mosquitto, msg)
}
