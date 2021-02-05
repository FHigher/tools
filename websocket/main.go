package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	host := &url.URL{
		Scheme: "ws",
		Host:   "111.19.254.147:1234",
		Path:   "/rpc/v0",
	}

	origin := &url.URL{
		Scheme: "http",
		Host:   "111.19.254.147",
	}

	var header = map[string][]string{}
	header["Content-Type"] = []string{"application/json"}
	header["Authorization"] = []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.Jht46zDKEWjQ2OVgnT-URPMxbuWdtbbbaDC3R_AES-U"}

	client := NewWebsocketClient(host, header, 1024, origin)

	err := client.InitConn()
	if nil != err {
		return
	}

	defer client.Close()

	var body []byte
	content := `{"jsonrpc": "2.0", "method": "Filecoin.ChainHead", "params": [], "id": 1}`
	body, err = json.Marshal(content)
	if nil != err {
		log.Printf("Marshal error: %s", err)
		return
	}
	err = client.SendMessage(body)
	if nil != err {
		log.Printf("SendMessage error: %s", err)
		return
	}

	for {
		time.Sleep(5 * time.Second)
		msg, err := client.ReadMessage()
		if nil != err {
			log.Println(err)
			continue
		}

		fmt.Println(msg)

	}
}

// Client ..
type Client struct {
	Host   *url.URL
	Header http.Header
	Ws     *websocket.Conn
	BufLen int
	Origin *url.URL
}

// NewWebsocketClient ..
func NewWebsocketClient(host *url.URL, header http.Header, len int, origin *url.URL) *Client {
	return &Client{
		Host:   host,
		Header: header,
		BufLen: len,
		Origin: origin,
	}
}

// InitConn ..
func (c *Client) InitConn() (err error) {
	var config = &websocket.Config{
		Location: c.Host,
		Header:   c.Header,
		Origin:   c.Origin,
		Version:  websocket.ProtocolVersionHybi13,
	}
	c.Ws, err = websocket.DialConfig(config)

	if err != nil {
		log.Printf("DialConfig connect error: %s", err)
		return err
	}

	return err
}

// SendMessage ..
func (c *Client) SendMessage(body []byte) (err error) {

	_, err = c.Ws.Write(body)
	if err != nil {
		log.Printf("Write error: %s", err)
		return err
	}

	return nil
}

// ReadMessage ..
func (c *Client) ReadMessage() (buf []byte, err error) {
	buf = make([]byte, c.BufLen)
	_, err = c.Ws.Read(buf)
	if nil != err {
		log.Printf("Read error: %s", err)
		return
	}

	return
}

// Close ..
func (c *Client) Close() error {
	return c.Ws.Close()
}
