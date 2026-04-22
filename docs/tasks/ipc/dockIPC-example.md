# **?од?л? dockIPC**

?еализа?и? мод?л? дл? ?або?? ? IPC ?е?ез Unix domain sockets:

```go
package dockIPC

import (
	"errors"
	"net"
	"os"
	"syscall"
)

// StartServer - зап??кае? IPC ?е?ве?
// fileName: п??? к ?оке?? (нап?име? "/tmp/hypr-appdock.sock")
// handler: ??нк?и? об?або?ки в?од??и? команд
func StartServer(fileName string, handler func(string) ([]byte, error)) error {
	// Удал?ем ??а??й ?оке? е?ли ???е??в?е?
	if err := os.RemoveAll(fileName); err != nil {
		return err
	}

	// Создаем Unix domain socket
	listener, err := net.Listen("unix", fileName)
	if err != nil {
		return err
	}
	defer listener.Close()

	// У??анавлием п?ава на ?оке?
	if err := os.Chmod(fileName, 0666); err != nil {
		return err
	}

	// ?б?або?ка в?од??и? ?оединений
	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			continue
		}

		go handleConnection(conn, handler)
	}
}

// handleConnection об?аба??вае? одно ?оединение
func handleConnection(conn net.Conn, handler func(string) ([]byte, error)) {
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	command := string(buf[:n])
	response, err := handler(command)
	if err != nil {
		// Фо?ма?и??ем о?ибк? в ??анда??н?й ?о?ма?
		response = []byte("error: " + err.Error() + "\n")
	}

	conn.Write(response)
}

// Send о?п?авл?е? команд? в IPC ?оке?
// command: ???ока команд? (нап?име? "j/layer get")
func Send(fileName string, command string) ([]byte, error) {
	conn, err := net.Dial("unix", fileName)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(command + "\n"))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil
}

// StopServer о??анавливае? ?е?ве? (в?помога?ел?на? ??нк?и?)
func StopServer(fileName string) error {
	return syscall.Unlink(fileName)
}
```

## **??пол?зование мод?л?**

### **1. Се?ве?на? ?а???**
```go
package main

import (
	"fmt"
	"dockIPC"
)

func main() {
	// ?б?або??ик команд
	handler := func(command string) ([]byte, error) {
		switch command {
		case "j/layer get":
			return []byte(`{"layers":["bottom","top"]}`), nil
		case "layer get":
			return []byte("bottom\ntop"), nil
		default:
			return nil, fmt.Errorf("unknown command")
		}
	}

	// ?ап??к ?е?ве?а
	err := dockIPC.StartServer("/tmp/hypr-appdock.sock", handler)
	if err != nil {
		panic(err)
	}
}
```

### **2. ?лиен??ка? ?а???**
```go
package main

import (
	"fmt"
	"dockIPC"
)

func main() {
	// ??п?авка команд?
	response, err := dockIPC.Send("/tmp/hypr-appdock.sock", "j/layer get")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(response))
}
```

## **??обенно??и ?еализа?ии**

1. **?езопа?но???**:
   - Удаление ??а?ого ?оке?а пе?ед ?озданием
   - У??ановка п?ав 0666 на ?оке?
   - ?о??ек?на? об?або?ка зак???и? ?оединений

2. **??оизводи?ел?но???**:
   - ?аждое ?оединение об?аба??вае??? в о?дел?ной goroutine
   - ???е?изи?ованное ??ение (4096 бай?)

3. **?ибко???**:
   - ?б?або??ик команд може? возв?а?а?? л?б?е бина?н?е данн?е
   - ?одде?жка ?ек??ов?? и JSON-команд

4. **??помога?ел?н?е ??нк?ии**:
   - `StopServer()` дл? ко??ек?ного заве??ени?
   - ?в?ома?и?е?кое ?о?ма?и?ование о?ибок

?од?л? го?ов к ин?ег?а?ии в п?оек? `hypr-appdock-ctl` и може? б??? ?а??и?ен п?и необ?одимо??и.
