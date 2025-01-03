# Expand CIDR

A simple Go utility that expands CIDR notation into individual IP addresses. 

## Installation

### Using `go install`

```bash
go install github.com/syndrome-impostor/expand-cidr@latest
```

### Building from source

```bash
git clone https://github.com/syndrome-impostor/expand-cidr.git
cd expand-cidr
go build -o expand-cidr
```

## Usage


#### Expand a CIDR block
```bash
expand-cidr 192.168.1.0/24
```

#### Expand multiple IPs and/or CIDRs
```bash
expand-cidr 192.168.1.1 10.0.0.0/28 172.16.0.100/29
```

#### Process single IP 
*Returns the same IP, useful for batch processing*
```bash
expand-cidr 192.168.1.1
```

#### Process multiple CIDR blocks from a file, output to stdout:

```bash
cat cidrs.txt | expand-cidr > output.txt
```

Example `cidrs.txt`, newline delimited:
```bash
192.168.1.0/30
10.0.0.1
172.16.0.0/29
```


### Example Output

Input:
```bash
expand-cidr 192.168.1.0/30
```

Output:
```
192.168.1.0
192.168.1.1
192.168.1.2
192.168.1.3
```

Multiple inputs:
```bash
expand-cidr 192.168.1.1 10.0.0.0/31
```

Output:
```bash
192.168.1.1
10.0.0.0
10.0.0.1
```

## License

MIT License - see [LICENSE](LICENSE) file for details