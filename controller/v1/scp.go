package v1

import (
	"config-check/utils"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ScpRPM 判断是否需要拷贝
func ScpRPM(c *utils.Cli) {
	var _ error
	srcPath := c.Src
	desPath := c.Des
	Client, _ := c.Dail()
	//遍历并复制
	_ = filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			newPath := strings.Replace(path, srcPath, desPath, len(srcPath))
			dName := strings.Replace(newPath, "\\", "/", -1)
			mkdir(Client, dName)
			return nil
		}
		//不是路径的，就复制文件
		File, err := os.Open(path)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			os.Exit(1)
		}
		info, _ := File.Stat()
		defer File.Close()
		newPath := strings.Replace(path, srcPath, desPath, len(srcPath))
		dName := strings.Replace(newPath, "\\", "/", -1)
		fmt.Printf("dname %s\n", dName)
		scp(Client, File, info.Size(), dName)
		return nil
	})
}

// scp拷贝文件
func scp(client *ssh.Client, File io.Reader, size int64, path string) {
	filename := filepath.Base(path)
	dirname := strings.Replace(filepath.Dir(path), "\\", "/", -1)
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("创建Session失败:", err)
		return
	}
	go func() {
		w, _ := session.StdinPipe()
		fmt.Fprintln(w, "C0644", size, filename)
		io.CopyN(w, File, size)
		fmt.Fprint(w, "\x00")
		w.Close()
	}()
	if err := session.Run(fmt.Sprintf("/usr/bin/scp -qrt %s/", dirname)); err != nil {
		fmt.Println("执行scp命令失败:", err)
		if err != nil {
			session.Close()
			return
		}
	} else {
		fmt.Printf("%s 发送成功.\n", path)
		session.Close()
	}
	if session, err = client.NewSession(); err == nil {
		defer session.Close()
		buf, err := session.Output(fmt.Sprintf("/usr/bin/md5sum %s", path))
		if err != nil {
			fmt.Println("检查md5失败:", err)
			return
		}
		fmt.Printf("MD5:\n%s\n", string(buf))
	}
}

// Mkdir 创建目录
func mkdir(client *ssh.Client, path string) {
	fmt.Printf("create path %s\n", path)
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("创建Session失败:", err)
		return
	}
	session.Run(fmt.Sprintf("[ ! -d %s ] && mkdir %s", path, path))
	defer session.Close()
}
