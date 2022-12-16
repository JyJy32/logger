# logger
This is a logger package in Go language. It provides five functions - Http, Debug, Warning, Mqtt and Error - to log messages at different levels.

Each function takes a variable number of arguments of type any, which is a placeholder for any type in Go. It means that you can pass arguments of any type to these functions.

The package also defines five global variables - http, mosquitto, debug, warning, and error - that hold prefixes for different logger levels. These prefixes are colored using ANSI escape codes, which are used to add color to the terminal output.

The Http function logs a message with the "HTTP" prefix, the Debug function logs a message with the "DEBG" prefix, the Warning function logs a message with the "WARN" prefix, and the Error function logs a message with the "ERR" prefix. The Mqtt function logs a message with the "MQTT" prefix.

The time at which the message was logged is also included in the output. It is formatted as hh:mm:ss using the time package's Clock function.

-ChatGPT
