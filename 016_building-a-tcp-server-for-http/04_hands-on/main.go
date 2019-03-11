package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(ln, conn)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(ln string, conn net.Conn) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]

	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/sign-up" {
		signUp(conn)
	}
	if m == "POST" && u == "/sign-up" {
		signUpProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head>
				<body>
					<strong>INDEX</strong>
					<ul>
						<li><a href="/about">About</a></li>
						<li><a href="/contact">Contact</a></li>
						<li><a href="/sign-up">Sign Up</a></li>
					</ul>
				</body></html>`

	fmt.Println("***METHOD GET")
	fmt.Println("***URL /")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head>
				<body>
					<strong>ABOUT</strong>
					<ul>
						<li><a href="/">Index</a></li>
						<li><a href="/contact">Contact</a></li>
						<li><a href="/sign-up">Sign Up</a></li>
					</ul>
				</body></html>`

	fmt.Println("***METHOD GET")
	fmt.Println("***URL /about")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head>
				<body>
					<strong>CONTACT</strong>
					<ul>
						<li><a href="/">Index</a></li>
						<li><a href="/about">About</a></li>
						<li><a href="/sign-up">Sign Up</a></li>
					</ul>
				</body></html>`

	fmt.Println("***METHOD GET")
	fmt.Println("***URL /contact")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func signUp(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head>
				<body>
					<strong>SIGN UP</strong>
					<ul>
						<li><a href="/">Index</a></li>
						<li><a href="/about">About</a></li>
						<li><a href="/contact">Contact</a></li>
					</ul>
					<form method="POST" action="/sign-up">
						<input type="submit" value="sign-up">
					</form>
				</body></html>`

	fmt.Println("***METHOD GET")
	fmt.Println("***URL /sign-up")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func signUpProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head>
				<body>
					<strong>SIGN UP PROCESS</strong>
					<ul>
						<li><a href="/">Index</a></li>
						<li><a href="/about">About</a></li>
						<li><a href="/contact">Contact</a></li>
					</ul>
					
				</body></html>`

	fmt.Println("***METHOD POST")
	fmt.Println("***URL /sign-up")
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
