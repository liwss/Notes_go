package Log

import "log"

/*init()函数会在main函数之前运行，一般做一些初始化工作*/
func init(){
	log.SetPrefix("TRACE:")	//设置一个字符串，作为每个日志的前缀
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用 Println()之后会接着调用 os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln 在调用 Println(之后会接着调用) panic()
	log.Panicln("panic message")
}