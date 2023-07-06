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

const monitoramento = 3
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
			fmt.Println("Showing Log's")
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
	name := "Gabriel"
	version := 0.1
	fmt.Println("Hello,", name)
	fmt.Println("This applicatios is on version:", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Log's")
	fmt.Println("0 - Leave The Program")
}

func readCommand() int {
	var commandWasRead int
	// Também é válido: fmt.Scanf("%d", &comand)
	fmt.Scan(&commandWasRead)
	fmt.Println("The chosen command was", commandWasRead)

	return commandWasRead
}

func startMonitoring() {
	fmt.Println(" ")
	fmt.Println("Sites in monitoring: ")
	fmt.Println(" ")

	readingFilesSites()

	for i := 0; i < monitoramento; i++ {
		fmt.Print("\n", i+1, "˚ rodada de monitoramento \n\n")
		for i, site := range readingFilesSites() {
			fmt.Println("Testando site", i, ": ", site)
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
		fmt.Println("[", resp.StatusCode, "]", "Carregado com sucesso")
		registerLog(site, true)
		fmt.Println(" ")
	} else {
		fmt.Println("[", resp.StatusCode, "]", "Não foi possivel abrir o site:")
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
		fmt.Println(sites)
	}
	file.Close()
	return sites
}

func errorHandling(err error) string {

	erro, _ := fmt.Println("Ocorreu um erro:", err)

	return string(rune(erro))
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		errorHandling(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
