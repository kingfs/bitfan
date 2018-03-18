// Code generated by "bitfanDoc "; DO NOT EDIT
package webfan

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "webfan",
  ImportPath: "github.com/vjeantet/bitfan/processors/webfan",
  Doc:        "",
  DocShort:   "Reads events from standard input",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:           "processors.CommonOptions",
        Alias:          ",squash",
        Doc:            "",
        Required:       false,
        Type:           "processors.CommonOptions",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Codec",
        Alias:          "",
        Doc:            "The codec used for posted data. Input codecs are a convenient method for decoding\nyour data before it enters the pipeline, without needing a separate filter in your bitfan pipeline\n\nDefault decode http request as plain text, response is json encoded.\nSet multiple codec with role to customize",
        Required:       false,
        Type:           "codec",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "codec => plain { role=>\"encoder\"} codec => json { role=>\"decoder\"}",
      },
      &doc.ProcessorOption{
        Name:           "Uri",
        Alias:          "uri",
        Doc:            "URI path /_/path",
        Required:       true,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Pipeline",
        Alias:          "pipeline",
        Doc:            "Path to pipeline's configuration to execute on request\nThis configuration should contains only a filter section an a output like ```output{pass{}}```",
        Required:       true,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Headers",
        Alias:          "headers",
        Doc:            "Headers to send back into outgoing response",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "{\"X-Processor\" => \"bitfan\"}",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}