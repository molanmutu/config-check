package utils

import (
	"config-check/settings"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type Cli struct {
	Port     int    // 端口号
	IP       string // IP地址
	Username string // 用户名
	Password string // 密码
	//Auth       string      // password or key,现在只做了password和rsa公钥
	//PublicKey  string      // 密钥路径
	Des        string      // 源路径
	Src        string      // 目标路径
	LastResult string      // 最近一次Run的结果
	client     *ssh.Client // ssh客户端

}

// New 创建命令行对象
//@param ip IP地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func New(ip string) *Cli {
	cfg := *settings.Conf.BaseInfo
	cli := &Cli{
		IP:       ip,
		Username: cfg.User,
		Password: cfg.Passwd,
		Des:      cfg.DesPath,
		Src:      cfg.SrcPath,
	}
	if cfg.Port <= 0 {
		cli.Port = 22
	} else {
		cli.Port = cfg.Port
	}
	return cli
}

// Dail 连接
func (c *Cli) Dail() (*ssh.Client, error) {
	Conf := ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	sshClient, err := ssh.Dial("tcp", addr, &Conf)
	if err != nil {
		return nil, err
	}
	c.client = sshClient
	return ssh.Dial("tcp", addr, &Conf)
}

// Run 执行shell
//@param shell shell脚本命令
func (c Cli) Run(shell string) (string, error) {
	if c.client == nil {
		if _, err := c.Dail(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}
