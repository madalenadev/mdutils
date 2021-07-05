package do

// Result is the standard for func asyncTask return
type Result struct {
	Data  interface{}
	Error error
}

// ChanResult channel for the result
type ChanResult chan *Result

// AsyncTask let the layer works async
func AsyncTask(fn func(*Result)) ChanResult {
	ch := make(chan *Result)

	go func() {
		r := new(Result)

		fn(r)

		ch <- r
	}()

	return ch
}
