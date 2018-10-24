package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mgutz/str"
	"github.com/sirupsen/logrus"
)

const DIG_HANDLE = " /dig?"
const TOP_HANDLE = "/topics/"
const CATE_HANDLE = "/categories/"
const TREATY_HANDLE = " HTTP/"

type cmdParams struct {
	logFilePath string
	routineNum  int
}

type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlData struct {
	data  digData
	uid   string
	unode urlNode
}

type urlNode struct {
	unType string //首页
	unRid  string //资源ID
	unUrl  string //页面url
	unTime string //访问页面时间
}

type storageBlock struct {
	counterType  string
	storageModel string
	unode        urlNode
}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)

}

func main() {
	logFilePath := flag.String("logFilePath", "/Users/joe/Documents/study/elk/dig/dig.log", "log file path")
	routineNum := flag.Int("routineNum", 5, "number goroutine")
	l := flag.String("l", "/Users/joe/go/src/joewt.com/joe/analysis/tmp/log", "runtime log file")
	flag.Parse()
	params := cmdParams{logFilePath: *logFilePath, routineNum: *routineNum}
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = logFd
		defer logFd.Close()
	}

	log.Infof("Start")
	log.Infof("Params: logFilePath=%s, routineNum=%d", params.logFilePath, params.routineNum)
	//初始化channel
	var logChannel = make(chan string, params.routineNum*6)
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)
	var storageChannel = make(chan storageBlock, params.routineNum)

	// redis 连接池
	redispool, err := pool.New("tcp", "localhost:6379", params.routineNum*3)
	if err != nil {
		log.Fatalln("Redis pool create failed")
		panic(err)
	} else {
		//在空闲的时候保证连接
		go func() {
			for {
				redispool.Cmd("PING")
				time.Sleep(3 * time.Second)
			}
		}()
	}

	go readFileLineByLine(params, logChannel)

	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel, redispool)

	go dataStorage(storageChannel, redispool)
	time.Sleep(1000 * time.Second)
}

func logConsumer(logChannel chan string, pvChannel chan urlData, uvChannel chan urlData) error {
	for logStr := range logChannel {
		data := cutLogData(logStr)

		hasher := md5.New()
		hasher.Write([]byte(data.refer + data.ua))
		uid := hex.EncodeToString(hasher.Sum(nil))

		// 一些解析的工作
		uData := urlData{data, uid, formatUrl(data.url, data.time)}

		uvChannel <- uData
		pvChannel <- uData
	}

	return nil
}

func cutLogData(logStr string) digData {
	logStr = strings.TrimSpace(logStr)
	pos1 := str.IndexOf(logStr, DIG_HANDLE, 0)
	if pos1 == -1 {
		return digData{}
	}

	pos1 += len(DIG_HANDLE)

	pos2 := str.IndexOf(logStr, TREATY_HANDLE, pos1)
	d := str.Substr(logStr, pos1, pos2-pos1)

	urlInfo, err := url.Parse("http://localhost/?" + d)
	if err != nil {
		return digData{}
	}
	data := urlInfo.Query()
	return digData{
		data.Get("time"),
		data.Get("refer"),
		data.Get("url"),
		data.Get("ua"),
	}
}

func readFileLineByLine(params cmdParams, logChannel chan string) (err error) {
	fd, err := os.Open(params.logFilePath)
	if err != nil {
		log.Warningf("ReadFileLineByLine can't open: %s", params.logFilePath)
		return err
	}
	defer fd.Close()
	bufferRead := bufio.NewReader(fd)
	count := 0
	for {
		line, err := bufferRead.ReadString('\n')

		logChannel <- line
		if err != nil {
			if err == io.EOF {
				time.Sleep(time.Second * 3)
				log.Infof("Read wait")
			} else {
				log.Warningf("ReadFileLineByLine")
			}
		}
		count++
		if count%(100*params.routineNum) == 0 {
			log.Infof("ReadFileLineByLine: %d", count)
		}
	}
	return nil
}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {
	for data := range pvChannel {
		sItem := storageBlock{"pv", "ZINCRBY", data.unode}
		storageChannel <- sItem
	}
}

func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock, redispool *pool.Pool) {
	for data := range uvChannel {
		//去重使用 HyperLoglog redis带这个
		hyperLogLogKey := "uv_hyperll_" + getTime(data.data.time, "day")
		ret, err := redispool.Cmd("PFADD", hyperLogLogKey, data.uid, "EX", 86400).Int()
		if err != nil {
			log.Warningln("uvCounter redis check hyperloglog failed: ", err)
		}
		if ret != 1 {
			continue
		}

		sItem := storageBlock{"uv", "ZINCRBY", data.unode}

		storageChannel <- sItem
	}
}

func dataStorage(storageChannel chan storageBlock, redispool *pool.Pool) {
	for block := range storageChannel {
		prefix := block.counterType + "_"
		//存储类型 Redis SortedSet
		setKeys := []string{
			prefix + "day_" + getTime(block.unode.unTime, "day"),
			prefix + "hour_" + getTime(block.unode.unTime, "hour"),
			prefix + "min_" + getTime(block.unode.unTime, "min"),
			prefix + block.unode.unType + "_day_" + getTime(block.unode.unTime, "day"),
			prefix + block.unode.unType + "_hour_" + getTime(block.unode.unTime, "hour"),
			prefix + block.unode.unType + "_min_" + getTime(block.unode.unTime, "min"),
		}

		rowId := block.unode.unRid
		for _, key := range setKeys {
			ret, err := redispool.Cmd(block.storageModel, key, 1, rowId).Int()
			if ret < 0 || err != nil {
				log.Errorln("DataStorage redis error: ", err)
			}
		}
	}
}

func formatUrl(url, t string) urlNode {
	pos1 := str.IndexOf(url, TOP_HANDLE, 0)
	if pos1 != -1 {
		pos1 += len(TOP_HANDLE)
		pos2 := str.IndexOf(url, TREATY_HANDLE, 0)
		idStr := str.Substr(url, pos1, pos2-pos1)
		return urlNode{"top", idStr, url, t}
	} else {
		pos1 := str.IndexOf(url, CATE_HANDLE, 0)
		if pos1 != -1 {
			pos1 += len(CATE_HANDLE)
			pos2 := str.IndexOf(url, TREATY_HANDLE, 0)
			idStr := str.Substr(url, pos1, pos2-pos1)
			return urlNode{"cate", idStr, url, t}
		} else {
			return urlNode{"home", "1", url, t}
		}
	}
	return urlNode{}
}

func getTime(logTime, timeType string) string {
	var item string
	switch timeType {
	case "day":
		item = "2006-01-02"
		break
	case "hour":
		item = "2006-01-02 15"
		break
	case "min":
		item = "2006-01-02 15:04"
	}

	t, _ := time.Parse(item, time.Now().Format(item))
	return strconv.FormatInt(t.Unix(), 10)
}
