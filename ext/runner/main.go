package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	sig "signature"
	HttpFetch "testext"

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

	sgo, err := scalefunc.Read("../local-testfngo-latest.scale")
	if err != nil {
		panic(err)
	}

	testfn(sgo)
	// Make sure things work if the extension returns an error...
	tryErrorsNew = true
	testfn(sgo)
	tryErrorsNew = false
	tryErrorsFetch = true
	testfn(sgo)

	tryErrorsNew = false
	tryErrorsFetch = false
	srs, err := scalefunc.Read("../local-testfnrs-latest.scale")
	if err != nil {
		panic(err)
	}

	testfn(srs)
	// Make sure things work if the extension returns an error...
	tryErrorsNew = true
	testfn(srs)
	tryErrorsNew = false
	tryErrorsFetch = true
	testfn(srs)

	//	sts, err := scalefunc.Read("../local-testfnts-latest.scale")
	//	if err != nil {
	//		panic(err)
	//	}

	//	testfn(sts)

}

func testfn(fn *scalefunc.V1BetaSchema) {
	fmt.Printf("Running scale function with ext... %s\n", fn.Language)

	ext_impl := &FetchExtension{}

	ctx := context.Background()

	// runtime
	config := scale.NewConfig(sig.New).
		WithContext(ctx).
		WithFunctions([]*scalefunc.V1BetaSchema{fn}).
		WithExtension(HttpFetch.New(ext_impl)).
		WithStdout(os.Stdout)

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
