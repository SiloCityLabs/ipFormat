package ipformat

import (
	"errors"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// IP details about the IP address
type IP struct {
	Address string // Original/full notation passed in
	Parts   []string
	TypeV6  bool
	Range   bool
	CIDR    int64
}

// ReadUserIP gets the users IP address from the request
// @TODO: fix this to handle IPS properly building on logic from https://husobee.github.io/golang/ip-address/2015/12/17/remote-ip-go.html
func ReadUserIP(r *http.Request) string {

	IPAddress := r.Header.Get("X-Forwarded-For")
	if IPAddress != "" {
		IPS := strings.Split(IPAddress, ",")
		IPAddress = IPS[0]
	}
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Real-IP")
	}
	if IPAddress == "" {
		IPAddress, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return IPAddress
}

// New creates the internal IP struct
func New(ipString string) (IP, error) {
	var ip IP

	// Set what the user sent us
	ip.Address = ipString

	//TODO: use regex validation?
	// check to see if its already an ipv6
	colonCount := strings.Count(ip.Address, ":")
	ip.TypeV6 = colonCount >= 2

	if colonCount > 7 || colonCount == 1 {
		return IP{}, errors.New("not a valid ip address")
	}

	//TODO: validate against "ip - ip"
	// Check if its a range
	ip.Range = strings.Contains(ip.Address, "/")

	// Grab the cidr
	cidrParts := strings.Split(ip.Address, "/")
	if ip.Range {
		var err error
		ip.CIDR, err = strconv.ParseInt(cidrParts[1], 10, 64)

		if err != nil {
			return IP{}, errors.New("cannot parse cidr notation")
		}
	}

	if len(cidrParts) > 2 {
		return IP{}, errors.New("not a valid ip address cidr notation")
	}

	// Split the parts
	if ip.TypeV6 {
		ip.Parts = strings.Split(cidrParts[0], ":")
	} else {
		ip.Parts = strings.Split(cidrParts[0], ".")
	}

	return ip, nil
}
