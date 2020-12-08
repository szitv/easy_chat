package random

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

var seeds []string  //种子，即生成源，0-9a-zA-Z总共62个
var batch int // 批次，传入第一个字符串为1
var total int // 总数
var counter int // 计数器
var codeLen int // id长度
var res string // 结果统一存储，如为写入数据库，则不需要。超过10w条不建议放在变量中，应直接写入目标数据库
var completedChan = make(chan int, 10) // 判断10条协程是否都执行完毕


func GetId(begin int, num int, len int) {
	counter = 0
	codeLen = len
	batch = begin
	seeds = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	total = num
	// completedChan := make(chan int, 10) // 判断10条协程是否都执行完毕
	perBegin := begin
	perEnd := 0

	if num >= 1000000 {
		last := num % 10

		per := num / 10 //每条长度
		for c := 0; c < 10; c++ {
			// 开10条协程处理结果
			perBegin = begin + c * per
			// 取模不为0，则最后一个协程拿取模的数量+递增
			if c == 9 && last != 0 {
				perEnd = perBegin + last + per
			} else {
				perEnd = perBegin + per
			}
			//fmt.Println(perBegin, perEnd)
			go getCode(perBegin, perEnd)
		}
	} else {
		go getCode(begin, begin+num)
	}
	for {
		_, ok := <-completedChan
		if ok {
			fmt.Println("finish")
			// 生成完毕回调，此处可以调用一个api将信号传递给其他进程
			goto EndLoop
		}
	}
	
	EndLoop: 
		var file = []byte(res)
		ioutil.WriteFile("./"+strconv.Itoa(batch)+".csv", file, 0666) //写入文件(字节数组)
		close(completedChan)
}

/**
	获取code
	@param begin 开始
	@param end 结束
*/
func getCode(begin int, end int) {
	for i := begin; i < end; i++ {
		calFunc(i)
	}
	completedChan <- 1
	//fmt.Println(counter)

}

/**
	线性同余算法
	@param idx 传入的动态计算的seed
	@param len 传入的需要密钥的id长度
*/
func calFunc(idx int) {
	var str string // 结果
	for i := 0; i < codeLen; i++ {
		// xn=(axn−1+b)mod(m) 线性同余
		idx = ( 9 * idx + 7 ) % 62
		str += seeds[idx]
	}
	res += str + "\n"
	//completedChan <- counter
}
