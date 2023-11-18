package log

import (
    "fmt"
    "sync"
    "testing"
)

type aa struct {
    Name string
    Age chan string
}

func TestLogger(t *testing.T) {
    t.Log("sfsf")

    var a = make(chan int)
    var b = aa{Name: "lfffffffffffffffff", Age: make(chan string, 1)}
    logger, err := New(
        TimestampFormat("2006-01-02 15:04:05.000000"),
        FileName("/Users/lifangfang/wwwroot/gowork/src/testtrace222/testtrace.log"),
    )
    fmt.Println(err)

    var wg sync.WaitGroup
    for i:=0;i<1000;i++ {
        wg.Add(1)
        go func(ii int) {
            defer wg.Done()
            fmt.Println("循环：", ii)

            logger.Info(Fields{
                "channel_a": a,
                "struct_b": b,
                "baol": false,
                "byte": []byte("默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数。后续调用相关方法日志将写到io.Writer中。现在，我们就能像上篇文章介绍log时一样，可以搞点事情了。传入一个io.MultiWriter，同时将日志写到bytes.Buffer、标准输出和文件中"),
                "rune": rune(5),
                "function": func() {},
                "msg": "默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数。后续调用相关方法日志将写到io.Writer中。现在，我们就能像上篇文章介绍log时一样，可以搞点事情了。传入一个io.MultiWriter，同时将日志写到bytes.Buffer、标准输出和文件中",
            }, "22", "8989")
        }(i)
    }

    wg.Wait()
    fmt.Println("DONE")
}