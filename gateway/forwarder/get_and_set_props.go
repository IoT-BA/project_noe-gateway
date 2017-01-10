package main

import (
    "bufio"
    "fmt"
    "net"
    "net/http"
    "net/url"
    "os"
    "strconv"
)

type Settings struct {
    gatewayId          string
    serverAddress      string
    servPortUp         int64
    servPortDown       int64
    keepAliveInterval  int64
    statInterval       int64
    pushTimeout_ms     int64
    forwardCrcValid    bool
    forwardCrcError    bool
    forwardCrcDisabled bool
}

func main() {
    var owner string = os.Getenv("GATEWAY_OWNER")
    var region string = os.Getenv("GATEWAY_REGION")
    var company string = os.Getenv("GATEWAY_COMPANY")
    var address string = os.Getenv("GATEWAY_ADDRESS")
    var gpsLat string = os.Getenv("GATEWAY_GPS_LAT")
    var gpsLon string = os.Getenv("GATEWAY_GPS_LON")
    reader := bufio.NewReader(os.Stdin)
    var err error
    // check owner variable
    if len(owner) == 0 {
        fmt.Printf("Enviroment variable is not set, please write gateway owner:")
        owner, err = reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Something happend while reading output: %v", err)
        }
    }
    // check region variable
    if len(region) == 0 {
        fmt.Printf("Enviroment variable is not set, please write gateway region:")
        region, err = reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Something happend while reading output: %v", err)
        }
    }
    // check company variable
    if len(company) == 0 {
        fmt.Printf("Enviroment variable is not set, please write gateway company:")
        company, err = reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Something happend while reading output: %v", err)
        }
    }
    // check address variable
    if len(address) == 0 {
        fmt.Printf("Enviroment variable is not set, please write gateway address:")
        address, err = reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Something happend while reading output: %v", err)
        }
    }
    // check gpsLat variable
    if len(gpsLat) == 0 {
        fmt.Printf("Enviroment variable is not set, please write gateway gpsLat:")
        gpsLat, err = reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Something happend while reading output: %v", err)
        }
    }
    // check gpsLon variable
    if len(gpsLon) == 0 {
        fmt.Printf("Enviroment variable is not set, please write gateway gpsLon:")
        gpsLon, err = reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Something happend while reading output: %v", err)
        }
    }
    apiUrl := "http://localhost:3000"
    resource := "/gateway/register/"
    data := url.Values{}
    data.Set("owner", owner)
    data.Add("region", region)
    data.Add("company", company)
    data.Add("address", address)
    data.Add("gpsLat", gpsLat)
    data.Add("gpsLon", gpsLon)

    //get ip addrs from interfaces
    ipInterfaces, _ := net.Interfaces()

    for i := 0; i < len(ipInterfaces); i++ {
        //obtain ip addresses
        ipAddrs, _ := ipInterfaces[i].Addrs()
        fmt.Printf("%d : %v", i, ipInterfaces[i])
        for j := 0; j < len(ipAddrs); j++ {
            fmt.Printf("\t %d.%d --> %v\n", i, j, ipAddrs[j])
        }
    }

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    u.RawQuery = data.Encode()
    urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/?name=foo&surname=bar"

    client := &http.Client{}
    r, _ := http.NewRequest("GET", urlStr, nil)
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    resp, _ := client.Do(r)
    fmt.Println(resp.Status)
}
