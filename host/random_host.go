package host

import (
	"encoding/json"
	"github.com/go-playground/lars"
	"io/ioutil"
	"net/http"
	"strconv"
)

const MaxStackSize = 1000

type RandomNumberReceiver struct {
	numbers *Stack
	key string
}

func NewRandomNumberReceiver() *RandomNumberReceiver {
	return &RandomNumberReceiver{NewStack(), ""}
}

func (r *RandomNumberReceiver) ListenAndServer(addr, key string) {
	r.key = key
	router := lars.New()
	router.Post("/" + key, r.postNumbers)
	router.Get("/numbers", r.getNumbers)
	router.Get("/random", r.getRandom)
	router.Get("/count", r.getNumbersCount)
	http.ListenAndServe(addr, router.Serve())
}

func (r *RandomNumberReceiver) GetRandomNumber() (int, error) {
	return r.numbers.Pop()
}

func (r *RandomNumberReceiver) GetStack() *Stack {
	return r.numbers
}

func (r *RandomNumberReceiver) GetKey() string {
	return r.key
}

func (r *RandomNumberReceiver) postNumbers(c lars.Context) {
	defer c.Request().Body.Close()
	data, err := ioutil.ReadAll(c.Request().Body)
	dataLen := len(data)
	if err != nil || dataLen < 3 {
		c.Response().WriteHeader(http.StatusBadRequest)
		c.Response().Write([]byte("invalid body"))
		return
	}
	dataStr := string(data)
	// remove redundant quotation marks if exists
	if dataStr[0] == '"' && dataStr[dataLen-1] == '"' {
		data = []byte(dataStr[1:dataLen-1])
	}
	var keys []int
	// unmarshal the json into the temporary holder
	err = json.Unmarshal(data, &keys)
	if err != nil {
		c.Response().WriteHeader(http.StatusUnprocessableEntity)
		c.Response().Write([]byte("invalid json"))
		return
	}
	if len(keys) + r.numbers.Size() > MaxStackSize {
		r.numbers.Clear()
	}
	// write new numbers to the cache
	r.numbers.PushMany(keys)
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Write([]byte("ok"))
}

func (r *RandomNumberReceiver) getNumbers(c lars.Context) {
	data, err := r.numbers.Json()
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().WriteString("internal error: " + err.Error())
		return
	}
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Write(data)
}

func (r *RandomNumberReceiver) getNumbersCount(c lars.Context) {
	c.Response().WriteHeader(http.StatusOK)
	c.Response().WriteString(strconv.Itoa(r.numbers.Size()))
}

func (r *RandomNumberReceiver) getRandom(c lars.Context) {
	if r.numbers.Size() < 1 {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().WriteString("internal error: not enough unique numbers provided by the source")
		return
	}
	n, err := r.numbers.Pop()
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().WriteString("internal error: not enough unique numbers provided by the source or " + err.Error())
		return
	}
	c.Response().WriteHeader(http.StatusOK)
	c.Response().WriteString(strconv.Itoa(n))
}