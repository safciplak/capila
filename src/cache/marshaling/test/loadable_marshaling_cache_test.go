package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/eko/gocache/v2/cache"
	"github.com/stretchr/testify/suite"

	capilaCache "github.com/safciplak/capila/src/cache"
	"github.com/safciplak/capila/src/cache/marshaling"
)

type LoadableMarshalerCacheSuite struct {
	suite.Suite
}

// TestClientHelpersSuite Runs the testsuite
func TestCWICacheSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(LoadableMarshalerCacheSuite))
}

type MarshalingExample struct {
	template func() interface{}
	loadFunc func(ctx context.Context, key interface{}) (interface{}, error)
	invoke   func(sut capilaCache.InterfaceGetter, message string)
	name     string
	key      string
}

const setGraceTime = 200 * time.Millisecond

//nolint:funlen // function is used to group test case categories
func (test *LoadableMarshalerCacheSuite) TestValidValues() {
	var examples = []*MarshalingExample{
		{
			name: "pointer to string",
			key:  "some key",
			template: func() interface{} {
				return new(string)
			},
			loadFunc: func(ctx context.Context, key interface{}) (interface{}, error) {
				v := "Some value"

				return &v, nil
			},
			invoke: func(sut capilaCache.InterfaceGetter, message string) {
				data, err := sut.Get(context.Background(), "some key")
				test.Nil(err, message)

				value, ok := data.(*string)
				test.True(ok, message)

				test.Equal("Some value", *value, message)
			},
		},
		{
			name: "struct",
			key:  "some key for struct",
			template: func() interface{} {
				return &exampleStruct{}
			},
			loadFunc: func(ctx context.Context, key interface{}) (interface{}, error) {
				return &exampleStruct{"Some value A", 7}, nil
			},
			invoke: func(sut capilaCache.InterfaceGetter, message string) {
				data, err := sut.Get(context.Background(), "some key for struct")
				test.Nil(err, message)

				value, ok := data.(*exampleStruct)
				test.True(ok, message)

				test.Equal("Some value A", value.ValueA, message)
			},
		},

		{
			name: "binaryMarshaler",
			key:  "some key for binaryMarshaler",
			template: func() interface{} {
				return &binaryMarshaler{}
			},
			loadFunc: func(ctx context.Context, key interface{}) (interface{}, error) {
				return &binaryMarshaler{"Some value A"}, nil
			},
			invoke: func(sut capilaCache.InterfaceGetter, message string) {
				data, err := sut.Get(context.Background(), "some key for binaryMarshaler")
				test.Nil(err, message)

				value, ok := data.(*binaryMarshaler)
				test.True(ok, message)

				test.Equal("Some value A", value.ValueA, message)
			},
		},
		{
			name: "binaryMarshalerWithValidation",
			key:  "some key for binaryMarshalerWithValidation",
			template: func() interface{} {
				return &binaryMarshalerWithValidation{}
			},
			loadFunc: func(ctx context.Context, key interface{}) (interface{}, error) {
				return &binaryMarshalerWithValidation{"Valid content"}, nil
			},
			invoke: func(sut capilaCache.InterfaceGetter, message string) {
				data, err := sut.Get(context.Background(), "some key for binaryMarshalerWithValidation")
				test.Nil(err, message)

				value, ok := data.(*binaryMarshalerWithValidation)
				test.True(ok, message)

				test.Equal("Valid content", value.Value, message)
			},
		},
	}

	for _, example := range examples {
		func(example *MarshalingExample) {
			test.Run("when not in cache -> loadFunc is called, value is stored in cache"+example.name, func() {
				inMem := NewMapStore()

				cacheManager := cache.New(inMem)

				loadFuncCallCounter := 0
				loadFunc := func(ctx context.Context, key interface{}) (interface{}, error) {
					loadFuncCallCounter++

					return example.loadFunc(ctx, key)
				}
				sut := marshaling.NewLoadableMarshalerCache(cacheManager, loadFunc, example.template)

				message := formatMessage(example)
				example.invoke(sut, message)

				// assert loadFunc is called
				test.Equal(1, loadFuncCallCounter, message)

				// wait for store to be updated (or timeout)
				select {
				case <-time.After(setGraceTime):
				case <-inMem.SetCalled:
				}

				// assert contains value
				data, err := inMem.Get(context.Background(), example.key)

				test.Nil(err)
				test.NotNil(data)
			})

			test.Run("when in cache -> cache is returned, loadFunc is not called"+example.name, func() {
				inMem := NewMapStore()

				cacheManager := cache.New(inMem)

				loadFuncCallCounter := 0
				loadFunc := func(ctx context.Context, key interface{}) (interface{}, error) {
					loadFuncCallCounter++

					return example.loadFunc(ctx, key)
				}
				sut := marshaling.NewLoadableMarshalerCache(cacheManager, loadFunc, example.template)

				message := formatMessage(example)
				example.invoke(sut, message)

				// assert loadFunc is called
				test.Equal(1, loadFuncCallCounter, message)

				// wait for store to be updated (or timeout)
				select {
				case <-time.After(setGraceTime):
				case <-inMem.SetCalled:
				}

				// assert contains value
				data, err := inMem.Get(context.Background(), example.key)
				test.Nil(err)
				test.NotNil(data)

				// invoke a second time
				example.invoke(sut, message)

				// imply cache is returned
				// assert loadFunc is not called a second time
				test.Equal(1, loadFuncCallCounter, message)
			})
		}(example)
	}
}

