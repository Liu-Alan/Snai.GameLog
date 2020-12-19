package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("2020-12-17.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	leopenid := make(map[string]int)
	enopenid := make(map[string]int)

	s := bufio.NewScanner(f)
	for s.Scan() {
		sc := s.Text()

		if strings.Contains(sc, "用户上报") && strings.Contains(sc, "\"appid\":\"326xqzwg\"") && strings.Contains(sc, "\"eventtype\":\"levelup\"") {
			if strings.Contains(sc, "\"openid\":\"") {
				idindex := strings.Index(sc, "\"openid\":\"") + 10
				openid := sc[idindex : idindex+12]
				if _, ok := leopenid[openid]; !ok {
					leopenid[openid] = 1
				}
			}
		}

		if strings.Contains(sc, "用户上报") && strings.Contains(sc, "\"appid\":\"326xqzwg\"") && strings.Contains(sc, "\"eventtype\":\"entersvr\"") {
			if strings.Contains(sc, "\"openid\":\"") {
				idindex := strings.Index(sc, "\"openid\":\"") + 10
				openid := sc[idindex : idindex+12]
				if _, ok := enopenid[openid]; !ok {
					enopenid[openid] = 1
				}
			}
		}

	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("levelup：%d\n", len(leopenid))
	fmt.Printf("entersvr：%d\n", len(enopenid))
}
