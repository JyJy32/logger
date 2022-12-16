package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	http      string `default:"[\033[32;1mHTTP\033[0m]"`
	mosquitto string `default:"[\033[34;1mMQTT\033[0m]"`
	debug     string `default:"[\033[36mDEBG\033[0m]"`
	warning   string `default:"[\033[33mWARN\033[0m]"`
	error     string `defaylt:"[\033[31mERR\033[0m ]"`
}

func (L Logger) Http(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, L.http, msg)
}
func (L Logger) Debug(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, L.debug, msg)
}
func (L Logger) Warning(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, L.warning, msg)
}
func (L Logger) Error(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, L.error, msg)
}
func (L Logger) Mqtt(msg ...any) {
	hour, min, sec := time.Now().Clock()
	now := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("%s %s %v\n", now, L.mosquitto, msg)
}
