package main

import ("fmt" ; "io" ; "os" ; "log" ; "net")

func main(){
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Port Number missing");
	}

	PORT := ":" + arguments[1]
	l,err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		c,err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go hendleConnection(c)
	}
}

func hendleConnection(c net.Conn){
	fmt.Printf("Serving %S", c.RemoteAddr().String())
	packet := make([]byte, 4096)
	tmp := make([]byte,4096)
	defer c.Close()
	for {
		_, err := c.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
		}
		packet = append(packet, tmp...)
	}
	c.Write(packet)
}