package helpers

import (
	"fmt"
	"os"
	"time"
)

func SendErrorLog(code int, message string) {
	timeNow := time.Now().Format("Monday 02/01/2006 15:04:05")
	logMessage := fmt.Sprintf("[%s]: [ERROR %d] %s\n", timeNow, code, message)

	file, err1 := os.OpenFile("publics/documents/error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err1 != nil {
		fmt.Println("error:", err1.Error())
	}

	_, err2 := file.WriteString(logMessage)
	if err2 != nil {
		fmt.Println("error:", err2.Error())
	}
}
