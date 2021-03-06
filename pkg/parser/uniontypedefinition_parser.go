package parser

import (
	"github.com/jensneuse/graphql-go-tools/pkg/lexing/keyword"
	"github.com/jensneuse/graphql-go-tools/pkg/lexing/token"
)

func (p *Parser) parseUnionTypeDefinition(description *token.Token, index *[]int) error {

	start, err := p.readExpect(keyword.UNION, "parseUnionTypeDefinition")
	if err != nil {
		return err
	}

	unionName, err := p.readExpect(keyword.IDENT, "parseUnionTypeDefinition")
	if err != nil {
		return err
	}

	definition := p.makeUnionTypeDefinition()
	definition.Name = p.putByteSliceReference(unionName.Literal)

	if description != nil {
		definition.Position.MergeStartIntoStart(description.TextPosition)
	} else {
		definition.Position.MergeStartIntoStart(start.TextPosition)
	}

	err = p.parseDirectives(&definition.DirectiveSet)
	if err != nil {
		return err
	}

	shouldParseMembers := p.peekExpect(keyword.EQUALS, true)

	for shouldParseMembers {

		member, err := p.readExpect(keyword.IDENT, "parseUnionTypeDefinition")
		if err != nil {
			return err
		}

		definition.UnionMemberTypes = append(definition.UnionMemberTypes, p.putByteSliceReference(member.Literal))

		shouldParseMembers = p.peekExpect(keyword.PIPE, true)
	}

	definition.Position.MergeStartIntoEnd(p.TextPosition())
	*index = append(*index, p.putUnionTypeDefinition(definition))
	return nil
}
