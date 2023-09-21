package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"HttpFetch"
	sig "signature"

	scale "github.com/loopholelabs/scale"
	scalefunc "github.com/loopholelabs/scale/scalefunc"
)

type FetchExtension struct {
}

// implementor

type HttpConnector struct {
}

func (fe *FetchExtension) New(c *HttpFetch.HttpConfig) (HttpFetch.HttpConnector, error) {
	return &HttpConnector{}, nil
}

func (hc *HttpConnector) Fetch(u *HttpFetch.ConnectionDetails) (HttpFetch.HttpResponse, error) {
	r := HttpFetch.HttpResponse{}
	// Do the actual fetch here...

	resp, err := http.Get(u.Url)
	if err != nil {
		return r, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}

	return HttpFetch.HttpResponse{
		StatusCode: int32(resp.StatusCode),
		Body:       body,
	}, nil
}

func main() {
	fmt.Printf("Running scale function with ext...\n")

	sgo, err := scalefunc.Read("../local-testfngo-latest.scale")
	if err != nil {
		panic(err)
	}

	testfn(sgo)

	srs, err := scalefunc.Read("../local-testfnrs-latest.scale")
	if err != nil {
		panic(err)
	}

	testfn(srs)

	sts, err := scalefunc.Read("../local-testfnts-latest.scale")
	if err != nil {
		panic(err)
	}

	testfn(sts)

}

func testfn(fn *scalefunc.Schema) {
	fmt.Printf("Running scale function with ext... %s\n", fn.Language)

	ext_impl := &FetchExtension{}

	ctx := context.Background()

	// runtime
	config := scale.NewConfig(sig.New).
		WithContext(ctx).
		WithFunctions([]*scalefunc.Schema{fn}).
		WithExtension(HttpFetch.New(ext_impl))

	r, err := scale.New(config)
	if err != nil {
		panic(err)
	}

	i, err := r.Instance(nil)
	if err != nil {
		panic(err)
	}

	sigctx := sig.New()

	sigctx.Context.MyString = "hello world"
	err = i.Run(context.Background(), sigctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data(1)[%s] from scaleFunction: %s\n", fn.Language, sigctx.Context.MyString)

	err = i.Run(context.Background(), sigctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data(2)[%s] from scaleFunction: %s\n", fn.Language, sigctx.Context.MyString)

}
