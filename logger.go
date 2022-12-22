package logger

import (
	"fmt"
	"time"
)

//prefix for logger level
var http string = "[\033[32;1mHTTP\033[0m]"
var mosquitto string = "[\033[34;1mMQTT\033[0m]"
var debug string = "[\033[36mDEBG\033[0m]"
var warning string = "[\033[33mWARN\033[0m]"
var error string = "[\033[31mERR\033[0m ]"

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
