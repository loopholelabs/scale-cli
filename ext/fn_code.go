package testfn

import (
	"HttpFetch"
	"fmt"
	"signature"
)

func Scale(ctx *signature.Context) (*signature.Context, error) {
	c := &HttpFetch.HttpConfig{
		Timeout: 10,
	}

	fetcher, err := HttpFetch.New(c)
	if err != nil {
		ctx.MyString = "Error"
		return ctx, nil
	}

	res, err := fetcher.Fetch(&HttpFetch.ConnectionDetails{
		Url: "https://ifconfig.me",
	})

	if err != nil {
		ctx.MyString = "Error"
		return ctx, nil
	}

	ctx.MyString = fmt.Sprintf("Fetch extension StatusCode=%d Body=%s", res.StatusCode, string(res.Body))

	return signature.Next(ctx)
}
