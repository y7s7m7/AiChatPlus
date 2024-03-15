package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) GetAllFile(pathname string) [][]SaveMd {
	var allFiles [][]SaveMd
	var today []SaveMd
	var yesterday []SaveMd
	var lastSevenDays []SaveMd
	var lastThirtyDays []SaveMd

	// 获取运行目录
	pathname, _ = filepath.Abs(pathname)

	// 判断目录是否存在，不存在则创建
	if _, err := os.Stat(pathname); os.IsNotExist(err) {
		if err := os.Mkdir(pathname, os.ModePerm); err != nil {
			log.Println("Error creating directory:", err)
			return allFiles
		}
	}

	// 获取目录下所有md文件
	err := filepath.Walk(pathname, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error accessing file:", err)
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		if filepath.Ext(filename) == ".md" {
			// 读取并解析md文件
			file, err := os.ReadFile(filename)
			if err != nil {
				log.Println("Error reading file:", err)
				return nil
			}
			var ReadMd SaveMd
			if err := sonic.Unmarshal(file, &ReadMd); err != nil {
				log.Println("Error unmarshalling file:", err)
				return nil
			}

			// 获取文件名中的时间作为Md结构体的Time字段
			filename = filepath.Base(filename)
			filename = strings.TrimSuffix(filename, filepath.Ext(filename))
			Time, _ := strconv.Atoi(filename)
			fileTime := time.Unix(int64(Time), 0)

			// 创建Md结构体并添加到相应日期分组中
			md := SaveMd{ReadMd.Title, Time, ReadMd.Desc, ReadMd.Data}
			diff := time.Since(fileTime)
			if diff < 24*time.Hour {
				today = append(today, md)
			} else if diff < 48*time.Hour {
				yesterday = append(yesterday, md)
			} else if diff < 7*24*time.Hour {
				lastSevenDays = append(lastSevenDays, md)
			} else if diff < 30*24*time.Hour {
				lastThirtyDays = append(lastThirtyDays, md)
			} else {
				// 删除超过三十天的文件
				if err := os.Remove(filename); err != nil {
					log.Println("Error deleting file:", err)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Error walking through directory:", err)
		return allFiles
	}

	// 根据Time字段对files进行排序
	sort.Slice(today, func(i, j int) bool {
		return today[i].Time < today[j].Time
	})
	sort.Slice(yesterday, func(i, j int) bool {
		return yesterday[i].Time < yesterday[j].Time
	})
	sort.Slice(lastSevenDays, func(i, j int) bool {
		return lastSevenDays[i].Time < lastSevenDays[j].Time
	})
	sort.Slice(lastThirtyDays, func(i, j int) bool {
		return lastThirtyDays[i].Time < lastThirtyDays[j].Time
	})

	// 添加四个列表到allFiles中
	allFiles = append(allFiles, today)
	allFiles = append(allFiles, yesterday)
	allFiles = append(allFiles, lastSevenDays)
	allFiles = append(allFiles, lastThirtyDays)

	return allFiles
}

func (a *App) CreateMd(folderName string, fileName string, timestamp int, desc string) bool {
	// 创建 Md 结构体
	md := SaveMd{
		Title: fileName,
		Time:  timestamp,
		Desc:  desc,
		Data:  nil,
	}

	// 将 Md 结构体转换为 JSON 格式
	snc := sonic.Config{
		CompactMarshaler: true,
		SortMapKeys:      true,
	}.Froze()
	data, err := snc.Marshal(&md)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// 确定保存路径
	pathname, err := filepath.Abs(folderName)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// 创建文件夹
	if _, err := os.Stat(pathname); os.IsNotExist(err) {
		if err := os.MkdirAll(pathname, os.ModePerm); err != nil {
			log.Fatal(err)
			return false
		}
	}

	// 创建并写入文件
	filename := fmt.Sprintf("%d.md", timestamp)
	filePath := filepath.Join(pathname, filename)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0776)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (a *App) GetMd(Path string, Time int) {
	pathname, _ := filepath.Abs(".")
	filename := fmt.Sprintf("%d.md", Time)
	filePath := filepath.Join(pathname, Path, filename)
	content, err := os.ReadFile(filePath)

	var data SaveMd
	if err != nil {
		log.Println(err)
		data.Title = "New Chat"
		data.Time = Time
		err = sonic.Unmarshal(content, &data)
		runtime.EventsEmit(a.ctx, "md_get", data)
		return
	}
	err = sonic.Unmarshal(content, &data)
	runtime.EventsEmit(a.ctx, "md_get", data)
}

func (a *App) PostOllma(model string, prompt string) Message {
	var input []Message
	fmt.Println(prompt)
	_ = json.Unmarshal([]byte(prompt), &input)
	messageData := ModelData{
		Model:    model,
		Messages: input,
	}
	data, _ := json.Marshal(messageData)
	fmt.Println(input)
	fmt.Println(string(data))
	//post请求http://localhost:11434/api/generate
	res, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	runtime.EventsEmit(a.ctx, "CreateMd")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	var buffer [1024]byte
	var out = ""
	for {
		//接受数据
		n, err := res.Body.Read(buffer[:])
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			break
		}

		//处理数据
		t := string(buffer[:n])
		fmt.Println(t)
		var JsonData Data
		_ = sonic.Unmarshal([]byte(t), &JsonData)
		if JsonData.Done {
			fmt.Println(out)
		} else {
			out += JsonData.Message.Content
			var out2 = Message{
				Role:    "assistant",
				Content: out,
			}
			if err != nil {
				return Message{}
			}
			runtime.EventsEmit(a.ctx, "onMessage", out2)
		}
	}
	//创建一个map make
	var out2 = Message{
		Role:    "assistant",
		Content: out,
	}
	return out2
}
