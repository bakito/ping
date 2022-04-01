package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-ping/ping"
	"github.com/jackpal/gateway"
)

const (
	envPingTarget = "PING_TARGET"
	envInterval   = "PING_INTERVAL"
)

func main() {
	target := os.Getenv(envPingTarget)
	if target == "" {
		gw, err := gateway.DiscoverGateway()
		if err != nil {
			log.Fatalf("error discovering default gateway: %v", err)
		}
		target = gw.String()
	}

	pinger, err := ping.NewPinger(target)
	if err != nil {
		panic(err)
	}

	if value, ok := os.LookupEnv(envInterval); ok {
		d, err := time.ParseDuration(value)
		if err != nil {
			log.Fatalf("error parsing duration %q fom env %q", value, envInterval)
		}
		pinger.Interval = d
	}

	log.Printf("Interval: %s\n", pinger.Interval.String())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	go func() {
		for range sigs {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		log.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		log.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		log.Printf("--- %s ping statistics ---\n", stats.Addr)
		log.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		log.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	log.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		log.Fatalf("error running ping: %v", err)
	}
}
