package conf

func initProdConf() {
	LogicConf = logicConf{
		MySQL:                  "root:liu123456@tcp(localhost:3306)/gim?charset=utf8&parseTime=true",
		NSQIP:                  "127.0.0.1:4150",
		RedisIP:                "127.0.0.1:6379",
		RPCIntListenAddr:       ":50000",
		ClientRPCExtListenAddr: ":50001",
		ServerRPCExtListenAddr: ":50002",
		ConnRPCAddrs:           "addrs:///127.0.0.1:60000",
	}

	ConnConf = connConf{
		TCPListenAddr: ":8080",
		RPCListenAddr: ":60000",
		LocalAddr:     "127.0.0.1:60000",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	WSConf = wsConf{
		WSListenAddr:  ":8080",
		RPCListenAddr: ":60000",
		LocalAddr:     "127.0.0.1:60000",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}
}
