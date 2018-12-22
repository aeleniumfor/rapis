package cash_routing

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Cach struct {
	protocol string
	host     string
	client   redis.Conn
}

func Init(host, protocol string) *Cach {
	fmt.Println("Redis_Init")
	return &Cach{host: host, protocol: protocol}
}

func (c *Cach) Connect() {
	c.client, _ = redis.Dial(c.protocol, c.host)
}

func (c Cach) GetCash(host string) (res string, err error) {
	//	"web.test.fukuda.localhost"
	s, err := redis.String(c.client.Do("GET", host))
	return s, err
}

func (c *Cach) SetCach(v_host, t_host string) {
	c.client.Do("SETEX", v_host, 10, t_host)
}
