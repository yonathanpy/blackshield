package main

import (
	"log"
	"net"
	"sync"
	"time"
)

type node struct {
	hits      uint64
	strikes   uint8
	lastSeen  time.Time
}

var (
	table   = make(map[string]*node)
	mu      sync.Mutex

	limit        uint64 = 100
	strikeLimit  uint8  = 4
)

func inspect(ip string) {
	mu.Lock()
	defer mu.Unlock()

	n, ok := table[ip]
	if !ok {
		table[ip] = &node{hits: 1, lastSeen: time.Now()}
		return
	}

	n.hits++
	n.lastSeen = time.Now()

	if n.hits > limit {
		n.strikes++
		log.Printf("violation src=%s hits=%d strikes=%d", ip, n.hits, n.strikes)

		if n.strikes >= strikeLimit {
			log.Printf("enforced src=%s action=drop", ip)
		}
	}
}

func collector() {
	for {
		time.Sleep(30 * time.Second)
		mu.Lock()
		for k, v := range table {
			if time.Since(v.lastSeen) > 180*time.Second {
				delete(table, k)
			}
		}
		mu.Unlock()
	}
}

func main() {
	go collector()

	l, err := net.Listen("tcp", ":8443")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		ip, _, _ := net.SplitHostPort(c.RemoteAddr().String())

		go func() {
			inspect(ip)
			c.Close()
		}()
	}
}
