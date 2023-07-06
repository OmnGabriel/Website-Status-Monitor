package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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

func showIntro() {
	fmt.Println("Website monitoring application")
	version := 0.1
	fmt.Print("This applications is on version: ", version, "\n\n")
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Delete logs")
	fmt.Println("0 - Leave The Program")
}

func readCommand() int {
	var commandWasRead int
	// Também é válido: fmt.Scanf("%d", &comand)
	fmt.Scan(&commandWasRead)

	return commandWasRead
}

func startMonitoring() {
	fmt.Print("\nMonitoring sites... \n")

	for i := 0; i < monitoring; i++ {
		fmt.Print("\n", i+1, "˚ first round of monitoring \n\n")
		for i, site := range readingFilesSites() {
			fmt.Println("Testing sites", i, ": ", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println(" ")

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
		fmt.Println("[", resp.StatusCode, "]", "Unable to open website:")
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
		fmt.Print(sites, "\n")
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
	file, err := ioutil.ReadFile("log.txt")
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
