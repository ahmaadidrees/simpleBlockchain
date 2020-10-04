package uri

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"SimpleGet",
		"GET",
		"/simpleget",
		SimpleGet,
	},
	Route{
		"AnotherGet",
		"GET",
		"/anotherget/{number}",
		AnotherGet,
	},
	Route{
		"SimplePost",
		"POST",
		"/simplepost",
		SimplePost,
	},
	Route{
		"AskOddOrEven",
		"GET",
		"/askoddoreven/{number}", // sample => askoddoreven/5
		AskOddOrEven,             //api to send post request
	},
	Route{
		"OddOrEven",
		"POST",
		"/oddoreven",
		OddOrEven,
	},
	Route{
		"Register",
		"POST",
		"/register",
		Register,
	},
	Route{
		"Show",
		"GET",
		"/show",
		ShowHandler,
	},
	Route{
		"Upload",
		"GET",
		"/upload",
		Upload,
	},
	Route{
		"Heartbeat",
		"POST",
		"/heartbeat",
		ReceiveHeartBeat,

	},
}
