package frontparser

import (
	"bytes"
	"errors"

	"gopkg.in/yaml.v2"
)

// error messages
var ErrorFrontmatterNotFound error = errors.New("frontmatter header not found")

var FmYAMLDelimiter []byte = []byte("---")

// returns whether the provided content
func HasFrontmatterHeader(input []byte) bool {
	// remove heading and trailing spaces (and CR, LF, ...)
	input = bytes.TrimSpace(input)
	// test for frontmatter delimiter
	if !bytes.HasPrefix(input, FmYAMLDelimiter) {
		return false
	}
	// trim heading frontmatter delimiter
	input = bytes.TrimPrefix(input, FmYAMLDelimiter)
	// split on frontmatter delimiter to separate frontmatter from the rest
	elements := bytes.SplitN(input, FmYAMLDelimiter, 2)
	if len(elements) != 2 {
		// malformed input
		return false
	}
	// parse frontmatter to validate it is valid YAML
	var out map[string]interface{} = make(map[string]interface{})
	err := yaml.Unmarshal(input, out)
	return err == nil
}

//
func ParseFrontmatter(input []byte, out interface{}) error {
	// remove heading and trailing spaces (and CR, LF, ...)
	input = bytes.TrimSpace(input)
	// test for frontmatter delimiter
	if !bytes.HasPrefix(input, FmYAMLDelimiter) {
		return errors.New("heading frontmatter delimiter not found")
	}
	// trim heading frontmatter delimiter
	input = bytes.TrimPrefix(input, FmYAMLDelimiter)
	// split on frontmatter delimiter to separate frontmatter from the rest
	elements := bytes.SplitN(input, FmYAMLDelimiter, 2)
	if len(elements) != 2 {
		// malformed input
		return errors.New("more than two frontmatter delimiters were found")
	}
	// parse frontmatter to validate it is valid YAML
	return yaml.Unmarshal(input, out)
}
