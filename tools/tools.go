package tools
import(
	"time"
	"fmt"
)

func MySQLDate () string {
	timeStorage := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			timeStorage.Year(),timeStorage.Month(),
			timeStorage.Day(),timeStorage.Hour(),
			timeStorage.Minute(),timeStorage.Second()) 
}