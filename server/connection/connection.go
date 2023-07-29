package connection

import (
	"io"
	"net"
	"os"
	"strconv"
	"time"

	"ehang.io/nps/lib/conn"
	"ehang.io/nps/lib/crypt"
	"ehang.io/nps/lib/pmux"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var pMux *pmux.PortMux
var bridgePort string
var httpsPort string
var httpPort string
var webPort string

func InitConnectionService() {
	bridgePort = beego.AppConfig.String("bridge_port")
	httpsPort = beego.AppConfig.String("https_proxy_port")
	httpPort = beego.AppConfig.String("http_proxy_port")
	webPort = beego.AppConfig.String("web_port")

	if httpPort == bridgePort || httpsPort == bridgePort || webPort == bridgePort {
		port, err := strconv.Atoi(bridgePort)
		if err != nil {
			logs.Error(err)
			os.Exit(0)
		}
		pMux = pmux.NewPortMux(port, beego.AppConfig.String("web_host"))
	}
}

func GetBridgeListener(tp string) (net.Listener, error) {
	logs.Info("server start, the bridge type is %s, the bridge port is %s", tp, bridgePort)
	var p int
	var err error
	if p, err = strconv.Atoi(bridgePort); err != nil {
		return nil, err
	}
	if pMux != nil {
		return pMux.GetClientListener(), nil
	}
	return net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(beego.AppConfig.String("bridge_ip")), p, ""})
}

func GetHttpListener() (net.Listener, error) {
	if pMux != nil && httpPort == bridgePort {
		logs.Info("start http listener, port is", bridgePort)
		return pMux.GetHttpListener(), nil
	}
	logs.Info("start http listener, port is", httpPort)
	return getTcpListener(beego.AppConfig.String("http_proxy_ip"), httpPort)
}

func GetHttpsListener() (net.Listener, error) {
	if pMux != nil && httpsPort == bridgePort {
		logs.Info("start https listener, port is", bridgePort)
		return pMux.GetHttpsListener(), nil
	}
	logs.Info("start https listener, port is", httpsPort)
	return getTcpListener(beego.AppConfig.String("http_proxy_ip"), httpsPort)
}

func GetWebManagerListener() (net.Listener, error) {
	if pMux != nil && webPort == bridgePort {
		logs.Info("Web management start, access port is", bridgePort)
		return pMux.GetManagerListener(), nil
	}
	logs.Info("web management start, access port is", webPort)
	return getTcpListener(beego.AppConfig.String("web_ip"), webPort)
}

func getTcpListener(ip, p string) (net.Listener, error) {
	port, err := strconv.Atoi(p)
	if err != nil {
		logs.Error(err)
		os.Exit(0)
	}
	if ip == "" {
		ip = "0.0.0.0"
	}
	return net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
}

/*
主要做config的处理
*/
type ConfigModeServer struct {
	listener net.Listener
}

func NewConfigServer() *ConfigModeServer {
	return &ConfigModeServer{}
}
func Gettime() (time_Now1 string) {
	time_Now := time.Now().Unix()
	s := strconv.FormatInt(time_Now, 10)
	return s
}
func Handler(c net.Conn) {
	var salt = "学习新思想"
	/*
		思路，客户端发送时间戳+特定字符 进行aes加密，服务端收到进行解密，如果后几个字符不同则解密错误
		保证了传输过程中动态, 如果发生错误，直接给客户端返回 salt+时间戳
	*/
	conn1 := conn.NewConn(c)
	defer conn1.Close()
	b2, err11 := conn1.GetShortLenContent()
	// fmt.Printf("string(b2): %v\n", string(b2))
	if err11 != nil {
		conn1.WriteLenContent([]byte(crypt.Md5(Gettime() + salt)))
		logs.Error("Get Short Content Error", err11.Error())
		return
	}
	b3, err12 := crypt.AesDecrypt(b2, []byte("学习新思想争做新青年!!"))
	if err12 != nil {
		conn1.WriteLenContent([]byte(crypt.Md5(Gettime() + salt)))
		logs.Error("AesDecode Error", err12.Error())
		return
	}
	// fmt.Printf("b3: %v\n", b3)
	if string(b3[0:6]) != "@#@113" {
		conn1.WriteLenContent([]byte(crypt.Md5(Gettime() + salt)))
		logs.Error("Client Passwd Error")
		return
	}
	f, err := os.Open("./conf/npc.conf")
	if err != nil {
		conn1.WriteLenContent([]byte(crypt.Md5(Gettime() + salt)))
		logs.Error("open File ./conf/npc.conf fail", err.Error())
		return
	}
	b, err1 := io.ReadAll(f)
	if err1 != nil {
		conn1.WriteLenContent([]byte(crypt.Md5(Gettime() + salt)))
		logs.Error("read file fail", err1.Error())
		return
	}
	b4, err2 := crypt.AesEncrypt(b, []byte("好好学习天天向上冲鸭!!"))
	if err2 != nil {
		conn1.WriteLenContent([]byte(crypt.Md5(Gettime() + salt)))
		logs.Error("AesEncode Server Send fail", err2.Error())
		return
	}
	err3 := conn1.WriteLenContent(b4)
	if err3 != nil {
		conn1.WriteLenContent([]byte(crypt.Md5(Gettime() + salt)))
		logs.Error("Write Server Send fail", err2.Error())
		return
	}
}
func (CServer *ConfigModeServer) ConfigHostHandler(Host string) {
	conn.NewTcpListenerAndProcess(Host, Handler, &CServer.listener)
}
