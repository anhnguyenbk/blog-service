package main

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc errorHandler
}

type Routes []Route

var routes = Routes{
	Route{
		"PostIndex",
		"GET",
		"/posts",
		PostIndex,
	},
	Route{
		"PostCreate",
		"POST",
		"/posts",
		PostSave,
	},
	Route{
		"PostUpdate",
		"POST",
		"/posts/{postId}",
		PostSave,
	},
	Route{
		"PostShow",
		"GET",
		"/posts/{postId}",
		PostShow,
	},
}
