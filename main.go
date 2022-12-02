package main

import (
	"fmt"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/afpacket"
	"github.com/google/gopacket/layers"
)

func main() {
	fmt.Println("vim-go")

	var wg sync.WaitGroup

	frameSize := afpacket.OptFrameSize(1 << 11)
	blockSize := afpacket.OptBlockSize(frameSize << 11)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			tpkt, err := afpacket.NewTPacket(
				afpacket.OptInterface("wlp4s0"),
				frameSize,
				blockSize,
				afpacket.OptNumBlocks(64),
			)
			if err != nil {
				panic(err)
			}

			if err := tpkt.SetFanout(afpacket.FanoutHashWithDefrag, 1984); err != nil {
				panic(err)
			}

			var eth layers.Ethernet
			var ipv4 layers.IPv4
			var ipv6 layers.IPv6
			var tcp layers.TCP
			var udp layers.UDP

			parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ipv4, &ipv6, &tcp, &udp)
			parser.IgnoreUnsupported = true
			decoded := make([]gopacket.LayerType, 10)

			for {
				pkt, ci, err := tpkt.ZeroCopyReadPacketData()
				if err != nil {
					panic(err)
				}

				if err := parser.DecodeLayers(pkt, &decoded); err != nil {
					fmt.Println("error decoding layers:", err)
					continue
				}

				fmt.Println(i, ci)

				for _, ltyp := range decoded {
					switch ltyp {
					case layers.LayerTypeIPv4:
						fmt.Println("-- IPv4:", ipv4.SrcIP, ipv4.DstIP)
					case layers.LayerTypeIPv6:
						fmt.Println("-- IPv6:", ipv6.SrcIP, ipv6.DstIP)
					case layers.LayerTypeTCP:
						fmt.Println("-- TCP:", tcp.SrcPort, tcp.DstPort)
					case layers.LayerTypeUDP:
						fmt.Println("-- UDP:", udp.SrcPort, udp.DstPort)
					}
				}

			}
		}(i)
	}

	wg.Wait()
}
