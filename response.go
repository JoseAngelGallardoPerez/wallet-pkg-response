package response

type ErrorTitled interface {
	error
	Title() string
}

type ErrorWithDetails interface {
	error
	Details() string
}

type ErrorWithCode interface {
	error
	Code() string
}

type ErrorWithHttpStatus interface {
	error
	HttpStatus() int
}

type ErrorWithMeta interface {
	error
	Meta() interface{}
}

type Error struct {
	Title   string      `json:"title"`
	Details *string     `json:"details"`
	Code    *string     `json:"code"`
	Source  *string     `json:"source"`
	Target  string      `json:"target"`
	Meta    interface{} `json:"meta"`
}

type Response struct {
	Data     interface{} `json:"data,omitempty"`
	Messages []string    `json:"messages,omitempty"`
	Errors   []*Error    `json:"errors,omitempty"`
}

func NewResponse(params ...interface{}) *Response {
	response := new(Response)
	if len(params) > 0 {
		response.Data = params[0]
	}
	if len(params) > 1 {
		if messages, ok := params[1].([]string); ok {
			response.Messages = messages
		}
	}
	if len(params) >= 2 {
		if errors, ok := params[2].([]*Error); ok {
			response.Errors = errors
		}
	}
	return response
}

func NewResponseWithError(
	title string,
	details *string,
	code *string,
	source *string,
	target string,
) *Response {
	return NewResponse().AddError(title, details, code, source, target, nil)
}

func (r *Response) AddError(
	title string,
	details *string,
	code *string,
	source *string,
	target string,
	meta interface{},
) *Response {
	e := &Error{
		Title:   title,
		Details: details,
		Code:    code,
		Source:  source,
		Target:  target,
		Meta:    meta,
	}
	r.Errors = append(r.Errors, e)
	return r
}

func (r *Response) AddMessage(message string) *Response {
	r.Messages = append(r.Messages, message)
	return r
}
