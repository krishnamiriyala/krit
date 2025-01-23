import (
	"fmt"
	"os"
	"time"
)

// Function to calculate the time ago from the provided timestamp
func timeAgo(timestampStr string) (string, error) {
	// Parse the timestamp string into time.Time
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		return "", fmt.Errorf("error parsing timestamp: %v", err)
	}

	// Get current time
	now := time.Now()

	// Calculate the difference between current time and provided timestamp
	duration := now.Sub(timestamp)

	// Get total number of seconds
	totalSeconds := int(duration.Seconds())

	// Calculate days, hours, minutes, and seconds
	days := totalSeconds / (24 * 3600)
	hours := (totalSeconds % (24 * 3600)) / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	// Build the string for display
	ago := ""
	if days > 0 {
		ago += fmt.Sprintf("%d day(s) ", days)
	}
	if hours > 0 || days > 0 {
		ago += fmt.Sprintf("%d hour(s) ", hours)
	}
	if minutes > 0 || hours > 0 || days > 0 {
		ago += fmt.Sprintf("%d minute(s) ", minutes)
	}
	ago += fmt.Sprintf("%d second(s) ago", seconds)

	return ago, nil
}

func main() {
	// Ensure the program has been called with an argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a timestamp string as an argument in the format 'YYYY-MM-DDTHH:MM:SSZ'")
		return
	}

	// Get the timestamp string from command-line argument
	timestampStr := os.Args[1]

	// Get the time ago
	ago, err := timeAgo(timestampStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the time ago
	fmt.Println(ago)
}
