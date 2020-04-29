package ipformat

// ToV4 converts ipv6 to ipv4
// func (ip IP) ToV4() (IP, error) {
// 	if ip.TypeV6 == false {
// 		return ip, nil
// 	}

// 	// start of ipv4
// 	var newParts []string

// 	// Loop through the old parts converting to ipv4
// 	for i, part := range ip.Parts {

// 		// turn the string into a int
// 		intConverted, _ := strconv.ParseInt(part, 16, 0)
// 		ip.Parts[i] = strconv.FormatInt(intConverted, 10)

// 		// Make sure that no 0s are missing
// 		if len(ip.Parts[i]) == 1 {
// 			ip.Parts[i] = "0" + ip.Parts[i]
// 		}
// 	}

// 	//Set the parts to ipv6
// 	newParts = append(newParts, ip.Parts[0]+ip.Parts[1])
// 	newParts = append(newParts, ip.Parts[2]+ip.Parts[3])
// 	ip.Parts = newParts

// 	ip.TypeV6 = true

// 	ip.Address = strings.Join(ip.Parts, ":")

// 	if ip.Range {
// 		// Convert cidr to ipv6
// 		ip.CIDR = 128 - (32 - ip.CIDR)
// 		ip.Address = ip.Address + "/" + strconv.FormatInt(ip.CIDR, 10)
// 	}

// 	// return the new ip address
// 	return ip, nil
// }
