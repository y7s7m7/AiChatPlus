package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// 写入文档
//func (a *App) WriteMd(filename string, Title string, Desc string, data string) {
//	fmt.Println(filename)
//	fmt.Println(data)
//	//创建文件在运行目录/chat下
//	//运行目录
//	pathname, _ := filepath.Abs(".")
//	file, err := os.OpenFile(filepath.Join(pathname, "chats", filename+".md"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 776)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//	var (
//		md     SaveMd
//		MdData []Message
//	)
//	md.Title = Title
//	md.Time, _ = strconv.Atoi(filename)
//	md.Desc = Desc
//	_ = sonic.Unmarshal([]byte(data), &MdData)
//	//遍历列表步长为2
//	var SData []SavMdData
//	for i := 0; i < len(MdData); i += 2 {
//		//格式化为mddata格式'
//		mTime, _ := strconv.Atoi(filename)
//		msg := SavMdData{
//			Time: mTime,
//			Send: MdData[i].Content,
//			Rev:  MdData[i+1].Content,
//		}
//		SData = append(SData, msg)
//	}
//	md.Data = SData
//	Sdata, _ := sonic.Marshal(md)
//	fmt.Println(string(Sdata))
//	_, _ = file.WriteString(string(Sdata))
//}

func (a *App) ChatRouter(time string) {
	runtime.EventsEmit(a.ctx, "OnLoad", time)
}
