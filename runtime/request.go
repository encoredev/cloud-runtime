package runtime

import (
	"fmt"
	"sync/atomic"
	"time"

	"encore.dev/runtime/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"

	// These imports are used only by the generated wrappers in the compiler,
	// but add them here so the 'go' command doesn't remove them from go.mod.
	_ "github.com/felixge/httpsnoop"
)

var (
	reqIDCtr  uint32
	callIDCtr uint64
)

var (
	RootLogger *zerolog.Logger
	Config     *config.ServerConfig
)

var json = jsoniter.Config{
	EscapeHTML:             false,
	SortMapKeys:            false,
	ValidateJsonRawMessage: true,
}.Froze()

type UID string

func BeginOperation() {
	encoreBeginOp(true /* always trace */)
}

func FinishOperation() {
	encoreFinishOp()
}

type SpanID [8]byte

type Type byte

const (
	RPCCall     Type = 0x01
	AuthHandler Type = 0x02
)

type Request struct {
	Type     Type
	SpanID   SpanID
	ParentID SpanID
	UID      UID
	AuthData interface{}

	Service  string
	Endpoint string
	Start    time.Time
	Logger   zerolog.Logger
	Traced   bool
}

type RequestData struct {
	Type            Type
	Service         string
	Endpoint        string
	CallExprIdx     int32
	EndpointExprIdx int32
	Inputs          [][]byte
	UID             UID
	AuthData        interface{}
}

func BeginRequest(data RequestData) error {
	spanID, err := genSpanID()
	if err != nil {
		return err
	}
	return beginReq(spanID, data)
}

func FinishRequest(status int, outputs [][]byte, err error) {
	finishReq(status, outputs, err)
}

type Call struct {
	CallID uint64
	SpanID SpanID
}

type CallParams struct {
	Service         string
	Endpoint        string
	CallExprIdx     int32
	EndpointExprIdx int32
}

func BeginCall(params CallParams) (*Call, error) {
	spanID, err := genSpanID()
	if err != nil {
		return nil, err
	}

	callID := atomic.AddUint64(&callIDCtr, 1)

	if g := encoreGetG(); g != nil && g.req != nil && g.req.data.Traced {
		tb := NewTraceBuf(8 + 4 + 4 + 4)
		tb.UVarint(callID)
		tb.Bytes(g.req.data.SpanID[:])
		tb.Bytes(spanID[:])
		tb.UVarint(uint64(g.goid))
		tb.UVarint(uint64(params.CallExprIdx))
		tb.UVarint(uint64(params.EndpointExprIdx))
		encoreTraceEvent(CallStart, tb.Buf())
	}

	return &Call{
		CallID: callID,
		SpanID: spanID,
	}, nil
}

func (c *Call) Finish(err error) {
	if g := encoreGetG(); g != nil && g.req != nil && g.req.data.Traced {
		tb := NewTraceBuf(8 + 4 + 4 + 4)
		tb.UVarint(c.CallID)
		if err != nil {
			msg := err.Error()
			if msg == "" {
				msg = "unknown error"
			}
			tb.String(msg)
		} else {
			tb.String("")
		}
		encoreTraceEvent(CallEnd, tb.Buf())
	}
}

func (c *Call) BeginReq(data RequestData) error {
	return beginReq(c.SpanID, data)
}

func (c *Call) FinishReq(status int, outputs [][]byte, err error) {
	finishReq(status, outputs, err)
}

type AuthCall struct {
	SpanID SpanID
	CallID uint64
}

func BeginAuth(authHandlerExprIdx int32, token string) (*AuthCall, error) {
	spanID, err := genSpanID()
	if err != nil {
		return nil, fmt.Errorf("could not generate request id: %v", err)
	}
	callID := atomic.AddUint64(&callIDCtr, 1)

	if g := encoreGetG(); g != nil && g.op.trace != nil {
		tb := NewTraceBuf(8 + 4 + 4 + 4)
		tb.UVarint(callID)
		tb.Bytes(spanID[:])
		tb.UVarint(uint64(g.goid))
		tb.UVarint(uint64(authHandlerExprIdx))
		encoreTraceEvent(AuthStart, tb.Buf())
	}

	return &AuthCall{
		SpanID: spanID,
		CallID: callID,
	}, nil
}

func (ac *AuthCall) Finish(uid UID, err error) {
	if g := encoreGetG(); g != nil && g.op.trace != nil {
		tb := NewTraceBuf(64)
		tb.UVarint(ac.CallID)
		tb.String(string(uid))
		if err != nil {
			msg := err.Error()
			if msg == "" {
				msg = "unknown error"
			}
			tb.String(msg)
		} else {
			tb.String("")
		}
		encoreTraceEvent(AuthEnd, tb.Buf())
	}
}

