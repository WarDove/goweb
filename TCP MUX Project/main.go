package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type person struct {
	Id        int
	FirstName string
	LastName  string
	License   bool
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func handle(conn net.Conn, rawData map[int]person) error {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	var firstLine bool
	var method string
	var uri string

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if !firstLine {
			firstLine = true
			method = strings.Fields(line)[0]
			uri = strings.Fields(line)[1]
			if method != "GET" {
				return errors.New("unsupported http method")
			}
		}

		if line == "" {
			// Request line and Headers are done
			// Beginning to respond
			intUri, err := strconv.Atoi(string([]byte(uri)[1:]))
			if err != nil {
				fmt.Fprintf(conn, "HTTP/1.1 404 Uri not Found\r\n\r\n")
				break
			}
			if val, ok := rawData[intUri]; ok {

				// Some Experimental Reminders :)
				//var jsonStr strings.Builder
				//json.NewEncoder(&jsonStr).Encode(val)
				//fmt.Println(jsonStr.String())

				//bs, _ := json.Marshal(val)

				fmt.Fprintf(conn, "HTTP/1.1 202 OK\r\n")
				fmt.Fprintf(conn, "Content-Length: %d\r\n", 100)
				fmt.Fprintf(conn, "Content-Type: application/json\r\n")
				fmt.Fprintf(conn, "\n")
				json.NewEncoder(conn).Encode(val)
			} else {
				fmt.Fprintf(conn, "HTTP/1.1 404 ID not Found\r\n\r\n")
				fmt.Fprintf(conn, "\n")
			}

			break
		}
	}

	return nil
}

func main() {

	peopleMap := make(map[int]person)

	people := []person{
		person{0, "Tarlan", "Huseynov", true},
		person{1, "Kamran", "Huseynov", true},
		person{2, "Sevil", "Tarlanova", true},
	}

	for _, v := range people {
		peopleMap[v.Id] = v

	}

	li, err := net.Listen("tcp", ":8080")
	check(err)
	defer li.Close()
	for {
		connection, err := li.Accept()
		check(err)
		go handle(connection, peopleMap)

	}

}
