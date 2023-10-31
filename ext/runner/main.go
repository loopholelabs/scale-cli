package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

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

var tryErrorsNew = false
var tryErrorsFetch = false

func (fe *FetchExtension) New(c *HttpFetch.HttpConfig) (HttpFetch.HttpConnector, error) {
	if tryErrorsNew {
		return nil, errors.New("Error from New")
	}
	return &HttpConnector{}, nil
}

func (hc *HttpConnector) Fetch(u *HttpFetch.ConnectionDetails) (HttpFetch.HttpResponse, error) {
	if tryErrorsFetch {
		return HttpFetch.HttpResponse{}, errors.New("Error from Fetch")
	}

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

	functions := []string{
		//		"../local-testfngo-latest.scale",
		//		"../local-testfnrs-latest.scale",
		"../local-testfnts-latest.scale",
	}

	for _, n := range functions {

		sgo, err := scalefunc.Read(n)
		if err != nil {
			panic(err)
		}

		tryErrorsNew = false
		tryErrorsFetch = false
		testfn(sgo)
		// Make sure things work if the extension returns an error...
		tryErrorsNew = true
		testfn(sgo)
		tryErrorsNew = false
		tryErrorsFetch = true
		testfn(sgo)

	}

}

func testfn(fn *scalefunc.Schema) {
	fmt.Printf("Running scale function with ext... %s\n", fn.Language)

	ext_impl := &FetchExtension{}

	ctx := context.Background()

	// runtime
	config := scale.NewConfig(sig.New).
		WithContext(ctx).
		WithFunctions([]*scalefunc.Schema{fn}).
		WithExtension(HttpFetch.New(ext_impl)).
		WithStdout(os.Stdout).
		WithRawOutput(true)

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
