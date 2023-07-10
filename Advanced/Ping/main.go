package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"time"
)

var (
	timeout      int64
	count        int
	size         int
	sendCount    int
	successCount int
	failCount    int
	minTs        int64 = math.MaxInt32
	maxTs        int64 = 0
	totalTs      int64
)

type ICMP struct {
	Type     uint8
	Code     uint8
	CheckSum uint16
	ID       uint16
	SeqNum   uint16
}

func main() {
	getArgs()
	desIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	fmt.Printf("正在 Ping %s [%s] 具有 %d 字节的数据:\n", desIp, conn.RemoteAddr(), size)
	for i := 0; i < count; i++ {
		sendCount++
		icmp := ICMP{
			Type:     8,
			Code:     0,
			CheckSum: 0,
			ID:       uint16(i),
			SeqNum:   uint16(i),
		}
		var buffer bytes.Buffer
		binary.Write(&buffer, binary.BigEndian, icmp)
		data := make([]byte, size)
		buffer.Write(data)
		data = buffer.Bytes()
		checksum, err := checkSum(data)
		if err != nil {
			failCount++
			continue
		}
		data[2] = byte(checksum >> 8)
		data[3] = byte(checksum)
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		t1 := time.Now()
		_, err = conn.Write(data)
		if err != nil {
			failCount++
			log.Println(err)
			continue
		}
		buf := make([]byte, 1<<16)
		n, err := conn.Read(buf)
		ts := time.Since(t1).Milliseconds()
		totalTs += ts
		if minTs > ts {
			minTs = ts
		}
		if maxTs < ts {
			maxTs = ts
		}
		if err != nil {
			failCount++
			fmt.Println(err)
			continue
		}
		fmt.Printf("来自 %d.%d.%d.%d 的回复: 字节=%d 时间=%dms TTL=%d\n", buf[12], buf[13], buf[14], buf[15], n-28, ts, buf[8])
		successCount++

	}
	fmt.Printf("%s 的 Ping 统计信息:\n    数据包: 已发送 = %d，已接收 = %d，丢失 = %d (%.2f%% 丢失)，\n往返行程的估计时间(以毫秒为单位):\n    最短 = %dms，最长 = %dms，平均 = %dms",
		conn.RemoteAddr(), sendCount, successCount, failCount, float64(failCount)/float64(sendCount), minTs, maxTs, totalTs/int64(sendCount))
}

func checkSum(data []byte) (uint16, error) {
	length := len(data)
	index := 0
	var sum uint32
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	if length == 1 {
		sum += uint32(data[index])
	}
	hi16 := sum >> 16
	for hi16 != 0 {
		sum = hi16 + uint32(uint16(sum))
		hi16 = sum >> 16
	}
	return uint16(^sum), nil
}
func getArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时间")
	flag.IntVar(&count, "n", 4, "请求次数")
	flag.IntVar(&size, "l", 32, "请求次数")
	flag.Parse()
}
