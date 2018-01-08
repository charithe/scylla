package graphql

import (
	"io"
	"io/ioutil"

	"github.com/charithe/scylla/graphql/internal"
	"github.com/charithe/scylla/graphql/symbols"
)

type Parser struct {
	rlp *internal.RagelParser
}

func NewParser(r io.Reader) (*Parser, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return &Parser{
		rlp: internal.NewRagelParser(data),
	}, nil
}

func (p *Parser) Start() (symbols.Symbol, error) {
	return p.rlp.Start()
}