func (ac *AuthCall) BeginReq(data RequestData) error {
	return beginReq(ac.SpanID, data)
}

func (ac *AuthCall) FinishReq(status int, outputs [][]byte, err error) {
	finishReq(status, outputs, err)
}

func Logger() *zerolog.Logger {
	if req, _, ok := CurrentRequest(); ok {
		return &req.Logger
	}
	return RootLogger
}

func CurrentRequest() (*Request, uint32, bool) {
	return currentReq()
}

func currentReq() (*Request, uint32, bool) {
	if g := encoreGetG(); g != nil && g.req != nil {
		return g.req.data, g.goid, true
	}
	return nil, 0, false
}

func TraceLog(event TraceEvent, data []byte) {
	encoreTraceEvent(event, data)
}

func SerializeInputs(inputs ...interface{}) ([][]byte, error) {
	var res [][]byte
	for _, input := range inputs {
		data, err := json.Marshal(input)
		if err != nil {
			return nil, fmt.Errorf("could not serialize input %v: %v", input, err)
		}
		res = append(res, data)
	}
	return res, nil
}

func CopyInputs(inputs [][]byte, outputs []interface{}) error {
	if len(inputs) != len(outputs) {
		panic(fmt.Sprintf("encore.dev/runtime.CopyInputs: len(inputs) != len(outputs): %v != %v",
			len(inputs), len(outputs)))
	}
	for i, data := range inputs {
		if err := json.Unmarshal(data, outputs[i]); err != nil {
			return fmt.Errorf("could not serialize input #%d: %v", i, err)
		}
	}
	return nil
}

func beginReq(spanID SpanID, data RequestData) error {
	req := &Request{
		Type:     data.Type,
		SpanID:   spanID,
		Service:  data.Service,
		Endpoint: data.Endpoint,
		Start:    time.Now(),
		UID:      data.UID,
		AuthData: data.AuthData,
	}

	if prev, _, ok := currentReq(); ok {
		req.UID = prev.UID
		req.AuthData = prev.AuthData
		req.ParentID = prev.SpanID
		encoreClearReq()
	}

	encoreBeginReq(spanID, req, true /* always trace */)

	ctx := RootLogger.With().
		Str("service", req.Service).
		Str("endpoint", req.Endpoint)
	if req.UID != "" {
		ctx = ctx.Str("uid", string(req.UID))
	}
	req.Logger = ctx.Logger()

	g := encoreGetG()
	req.Traced = g.op.trace != nil
	if req.Traced {
		tb := NewTraceBuf(1 + 8 + 8 + 8 + 8 + 8 + 8 + 64)
		tb.Bytes([]byte{byte(req.Type)})
		tb.Bytes(req.SpanID[:])
		tb.Bytes(req.ParentID[:])
		tb.UVarint(uint64(g.goid))
		tb.UVarint(uint64(data.CallExprIdx))
		tb.UVarint(uint64(data.EndpointExprIdx))
		tb.String(string(req.UID))
		tb.UVarint(uint64(len(data.Inputs)))
		for _, input := range data.Inputs {
			tb.UVarint(uint64(len(input)))
			tb.Bytes(input)
		}
		encoreTraceEvent(RequestStart, tb.Buf())
	}

	switch data.Type {
	case AuthHandler:
		req.Logger.Info().Msg("running auth handler")
	default:
		req.Logger.Info().Msg("starting request")
	}
	return nil
}

func finishReq(status int, outputs [][]byte, err error) {
	g := encoreGetG()
	if g == nil || g.req == nil {
		panic("encore: no current request running")
	}

	req := g.req.data
	if err != nil {
		switch req.Type {
		case AuthHandler:
			req.Logger.Error().Err(err).Msg("auth handler failed")
		default:
			req.Logger.Error().Err(err).Msg("request failed")
		}
	}

	if req.Traced {
		tb := NewTraceBuf(64)
		tb.Bytes(req.SpanID[:])
		if err == nil {
			tb.Bytes([]byte{0}) // no error
			tb.UVarint(uint64(len(outputs)))
			for _, output := range outputs {
				tb.UVarint(uint64(len(output)))
				tb.Bytes(output)
			}
		} else {
			tb.Bytes([]byte{1})
			tb.String(err.Error())
		}
		encoreTraceEvent(RequestEnd, tb.Buf())
	}

	dur := time.Since(req.Start)
	switch req.Type {
	case AuthHandler:
		req.Logger.Info().Dur("duration", dur).Msg("auth handler completed")
	default:
		req.Logger.Info().Dur("duration", dur).Int("status", status).Msg("request completed")
	}
	encoreCompleteReq()
}