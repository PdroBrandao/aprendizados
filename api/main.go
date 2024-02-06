package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type requestBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResponseJSON struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// func https
func handlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Get variable in url
		k := r.URL.Query().Get("key")
		log.Println(k)
		// Get Value of variable
		e, err := newEngine()
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}
		v, err := e.Get(k)
		log.Println(v, err)
		if err != nil {
			responseJSON(w, ResponseJSON{Message: err.Error(), Status: "Error"}, http.StatusNotFound)
		} else {
			responseJSON(w, requestBody{Key: k, Value: v}, http.StatusOK)
		}
		// if bad, bad, if good, good
	} else {
		log.Panicln("Method not allowed")
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func handlerSet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Get Body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body:", err)
			fmt.Fprintf(w, "Error reading body")
			return
		}
		// Deserialize Body
		var requestBody requestBody
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			log.Println("Error deserializing body:", err)
			fmt.Fprintf(w, "Error deserializing body")
			return
		}
		// Save to file
		e, err := newEngine()
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}
		err = e.Set(requestBody.Key, requestBody.Value)
		if err != nil {
			log.Println("Error saving to file:", err)
			fmt.Fprintf(w, "Error saving to file")
			return
		}
		// Return
		responseJSON(w, ResponseJSON{
			Message: "Key value pair saved successfully.",
			Status:  "Sucess",
		}, http.StatusOK)

	} else {
		log.Println("Method not allowed")
		fmt.Fprintf(w, "Method not allowed")
	}
}

func responseJSON(w http.ResponseWriter, r interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error serializing response"))
		return
	}
	w.WriteHeader(status)
	w.Write(response)
	w.Write([]byte("\n"))
}

func main() {
	e, err := newEngine()
	e.Set("1", "1.ok")
	e.Set("567", "2.ok")
	v, err := e.Get("567")
	fmt.Println(v, err)

	// definitions http handlers
	http.HandleFunc("/get", handlerGet)
	http.HandleFunc("/set", handlerSet)

	// server
	adress := ":8888"

	// start server
	fmt.Println("Starting server on port", adress)
	err = http.ListenAndServe(adress, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

type Engine struct {
	data    map[string]int64
	mu      sync.Mutex
	file    *os.File
	keys    []string
	offsets map[string]int64
}

func newEngine() (*Engine, error) {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file data:", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var keysRecovered []string
	offsets := make(map[string]int64)

	// Move cursor to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking text:", err)
		return nil, err
	}

	for scanner.Scan() {
		var key, v string
		line := scanner.Text()
		offset, _ := file.Seek(0, io.SeekCurrent) // Obtém o offset atual
		_, err = fmt.Sscanf(line, "%s %s", &key, &v)
		if err != nil {
			fmt.Println("Error scanning text:", err)
			return nil, err
		}
		keysRecovered = append(keysRecovered, key)
		offsets[key] = offset // Armazena o offset associado à chave
		// Avança o cursor do arquivo para o final da linha
		_, err = file.Seek(int64(len(line)+1), io.SeekCurrent)
		if err != nil {
			fmt.Println("Error seeking text:", err)
			return nil, err
		}
	}

	// Check error in read file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Print data
	for _, d := range keysRecovered {
		fmt.Println("Keys :", d)
	}

	for _, o := range offsets {
		fmt.Println("Offsets :", o)
	}

	return &Engine{
		data:    make(map[string]int64),
		mu:      sync.Mutex{},
		file:    file,
		keys:    keysRecovered,
		offsets: offsets,
	}, nil
}

var keyValueSeparator = " "

func (e *Engine) Set(key string, value string) error {
	// e.mu.Lock()
	// defer e.mu.Unlock()

	//Position cursor
	offset, err := e.file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println("Error seeking text:", err)
		return err
	}

	// Save data
	_, err = e.file.WriteString(key + keyValueSeparator + value + "\n")
	if err != nil {
		fmt.Println("Error appending text:", err)
		return err
	}

	e.data[key] = offset
	return nil
}

func (e *Engine) Get(key string) (string, error) {
	// e.mu.Lock()
	// defer e.mu.Unlock()

	var offset int64
	var ok bool
	for _, k := range e.keys {
		if k == key {
			log.Printf("Key recovered = %s || offset recovered = %d", k, e.offsets[k])
			offset, ok = e.offsets[k]
		}
	}
	if !ok {
		return "", errors.New("Key not found")
	}
	// use the Seek function to move the cursor to the offset of the key that is required by parameter on the method
	valueOffset, err := e.file.Seek(offset+int64(len(key))+1, 0)
	if err != nil {
		fmt.Println("Error seeking text:", err)
		return "", err
	}
	log.Printf("Key offset = %d || Value offset = %d ", offset, valueOffset)
	currentOffset, err := e.file.Seek(0, io.SeekCurrent)
	if err != nil {
		fmt.Println("Error seeking text:", err)
		return "", err
	}
	log.Printf("Current offset = %d", currentOffset)

	//create buffer to read the value
	buffer := make([]byte, 1)
	var c []byte

	for {
		i, err := e.file.Read(buffer)
		if err != nil {
			fmt.Println("Error reading text:", err)
			return "", err
		}
		if i == 0 {
			break
		}
		if buffer[0] == '\n' {
			break
		}
		c = append(c, buffer[0])
		log.Println(string(buffer[0]))
	}

	return string(c), nil
}
