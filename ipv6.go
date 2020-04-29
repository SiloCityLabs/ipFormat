package ipformat

import (
	"strconv"
	"strings"
)

// ToV6 converts ipv4 to ipv6
func (ip IP) ToV6() (IP, error) {
	// Return ip if its already v6 or if there are not enough parts to make a proper v6
	if ip.TypeV6 || len(ip.Parts) < 4 {
		return ip, nil
	}

	// Start of ipv6
	newParts := []string{"", "", "ffff"}

	// Loop through the old parts converting to ipv6
	for i, part := range ip.Parts {

		// turn the string into a int
		intConverted, _ := strconv.ParseInt(part, 10, 0)
		ip.Parts[i] = strconv.FormatInt(intConverted, 16)

		// Make sure that no 0s are missing
		if len(ip.Parts[i]) == 1 {
			ip.Parts[i] = "0" + ip.Parts[i]
		}
	}

	// Set the parts to ipv6
	newParts = append(newParts, ip.Parts[0]+ip.Parts[1])
	newParts = append(newParts, ip.Parts[2]+ip.Parts[3])
	ip.Parts = newParts

	ip.TypeV6 = true

	ip.Address = strings.Join(ip.Parts, ":")

	if ip.Range {
		// Convert cidr to ipv6
		ip.CIDR = 128 - (32 - ip.CIDR)
		ip.Address = ip.Address + "/" + strconv.FormatInt(ip.CIDR, 10)
	}

	// return the new ip address
	return ip, nil
}

// CompressV6 takes all formats of ipv6 and strips them of uneeded characters
func (ip IP) CompressV6() (IP, error) {

	// start of ipv6
	var ipParts []string

	// loop through the parts and make sure that no 0s are missing
	for _, part := range ip.Parts {
		ipParts = append(ipParts, strings.TrimLeft(part, "0"))
	}

	ip.Address = strings.Join(ipParts, ":")

	if ip.Range {
		ip.Address = ip.Address + "/" + strconv.FormatInt(ip.CIDR, 10)
	}

	return ip, nil
}
