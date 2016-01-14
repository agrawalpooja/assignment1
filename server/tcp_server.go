package main 
import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing 
import "strconv"
func writeFile(filename string, numBytes int, expTimeStr int){
	fmt.Println(filename+":"+strconv.Itoa(numBytes)+":"+strconv.Itoa(expTimeStr))
	
}
func parseCommand(command string){
	var strArr []string
	strArr=strings.Split(command," ")
	if strArr[0]=="write" {
		expT :=-1
		if len(strArr)>3{
			expT,_=strconv.Atoi(strArr[3][:len(strArr[3])-1])
		}
		numBytes,_:=strconv.Atoi(strArr[2])
		fmt.Print(expT)
		writeFile(strArr[1],numBytes,expT)
	}
}


func serverMain() {
	fmt.Println("Launching server...")   // listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")   // accept connection on port
	conn, _ := ln.Accept()   // run loop forever (or until ctrl-c)   
	for {     
		// will listen for message to process ending in newline (\n)     
		message, _ := bufio.NewReader(conn).ReadString('\n')     
		// output message received     
		fmt.Print("Message Received1:", string(message))     
		parseCommand(string(message))
		// sample process for string received     
		newmessage := strings.ToUpper(message)     
		// send new string back to client     
		conn.Write([]byte(newmessage + "\n"))   
	} 
}
func main() {
	serverMain()
}
