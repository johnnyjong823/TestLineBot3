package main

import (  
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func Send() string{  
    var jsonStr = []byte('{}')

    url := "http://saappd.cloudapp.net/Line/WebService1.asmx/HelloWorld"

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    return string(body)
}