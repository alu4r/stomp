package frame

import (
	"strconv"
)

const (
	AcceptVersion = "accept-version"
	Ack           = "ack"
	ContentLength = "content-length"
	ContentType   = "content-type"
	Destination   = "destination"
	HeartBeat     = "heart-beat"
	Host          = "host"
	Id            = "id"
	Login         = "login"
	Message       = "message"
	MessageId     = "message-id"
	Passcode      = "passcode"
	Receipt       = "receipt"
	ReceiptId     = "receipt-id"
	Server        = "server"
	Session       = "session"
	Subscription  = "subscription"
	Transaction   = "transaction"
	Version       = "version"
)

type headerKey struct {
	AcceptVersion string
	Ack           string
	ContentLength string
	ContentType   string
	Destination   string
	HeartBeat     string
	Host          string
	Id            string
	Login         string
	Message       string
	MessageId     string
	Passcode      string
	Receipt       string
	ReceiptId     string
	Server        string
	Session       string
	Subscription  string
	Transaction   string
	Version       string
}

var HeaderKeys = &headerKey{
	AcceptVersion: AcceptVersion,
	Ack:           Ack,
	ContentLength: ContentLength,
	ContentType:   ContentType,
	Destination:   Destination,
	HeartBeat:     HeartBeat,
	Host:          Host,
	Id:            Id,
	Login:         Login,
	Message:       Message,
	MessageId:     MessageId,
	Passcode:      Passcode,
	Receipt:       Receipt,
	ReceiptId:     ReceiptId,
	Server:        Server,
	Session:       Session,
	Subscription:  Subscription,
	Transaction:   Transaction,
	Version:       Version,
}

type Header struct {
	slice map[string]string
}

// NewHeader Create a new Header and store it in the head key/for odd numbers
// subscript key and even numbers subscript value
func NewHeader(headerEntries ...string) *Header {
	h := &Header{
		slice: make(map[string]string),
	}
	if headerEntries == nil || len(headerEntries)&1 != 0 {
		return h
	} else {
		if headerEntries != nil && len(headerEntries)%2 == 0 {
			for i := 0; i < len(headerEntries); i += 2 {
				h.slice[headerEntries[i]] = headerEntries[i+1]
			}
		}
		return h
	}
}

// Add Add a header key/value pair
func (h *Header) Add(key, value string) {
	h.slice[key] = value
}

func (h *Header) Set(key, value string) {
	h.Add(key, value)
}

func (h *Header) Get(key string) string {
	return h.slice[key]
}

// AddHeader 添加一个头信息，相同的key会被覆盖
func (h *Header) AddHeader(header *Header) {
	for k, v := range header.slice {
		h.slice[k] = v
	}
}

// AddFromArray  The array must contain an even number of data with an odd subscript key and an even subscript value
func (h *Header) AddFromArray(array []string) {
	if array != nil && len(array)&1 == 0 {
		for i := 0; i < len(array); i += 2 {
			h.slice[array[i]] = array[i+1]
		}
	}
}

// toString Gets the header string representation in the frame key:value\n contains newline characters
func (h *Header) toString() string {
	var s string
	for k, v := range h.slice {
		s += k + ColonStr + v + LfStr
	}
	return s
}

// Contains  Whether a key is included
func (h *Header) Contains(key string) (value string, ok bool) {
	value, ok = h.slice[key]
	return
}

// ContainsKey 是否包含某个header
func (h *Header) ContainsKey(key string) bool {
	_, ok := h.slice[key]
	return ok
}

func (h *Header) Del(key string) {
	delete(h.slice, key)
}

func (h *Header) Clone() *Header {
	hc := &Header{slice: make(map[string]string)}
	for k, v := range h.slice {
		hc.slice[k] = v
	}
	return hc
}

// ContentLength 获取ContentLength的值 int
func (h *Header) ContentLength() (value int, ok bool, err error) {
	text, ok := h.Contains(ContentLength)
	if !ok {
		return 0, false, nil
	}
	n, err := strconv.ParseUint(text, 10, 32)
	if err != nil {
		return 0, true, err
	}
	value = int(n)
	ok = true
	return value, ok, nil
}
