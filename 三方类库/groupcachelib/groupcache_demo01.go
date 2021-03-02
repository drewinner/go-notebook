package groupcachelib

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/golang/groupcache"
	"io"
	"net/http"
	"os"
	"strings"
)

//https://my.oschina.net/u/3470972/blog/1603171
const defaultHost = "127.0.0.1:9001" //对外提供
const group_addr = ":8081"           //集群内部peer互相通信使用

func GroupCacheDemo01() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s peer1 [peer2...]", os.Args[0])
		os.Exit(1)
	}

	//本地peer地址
	self := flag.String("self", defaultHost, "self node")
	flag.Parse()

	//cache集群所有节点
	cluster := os.Args[1:]

	//初始化本地groupcache, 并监听groupcache相应的端口
	setUpGroup("test_cache")
	//本地peer
	peers := groupcache.NewHTTPPool(addrsToUrl(*self)[0])
	//设置集群信息 用以本机缓存没命中的时候，一致性哈希查找key的存储节点, 并通过http请求访问
	peers.Set(addrsToUrl(cluster...)...)

	selfPort := strings.Split(*self, ":")[1]
	_ = http.ListenAndServe(":"+selfPort, peers) //监听本机集群内部通信的端口

}

//启动groupcache
func setUpGroup(name string) {
	//缓存池,
	//这里有一个惯用方法、声明一个函数类型、此函数类型实现某个接口、并且执行该函数、
	// 其它方法参数是接口类型、调用的时候、强制转换成该函数类型、golang http类库有类似实现方式
	stringGroup := groupcache.NewGroup(name, 1<<20, groupcache.GetterFunc(func(_ context.Context, key string, dest groupcache.Sink) error {
		//当cache miss之后，用来执行的load data方法
		fp, err := os.Open("./groupcache.conf")
		if err != nil {
			return err
		}
		defer fp.Close()

		fmt.Printf("look up for %s from config_file\n", key)
		//按行读取配置文件
		buf := bufio.NewReader(fp)
		for {
			line, err := buf.ReadString('\n')
			fmt.Printf("read err:%+v", err)
			if err != nil {
				if err == io.EOF {
					dest.SetBytes([]byte{})
					return nil
				} else {
					return err
				}
			}

			line = strings.TrimSpace(line)
			parts := strings.Split(line, "=")
			if len(parts) > 2 {
				continue
			} else if parts[0] == key {
				dest.SetBytes([]byte(parts[1]))
				return nil
			} else {
				continue
			}
		}
	}))

	http.HandleFunc("/config", func(rw http.ResponseWriter, r *http.Request) {
		k := r.URL.Query().Get("key")
		var dest []byte
		fmt.Printf("look up for %s from groupcache\n", k)
		if err := stringGroup.Get(nil, k, groupcache.AllocatingByteSliceSink(&dest)); err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("this key doesn't exists"))
		} else {
			rw.Write([]byte(dest))
		}

	})

	//能够直接访问cache的端口, 启动http服务
	//http://ip:group_addr/config?key=xxx
	go http.ListenAndServe(group_addr, nil)

}

//将ip:port转换成url的格式
func addrsToUrl(node_list ...string) []string {
	urls := make([]string, len(node_list))
	for k, addr := range node_list {
		urls[k] = "http://" + addr
	}

	return urls
}
