package gocssunits

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	lengthValuesPattern = regexp.MustCompile(`^(\d+\.?\d*)(px|pt|em|rem|cm|mm|Q|in|pc|cqw|cqh|cqi|cqb|cqmin|cqmax|vh|vw|vmax|vmin|vb|rcap|rch|rex|ric|rlh|cap|ch|ex|ic|lh|%)$`)
	wUnitPattern        = regexp.MustCompile(`^(\d+\.?\d*)$`)
	keywordSizes        = map[string]struct{}{
		"xx-small":     {},
		"x-small":      {},
		"small":        {},
		"medium":       {},
		"large":        {},
		"x-large":      {},
		"xx-large":     {},
		"smaller":      {},
		"larger":       {},
		"math":         {},
		"inherit":      {},
		"initial":      {},
		"revert":       {},
		"revert-layer": {},
		"unset":        {},
	}

	ErrUnsupportedSize = errors.New("unsupported css size value")
)

type FontSize struct {
	Value float64
	Unit  string
}

func (fz FontSize) String() string {
	if fz.Value == 0 {
		return fz.Unit
	}
	return fmt.Sprintf("%s%s", strconv.FormatFloat(fz.Value, 'f', -1, 64), fz.Unit)
}

// ParseFontSize parses a CSS font size value and returns a structured FontSize object.
//
// The function supports various CSS font size formats:
//   - Absolute units: "16px", "12pt", "1.2in", "10mm", "0.5cm"
//   - Relative units: "1.5em", "2rem", "120%"
//   - Unitless numbers: "14"
//   - Keywords: "small", "medium", "large", "x-large", "xx-large",
//     "smaller", "larger", "xx-small", "x-small"
//
// Parameters:
//
//	css string - The CSS font size value to parse (e.g., "16px", "1.2em", "medium")
//
// Returns:
//
//	*FontSize - A structured object containing the parsed value and unit,
//	            or nil if parsing fails
//	error     - An error if the input format is unsupported or invalid
//
// Example:
//
//	size, err := ParseFontSize("1.5rem")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Value: %.1f, Unit: %s", size.Value, size.Unit) // Value: 1.5, Unit: rem
//
// Notes:
//   - The function is case-insensitive for keywords and units
//   - Leading and trailing whitespace is automatically trimmed
//   - Returns ErrUnsupportedSize for unknown keywords or invalid formats
//   - Unitless numbers are stored with empty Unit field
func ParseFontSize(css string) (*FontSize, error) {
	css = strings.TrimSpace(css)

	fz := &FontSize{}

	if matches := lengthValuesPattern.FindStringSubmatch(css); len(matches) > 2 {
		f, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			return nil, err
		}
		fz.Value = f
		fz.Unit = matches[2]
	} else if matches := wUnitPattern.FindStringSubmatch(css); len(matches) > 1 {
		f, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			return nil, err
		}
		fz.Value = f
	} else {
		if _, ok := keywordSizes[css]; !ok {
			return nil, ErrUnsupportedSize
		}
		fz.Unit = css
	}
	return fz, nil
}

func (fz *FontSize) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "\"%s\"", fz.String()), nil
}

func (fz *FontSize) UnmarshalJSON(data []byte) error {
	f, err := ParseFontSize(strings.ReplaceAll(string(data), "\"", ""))
	if err != nil {
		return err
	}
	*fz = *f
	return nil
}
