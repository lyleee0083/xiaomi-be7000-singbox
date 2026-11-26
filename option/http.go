package option

import (
	"context"
	"reflect"

	"github.com/sagernet/sing-box/schema"
	"github.com/sagernet/sing/common/byteformats"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/json"
	"github.com/sagernet/sing/common/json/badjson"
	"github.com/sagernet/sing/common/json/badoption"
)

type HTTP2Options struct {
	IdleTimeout             badoption.Duration       `json:"idle_timeout,omitempty"`
	KeepAlivePeriod         badoption.Duration       `json:"keep_alive_period,omitempty"`
	StreamReceiveWindow     *byteformats.MemoryBytes `json:"stream_receive_window,omitempty"`
	ConnectionReceiveWindow *byteformats.MemoryBytes `json:"connection_receive_window,omitempty"`
	MaxConcurrentStreams    int                      `json:"max_concurrent_streams,omitempty"`
}

type _HTTPClientOptions struct {
	Tag                     string               `json:"tag,omitempty"`
	Engine                  string               `json:"engine,omitempty" enum:"go,apple"`
	Version                 int                  `json:"version,omitempty" enum:"0,1,2"`
	DisableVersionFallback  bool                 `json:"disable_version_fallback,omitempty"`
	Headers                 badoption.HTTPHeader `json:"headers,omitempty"`
	HTTP2Options            HTTP2Options         `json:"-"`
	DefaultOutbound         bool                 `json:"-"`
	DisableEmptyDirectCheck bool                 `json:"-"`
	ResolveOnDetour         bool                 `json:"-"`
	DirectResolver          bool                 `json:"-"`
	OutboundTLSOptionsContainer
	DialerOptions
}

type (
	HTTPClient        _HTTPClientOptions
	HTTPClientOptions _HTTPClientOptions
)

func (h HTTPClient) Options() HTTPClientOptions {
	options := HTTPClientOptions(h)
	options.Tag = ""
	return options
}

func (o HTTPClientOptions) IsEmpty() bool {
	if o.Tag != "" {
		return false
	}
	o.DefaultOutbound = false
	o.ResolveOnDetour = false
	o.DirectResolver = false
	return reflect.ValueOf(_HTTPClientOptions(o)).IsZero()
}

func (o HTTPClientOptions) MarshalJSON() ([]byte, error) {
	if o.Tag != "" {
		return json.Marshal(o.Tag)
	}
	return badjson.MarshallObjects(_HTTPClientOptions(o), httpClientVariant(_HTTPClientOptions(o)))
}

func (o *HTTPClientOptions) UnmarshalJSONContext(ctx context.Context, content []byte) error {
	if len(content) > 0 && content[0] == '"' {
		*o = HTTPClientOptions{}
		return json.Unmarshal(content, &o.Tag)
	}
	var options _HTTPClientOptions
	err := json.UnmarshalContext(ctx, content, &options)
	if err != nil {
		return err
	}
	err = unmarshalHTTPVersionOptions(ctx, content, &options, options.Version, &options.HTTP2Options)
	if err != nil {
		return err
	}
	options.Tag = ""
	*o = HTTPClientOptions(options)
	return nil
}

func (h HTTPClient) MarshalJSON() ([]byte, error) {
	return badjson.MarshallObjects(_HTTPClientOptions(h), httpClientVariant(_HTTPClientOptions(h)))
}

func (h *HTTPClient) UnmarshalJSONContext(ctx context.Context, content []byte) error {
	err := json.UnmarshalContext(ctx, content, (*_HTTPClientOptions)(h))
	if err != nil {
		return err
	}
	return unmarshalHTTPVersionOptions(ctx, content, (*_HTTPClientOptions)(h), h.Version, &h.HTTP2Options)
}

func unmarshalHTTPVersionOptions(ctx context.Context, content []byte, baseStruct any, version int, http2Options *HTTP2Options) error {
	switch version {
	case 1:
		return json.UnmarshalContextDisallowUnknownFields(ctx, content, baseStruct)
	case 0, 2:
		return badjson.UnmarshallExcludedContext(ctx, content, baseStruct, http2Options)
	default:
		return E.New("unknown HTTP version: ", version)
	}
}

func httpVersionVariant(version int, http2Options HTTP2Options) any {
	switch version {
	case 1:
		return nil
	case 0, 2:
		return http2Options
	default:
		return nil
	}
}

func httpClientVariant(options _HTTPClientOptions) any {
	return httpVersionVariant(options.Version, options.HTTP2Options)
}

func describeHTTPClientObject(builder schema.Builder) (*schema.Node, error) {
	node := schema.StrictObject()
	err := builder.FlattenStruct(node, reflect.TypeFor[HTTPClient]())
	if err != nil {
		return nil, err
	}
	err = builder.FlattenStruct(node, reflect.TypeFor[HTTP2Options]())
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (h HTTPClient) DescribeSchema(builder schema.Builder) (*schema.Node, error) {
	return builder.Define("HTTPClient", func() (*schema.Node, error) {
		return describeHTTPClientObject(builder)
	})
}

func (o HTTPClientOptions) DescribeSchema(builder schema.Builder) (*schema.Node, error) {
	return builder.Define("HTTPClientReference", func() (*schema.Node, error) {
		clientObject, err := describeHTTPClientObject(builder)
		if err != nil {
			return nil, err
		}
		clientObject.Properties.Remove("tag")
		return schema.AnyOf(schema.TagReferenceNode("http_client"), clientObject), nil
	})
}
