package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	showIntro()

	for {
		showMenu()
		command := readCommand()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 3:
			deleteLogs()
		case 0:
			fmt.Println("Leaving the Program")
			os.Exit(0)
		default:
			fmt.Println("Don't recognize this command")
			os.Exit(-1)
		}
	}
}

func showIntro() string {
	intro := "\nWebsite Monitoring Application\nIs on version: 1.0\n\n"
	print(intro)
	return intro
}

func showMenu() string {
	menu := "1 - Start Monitoring\n2 - Show logs\n3 - Delete logs\n0 - Leave The Program\n\n"
	print(menu)
	return menu
}

func readCommand() int {
	var commandWasRead int
	fmt.Scan(&commandWasRead)

	return commandWasRead
}

func startMonitoring() {
	fmt.Print("\nMonitoring sites... \n")

	for i := 0; i < monitoring; i++ {
		fmt.Printf("\n%d˚ round of monitoring \n\n", i+1)
		loopForMonitoring(i)
	}
}

func loopForMonitoring(i int) int {
	for i, site := range readingFilesSites() {
		fmt.Printf("Testing site %d: %q\n", i+1, site)
		testSite(site)
		time.Sleep(delay * time.Second)
	}
	return i
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		errorHandling(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("[", resp.StatusCode, "]", "Successfully opened")
		registerLog(site, true)
		fmt.Println(" ")
	} else {
		fmt.Println("[", resp.StatusCode, "]", "Unable to open website")
		registerLog(site, false)
		fmt.Println(" ")
	}

}

func readingFilesSites() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	reader := bufio.NewReader(file)

	if err != nil {
		errorHandling(err)
	}

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func errorHandling(err error) string {

	erro, _ := fmt.Print("\nAn error occurred: \n ", err, "\n")

	return string(rune(erro))
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		errorHandling(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n\n")

}

func showLogs() {
	file, err := os.ReadFile("log.txt")
	fmt.Print("\nShowing logs \n")
	if err != nil {
		errorHandling(err)
	}

	fmt.Println(string(file))
}

func deleteLogs() {
	fmt.Print("\nDeleting all existing logs \n\n")
	os.Remove("log.txt")
}
