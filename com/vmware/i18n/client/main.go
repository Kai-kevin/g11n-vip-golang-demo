package main

import (
	"fmt"
	"context"
	"time"
	"reflect"
	"strconv"
)

type ReflectionTest struct{
	name string `test1:"1111";test2:"2222"`
}

func main(){


	str := "\"123\""

	fmt.Println(str[1])

	fmt.Println(strconv.Unquote(str))
	testReflect := ReflectionTest{"atest"}

	s := reflect.TypeOf(testReflect)

	fmt.Println(s.Field(0).Tag.Get("test1"))


	//respData := bean.QueryTranslationByCompRespData{
	//	ProductName:"1111",
	//	Version:"1.0",
	//	Pseudo:true,
	//	Messages: map[string]string{
	//		"key1":"value1",
	//		"key2":"value2",
	//	},
	//}
	//
	//respEvent := bean.QueryTranslationByCompRespEvent{
	//	Data: respData,
	//	Signature: "sig11111",
	//}
	//
	//file,_ := json.Marshal(respEvent)
	//
	//fmt.Println(string(file))

	//util.GetTranslationByComponent()

	origin,wait := make(chan int),make(chan struct{})

	Processor(origin,wait)

	for num :=2;num < 10; num ++{
		origin <- num
	}

	close(origin)

	//<-wait

	TestContext()

	TestTimeout()

}

func TestContext(){
	ctx,cancel := context.WithCancel(context.Background())

	go func(ctx context.Context){
		for{
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(2 * time.Second)
			}
		}

	}(ctx)


	time.Sleep(10 * time.Second)
	fmt.Println("通知监控停止")
	cancel()

	time.Sleep(5*time.Second)
}

func Processor(seq chan int,wait chan struct{}){
	go func(){
		prime,ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out,wait)
		for num := range seq{
			if num %prime != 0{
				out <- num
			}
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("something is wrong")
			}
		}()

		panic("test")
		close(out)


	}()
}

func TestTimeout(){
	ctx,cancel := context.WithTimeout(context.Background(),50*time.Microsecond)

	defer cancel()

	select {
		case <- time.After(1 * time.Second):
			fmt.Println("overslept")
		case <- ctx.Done():
			fmt.Println(ctx.Err())
	}
}
