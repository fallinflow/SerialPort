package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("Error opening serial port: ", err)
		return
	}

	defer s.Close()

	reader := bufio.NewReader(s)
	fmt.Println("Serial Port !")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from serial port: ", err)
			break
		}
		fmt.Print(line)
	}
}
