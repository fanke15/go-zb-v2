package websocket

import (
	"bytes"
	"encoding/json"
	"github.com/fanke15/go-zb-v2/constant"
	"github.com/fanke15/go-zb-v2/handler"
	. "github.com/fanke15/go-zb-v2/log"
	"github.com/fanke15/go-zb-v2/types"
	"github.com/fanke15/go-zb-v2/util"
	"github.com/gorilla/websocket"
	"time"
)

type STATUS int

const (
	pingInterval        = time.Second * 3
	MaxNoReplyPingCount = 3
	ReadTimeout         = time.Second * 10
)

var (
	ping = []byte("{\"action\": \"ping\"}")
)

const (
	INIT  STATUS = iota
	START        //tcp建立成功
	STOP         //tcp close
)

type WSClient struct {
	url string

	*websocket.Conn
	remoteIPPort string
	localIPPort  string

	status  STATUS
	msgCh   chan STATUS //因为readLoop大部分时候都是处于阻塞，所以readLoop也收不到writeLoop发的消息，故只是readLoop单向发给writeLoop的消息通道，
	writeCh chan []byte //收到待发送数据

	noReplyPingCount int           //未响应的ping数量
	lastDataTime     time.Time     //最近一条数据的时间
	connectTime      time.Time     //连接建立时间
	readTimeout      time.Duration //数据推送间隔的最长时间，越过这个时间没有收到有效时间，则关闭

	handles map[string]handler.Handler
}

func NewWSClient(url string, readTimeout time.Duration) *WSClient {
	var c WSClient
	c.url = url
	c.status = INIT
	c.msgCh = make(chan STATUS, 1)
	c.writeCh = make(chan []byte, 10)
	c.readTimeout = readTimeout
	c.handles = make(map[string]handler.Handler)

	c.connect(false)

	go c.writeLoop()
	go c.readLoop()
	return &c
}

func (c *WSClient) readLoop() {
	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if c.status == STOP { //WriteLoop已经把状态设置为STOP，故不需要给WriteLoop发消息

			} else { //ReadMessage的报错
				Log.Warn().Msg(c.url + " " + c.localIPPort + " read error: " + err.Error())

				c.status = STOP
				c.msgCh <- STOP
				c.Close()
				c.handles = make(map[string]handler.Handler)
			}

			return
		}

		c.noReplyPingCount = 0 //收到新数据,所以清零此值

		ret := false
		if messageType == websocket.BinaryMessage {
			message = util.GzipUnCompress(message)
		}

		if message != nil {

			if bytes.Contains(message, []byte(constant.PONG)) {
				//do nothing
				//Log.Info().Msg(string(message))
			} else if bytes.Contains(message, []byte(constant.LOGIN)) {
				Log.Info().Msg(string(message))
			} else {
				ret = c.handle(message)
			}

		} else {
			ret = true
		}

		if ret {
			//先通知WriteLoop，再关闭tcp
			c.status = STOP
			c.msgCh <- STOP
			c.Close()

			Log.Error().Msg("exit....")

			return
		}
	}
}

func (c *WSClient) Write(msg []byte) {
	c.writeCh <- msg
}

func (c *WSClient) writeLoop() {
	pingTimer := time.NewTimer(pingInterval)
	defer pingTimer.Stop()

	for {
		select {
		case <-pingTimer.C:
			if c.noReplyPingCount > MaxNoReplyPingCount || time.Since(c.lastDataTime) > c.readTimeout { //超时没有响应，则关闭连接
				if c.noReplyPingCount > MaxNoReplyPingCount {
					Log.Warn().Msg(c.url + " " + c.localIPPort + " error: ping timeout")
				} else {
					Log.Warn().Msg(c.url + " " + c.localIPPort + " error: data timeout")
				}

				c.status = STOP
				c.Close()

				return
			}

			pingTimer.Reset(pingInterval)

			c.noReplyPingCount++
			c.WriteTimeout(websocket.TextMessage, ping, pingInterval) //由于写的情况非常少，所以不用考虑写出错的情况。即使出错，由于有Timeout，也不会一直阻塞

		case msg := <-c.msgCh: //这是从ReadLoop过来的消息，ReadLoop本身就会设置status，故不需要再次设置
			if msg == STOP {
				return
			} else if msg == START {
				pingTimer.Reset(pingInterval)
			}
		case data := <-c.writeCh:
			//Log.Debug().Msg("write:" + string(data))
			c.WriteTimeout(websocket.TextMessage, data, time.Second)
		}
	}
}

func (c *WSClient) connect(isReconnect bool) {
	c.mustConnectWebsocket(c.url)

	c.remoteIPPort = c.RemoteAddr().String()
	c.localIPPort = c.LocalAddr().String()

	c.noReplyPingCount = 0
	c.lastDataTime = time.Now()
	c.connectTime = c.lastDataTime

	c.status = START
	if isReconnect {
		Log.Info().Msg(c.url + " " + c.remoteIPPort + " " + c.localIPPort + " ReConnect success")
	} else {
		Log.Info().Msg(c.url + " " + c.remoteIPPort + " " + c.localIPPort + " Connect success")
	}
}

func (c *WSClient) mustConnectWebsocket(url string) {
	for {
		dialer := &websocket.Dialer{}
		dialer.HandshakeTimeout = time.Second * 5

		ws, _, err := dialer.Dial(url, nil)
		if err == nil {
			c.Conn = ws
			break

		} else {
			Log.Warn().Msg("connect:" + url + " error:" + err.Error() + ". try is 1s later")
			time.Sleep(time.Second)
		}
	}
}

//true:成功
//false:失败
func (c *WSClient) TryWriteMessage(buf []byte) error {
	var err error
	retries := 3

	i := 0
	for ; i < retries; i++ {
		err = c.WriteTimeout(websocket.TextMessage, buf, pingInterval)
		if err == nil {
			break
		}
	}
	return err
}

func (c *WSClient) WriteTimeout(messageType int, data []byte, timeout time.Duration) error {
	var err error
	err = c.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		return err
	}

	err = c.WriteMessage(messageType, data)
	if err != nil {
		return err
	}

	err = c.SetWriteDeadline(time.Time{})
	if err != nil {
		return err
	}

	return nil
}

//true: 关闭tcp连接  false: 不关闭tcp连接
func (c *WSClient) handle(message []byte) bool {
	Log.Debug().Msg(string(message))

	var response types.WSResponse
	if err := json.Unmarshal(message, &response); err != nil {
		Log.Error().Msg(err.Error())
		return true
	}

	if response.Code == 0 { //没有错误
		if handler, isOK := c.handles[response.Channel]; isOK {
			return handler(*response.Data)

		} else {
			Log.Error().Str("data", string(message)).Msg("channel Don't support")
			return true
		}
	} else {
		Log.Error().Msg(string(message))
		return false
	}

	return true
}

func (c *WSClient) AddHandler(channel string, handler handler.Handler) {
	c.handles[channel] = handler
}

func (c *WSClient) DeleteHandler(channel string) {
	if _, isOK := c.handles[channel]; isOK {
		delete(c.handles, channel)
	}
}
