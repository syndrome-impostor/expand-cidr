package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func expandCIDR(input string) []string {
	// Try parsing as a single IP first
	if ip := net.ParseIP(input); ip != nil {
		if ipv4 := ip.To4(); ipv4 != nil {
			return []string{ipv4.String()}
		}
		return nil
	}

	// Parse as CIDR
	ip, ipnet, err := net.ParseCIDR(input)
	if err != nil {
		return nil
	}

	// Ensure IPv4
	if ip = ip.To4(); ip == nil {
		return nil
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips
}

func processInput(input string, writer *bufio.Writer) {
	if ips := expandCIDR(input); ips != nil {
		for _, ip := range ips {
			fmt.Fprintln(writer, ip)
		}
	}
}

func main() {
	// Create a buffered writer for output
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Check if we're receiving input from pipe
	if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			processInput(scanner.Text(), writer)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Handle command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: expand-cidr <cidr or ip> [cidr or ip...]")
		fmt.Println("Examples:")
		fmt.Println("  expand-cidr 192.168.1.0/24")
		fmt.Println("  expand-cidr 192.168.1.1 10.0.0.0/28 172.16.1.1")
		fmt.Println("  cat cidrs.txt | expand-cidr > output.txt")
		return
	}

	for _, input := range os.Args[1:] {
		processInput(input, writer)
	}
}
