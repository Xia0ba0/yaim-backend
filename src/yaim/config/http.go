package config

import (
	"time"
	"fmt"
	"net"
)

func Ips() (string, error) {

    ips :=  make(map[string]string)

    interfaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }

    for _, i := range interfaces {
        byName, err := net.InterfaceByName(i.Name)
        if err != nil {
            return "", err
        }
        addresses, err := byName.Addrs()
        for _, v := range addresses {
            ips[byName.Name] = v.String()
        }
    }
    ip := ips["WLAN"]
    ip = ip[:len(ip)-3]
    fmt.Println("Yout WLAN IP is: " + ip)
    return ip, nil
}

var localIP, _ = Ips()
var ServerAddr = "http://" + localIP + Port

const (
	HostName = "http://localhost:9080"
	Port     = ":8090"

	CookieName    = "YaimSession"
	CookieExpires = 24 * time.Hour
	UserIdKey     = "userid"

	SMTPServer   = "smtp.163.com:25"
	SMTPAccount  = "m18569002382@163.com"
	SMTPPassword = "4702391byl"
	SMTPSubject = "Yaim Account Verification"

	TokenKey = "Yaim?:@$%"
)