//nolint:funlen // function is used to group test case categories
func (test *LoadableMarshalerCacheSuite) TestInValidValues() {
	var marshalFailsExamples = []*MarshalingExample{
		{
			name: "binaryMarshalerWithValidation",
			key:  "some key for binaryMarshalerWithValidation",
			template: func() interface{} {
				return &binaryMarshalerWithValidation{}
			},
			loadFunc: func(ctx context.Context, key interface{}) (interface{}, error) {
				// e at second place will trigger error in marshal
				return &binaryMarshalerWithValidation{"Very invalid content"}, nil
			},
			invoke: func(sut capilaCache.InterfaceGetter, message string) {
				data, err := sut.Get(context.Background(), "some key for binaryMarshalerWithValidation")
				test.Nil(err, message)

				value, ok := data.(*binaryMarshalerWithValidation)
				test.True(ok, message)

				// expect to receive the data even when it will not be stored
				test.Equal("Very invalid content", value.Value, message)
			},
		},
	}

	var unmarshalFailsExamples = []*MarshalingExample{
		{
			name: "string",
			key:  "some key",
			template: func() interface{} {
				return ""
			},
			loadFunc: func(ctx context.Context, key interface{}) (interface{}, error) {
				v := "Some value"

				return v, nil
			},
			invoke: func(sut capilaCache.InterfaceGetter, message string) {
				data, err := sut.Get(context.Background(), "some key")
				test.Nil(err, message)

				value, ok := data.(string)
				test.True(ok, message)

				test.Equal("Some value", value, message)
			},
		},
		{
			name: "binaryMarshalerWithValidation",
			key:  "some key for binaryMarshalerWithValidation",
			template: func() interface{} {
				return &binaryMarshalerWithValidation{}
			},
			loadFunc: func(ctx context.Context, key interface{}) (interface{}, error) {
				// m at first place will trigger error in unmarshal, while it was valid to marshal
				return &binaryMarshalerWithValidation{"maybe valid content"}, nil
			},
			invoke: func(sut capilaCache.InterfaceGetter, message string) {
				data, err := sut.Get(context.Background(), "some key for binaryMarshalerWithValidation")
				test.Nil(err, message)

				value, ok := data.(*binaryMarshalerWithValidation)
				test.True(ok, message)

				// expect to receive the data from the loadFunc even when it was stored (but invalid)
				test.Equal("maybe valid content", value.Value, message)
			},
		},
	}

	for _, example := range marshalFailsExamples {
		func(example *MarshalingExample) {
			test.Run("When not in cache, and marshal fails -> value is returned, but not stored "+example.name, func() {
				inMem := NewMapStore()

				cacheManager := cache.New(inMem)

				loadFuncCallCounter := 0
				loadFunc := func(ctx context.Context, key interface{}) (interface{}, error) {
					loadFuncCallCounter++

					return example.loadFunc(ctx, key)
				}

				sut := marshaling.NewLoadableMarshalerCache(cacheManager, loadFunc, example.template)
				message := formatMessage(example)
				example.invoke(sut, message)

				// assert loadFunc is called
				test.Equal(1, loadFuncCallCounter, message)

				// wait for store to be updated (or timeout)
				select {
				case <-time.After(setGraceTime):
				case <-inMem.SetCalled:
					test.Fail("it should not store when validation fails")
				}

				// assert no value stored in cache
				data, err := inMem.Get(context.Background(), example.key)

				test.NotNil(err)
				test.Nil(data)
			})
		}(example)
	}

	for _, example := range unmarshalFailsExamples {
		func(example *MarshalingExample) {
			test.Run("When in cache, and unmarshal fails -> loadFunc is called, cache is updated "+example.name, func() {
				inMem := NewMapStore()

				cacheManager := cache.New(inMem)

				loadFuncCallCounter := 0
				loadFunc := func(ctx context.Context, key interface{}) (interface{}, error) {
					loadFuncCallCounter++

					return example.loadFunc(ctx, key)
				}

				sut := marshaling.NewLoadableMarshalerCache(cacheManager, loadFunc, example.template)
				message := formatMessage(example)
				example.invoke(sut, message)

				// assert loadFunc is called
				test.Equal(1, loadFuncCallCounter, message)

				// wait for store to be updated (or timeout)
				select {
				case <-time.After(setGraceTime):
				case <-inMem.SetCalled:
				}

				// assert contains value
				data, err := inMem.Get(context.Background(), example.key)
				test.Nil(err)
				test.NotNil(data)

				// invoke a second time
				example.invoke(sut, message)

				// assert loadFunc is called a second time
				test.Equal(2, loadFuncCallCounter, message)

				// wait for store to be updated (or timeout)
				select {
				case <-time.After(setGraceTime):
				case <-inMem.SetCalled:
				}

				// assert contains value
				data, err = inMem.Get(context.Background(), example.key)
				test.Nil(err)
				test.NotNil(data)
			})
		}(example)
	}
}

func formatMessage(example *MarshalingExample) string {
	return fmt.Sprintf("while executing example '%s'", example.name)
}
