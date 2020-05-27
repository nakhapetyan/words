package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"sync"
)

var dict = NewDict([]string{"test", "tes", "tos"})

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func router(ctx *fasthttp.RequestCtx) {
	if bytes.Equal(ctx.Method(), []byte("GET")) {
		findWordsAction(ctx)
		return
	} else if bytes.Equal(ctx.Method(), []byte("POST")) {
		loadDictAction(ctx)
		return
	}

	ctx.SetStatusCode(404)
}

func findWordsAction(ctx *fasthttp.RequestCtx) {
	letters := ctx.QueryArgs().Peek("letters")

	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	words := dict.findWordsByLetters(string(letters))
	wordsJson, _ := json.Marshal(words)
	b.WriteString(fmt.Sprintf("{\"accounts\":%s}", wordsJson))

	ctx.SetBody(b.Bytes())
	ctx.SetStatusCode(200)
	defer bufPool.Put(b)
}

func loadDictAction(ctx *fasthttp.RequestCtx) {
	dict.loadWords(dict.parseWords(ctx.PostBody()))
	ctx.SetStatusCode(200)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	addr := ":80"
	if len(os.Args) >= 2 {
		addr = os.Args[1]
	}
	log.Fatal(fasthttp.ListenAndServe(addr, router))
}
