package main

import (
	"log"
	"github.com/alexflint/go-arg"
	"net"
	"os"
	"bufio"
	"regexp"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	rangeArr := make(map[int]*net.IPNet)
	var args struct {
		File string `arg:"required, -f, --file" help:"file to parse"`
		Ranges []string `arg:"required, -r, --ranges" help:"ranges to parse"`
		Debug bool `arg:"-d, --debug" help:"enable debug"`
	}
	arg.MustParse(&args)
	
	if args.Debug {
		log.Println("file: " + args.File)
		for _, v := range args.Ranges {
			log.Println("ranges: " + v)
		}
	}

	for i, v := range args.Ranges {
		_, mask, err := net.ParseCIDR(v)
		check(err)
		rangeArr[i] = mask
	}
	
	file, err := os.Open(args.File)
	check(err)
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`)
	for scanner.Scan() {
		matched := re.FindAllString(scanner.Text(), -1)
		for _, match := range matched {
			log.Println(match)
			for _, v := range rangeArr {
				ip := net.ParseIP(match)
				if v.Contains(ip) {
					log.Print("Alert: ", match, " | ", v, " row: ", count)
				}
			}
		}
		count++	
	}

	err = scanner.Err()
	check(err)
	
}
