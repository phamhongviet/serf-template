package main

import (
	"fmt"
	"io/ioutil"
)

func ExampleRenderTemplate() {
	members := []Member{
		{
			Name:   "web-1",
			Addr:   "172.16.0.21",
			Port:   7946,
			Status: "alive",
			Tags: map[string]string{
				"app":     "nginx",
				"webport": "8080",
			},
		},
		{
			Name:   "web-2",
			Addr:   "172.16.0.22",
			Port:   7946,
			Status: "alive",
			Tags: map[string]string{
				"app":     "nginx",
				"webport": "8080",
			},
		},
		{
			Name:   "web-3",
			Addr:   "172.16.0.23",
			Port:   7946,
			Status: "alive",
			Tags: map[string]string{
				"app":     "httpd",
				"webport": "80",
			},
		},
	}
	RenderTemplate("test/template_1.tpl", "test/result_1.txt", members)
	out, err := ioutil.ReadFile("test/result_1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
/*
Output:
BEGIN
server web-1 at 172.16.0.21 with serf at 7946 and nginx at 8080
server web-2 at 172.16.0.22 with serf at 7946 and nginx at 8080
server web-3 at 172.16.0.23 with serf at 7946 and httpd at 80

END
*/
}
