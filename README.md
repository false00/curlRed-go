# curlRed-go

This program lookups ip address (both ipv4 and ipv6) and domain names against the curl.red api. The curl.red API returns geolocation and domian information. 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine.  

### How to Run

Windows/Linux/MacOS
```
Usage: curlRed.exe / curlRed
Example (ipv4): curlRed.exe 1.1.1.1
Example (ipv6): curlRed.exe 2a03:2880:2130:cf05:face:b00c::1.Faceb00c
Example (domain): curlRed.exe google.com
Example (batch mode): curlRed.exe 1.1.1.1 2a03:2880:2130:cf05:face:b00c::1.Faceb00c google.com
```

Build
```
go build curlRed.go
```

## Authors

* **Juan Ortega** - *Initial work* - [falseShepherd](https://github.com/false00)


