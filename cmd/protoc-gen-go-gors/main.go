package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	showVersion = flag.Bool("version", false, "print the version and exit")
	omitempty   = flag.Bool("omitempty", true, "omit if google.api is empty")
	bizcode     = flag.Bool("bizcode", true, "open biz code model")
	setGinErr   = flag.Bool("setGinErr", true, "if has error, set to gin context")
)

type Config struct {
	omitempty bool
	bizCode   bool
	setGinErr bool
}

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-http %v\n", release)
		return
	}
	// init config
	config := Config{
		omitempty: *omitempty,
		bizCode:   *bizcode,
		setGinErr: *setGinErr,
	}
	// protoc run plugin
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f, config)
		}
		return nil
	})
}
