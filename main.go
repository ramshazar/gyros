package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "time"
  "net/http"
)

func main() {
        logmessage, err := notmain()
        if err != nil {
                fmt.Println("[ERROR] ", err)
        } else {
                fmt.Println("[*] Output-> \n", logmessage)
        }
}

func notmain() (string, error) {

        client := &http.Client{
        Timeout: 3 * time.Second,
        }

        req, err := http.NewRequest(http.MethodGet, "http://metadata.google.internal/computeMetadata/v1/?recursive=true&alt=json", nil)
        req.Header.Add("Metadata-Flavor", "Google")
        if err != nil {
                return "", err
        }

        resp, err := client.Do(req)
        if err != nil {
                return "", err
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        bodyString := string(body)

        return bodyString, nil
}
