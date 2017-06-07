//go:generate bitfanDoc -codec encoder,decoder
package rubydebugcodec

import (
	"io"

	"github.com/go-playground/validator"
	"github.com/k0kubun/pp"
	"github.com/mitchellh/mapstructure"
	"github.com/vjeantet/bitfan/codecs/lib"
)

type encoder struct {
	w       io.Writer
	options encoderOptions

	log lib.Logger
}

type encoderOptions struct {
}

func NewEncoder(w io.Writer) *encoder {
	return &encoder{
		w:       w,
		options: encoderOptions{},
	}
}

func (e *encoder) SetOptions(conf map[string]interface{}, logger lib.Logger, cwl string) error {
	e.log = logger

	if err := mapstructure.Decode(conf, &e.options); err != nil {
		return err
	}

	// validates codecs's user options
	if err := validator.New(&validator.Config{TagName: "validate"}).Struct(&e.options); err != nil {
		return err
	}

	return nil
}

func (e *encoder) Encode(data map[string]interface{}) error {
	pp.Fprintln(e.w, data)
	return nil
}