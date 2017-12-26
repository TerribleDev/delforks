package main

import (
	"context"
	"flag"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	list := flag.Bool("list", false, "list all forks")
	del := flag.Bool("del", false, "delete listed")
	flag.Parse()
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "MY TOKEN HERE"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	if *list && *del {
		panic("List and del are mutually exclusive!")
	}
	if *list {

		repos, _, _ := client.Repositories.List(ctx, "TerribleDev", nil)
		for _, v := range repos {
			if *v.Fork {
				print(*v.Name + " ")
			}
		}
		return
	}
	if *del {

		for _, v := range flag.Args() {
			res, _ := client.Repositories.Delete(ctx, "TerribleDev", v)
			println(res.StatusCode)
		}

	}

}
