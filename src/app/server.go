package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

// Initialize all templates
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// Package level struct pointer
var (
	requestData = newHTTPRequestData()
	tpl         *template.Template
)

// Define networking param
const (
	port    = ":8080"
	network = "tcp"
)

// HTTP interface
type HTTP interface {
	newHTTPRequestData()
}

// HTTPRequestData is collected data from HTTP request
type HTTPRequestData struct {
	Method      string
	RequestLine string
	StatusLine  string
}

// HTTPReponse reponse to client request
type HTTPReponse struct {
	Route string
	Body  string
}

// NewHTTPRequestData will return a default initialized struct
func newHTTPRequestData() *HTTPRequestData {
	return &HTTPRequestData{
		Method:      "",
		RequestLine: "",
		StatusLine:  "",
	}
}

func processRequest() {

}

// Request function will chop up http request into respected fields
// defined within RFC7230. return mapped request
func request(conn net.Conn) {
	lineCount := 0

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if lineCount == 0 {
			tok := strings.Fields(ln)
			requestData.Method = string(tok[0])
			requestData.RequestLine = string(tok[1])
			requestData.StatusLine = string(tok[2])
		} else if ln == "" {
			// headers are done
			break
		}
		lineCount++
	}
}

// UPDATE CONDITIONAL STATEMENT WITH IF STATEMENTS
func respond(conn net.Conn) {

	if strings.Compare(requestData.RequestLine, "/") == 0 {
		home(conn)
	} else if strings.Compare(requestData.RequestLine, "/api") == 0 {
		// switch requestData.Method {
		// case "GET":
		// 	get(conn)
		// case "POST":
		// 	post(conn)
		// case "DELETE":
		// 	delete(conn)
		// case "PUT":
		// 	put(conn)
		// default:
		// 	notFound(conn)
		// }
	} else {
		notFound(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)

}

func main() {

	li, err := net.Listen(network, port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Server listening...")
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}

}

// Request handler
func get(conn net.Conn) {

	body := `get`

	fmt.Fprint(conn, "HTTP/1.1 404 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func post(conn net.Conn) {
	body := `post`

	fmt.Fprint(conn, "HTTP/1.1 404 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func delete(conn net.Conn) {
	body := `delete`

	fmt.Fprint(conn, "HTTP/1.1 404 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func put(conn net.Conn) {
	body := `put`

	fmt.Fprint(conn, "HTTP/1.1 404 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func notFound(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>404</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 404 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

// Routes for HTML handlers
func home(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>HOME</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
func about() {

}
