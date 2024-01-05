package context

import (
	"context"
	"fmt"

	"golang.org/x/text/language"
)

// LanguageKey is the key for the context.
type languageKey string

//nolint:golint,stylecheck // Correct
const LanguageKey languageKey = "language"

// NewContext instantiates an initial application context
func NewContext() context.Context {
	return context.Background()
}

// GetLanguage returns the Language value stored in ctx, if any.
func getLanguage(ctx context.Context) language.Tag {
	lang, ok := ctx.Value(LanguageKey).(languageKey)

	if ok {
		return language.Make(string(lang))
	}

	return language.Make("en")
}

// GetTwoLetterLanguageCode returns the Language in four letters ( for example nl )
func GetTwoLetterLanguageCode(ctx context.Context) string {
	lang := getLanguage(ctx)
	base, _ := lang.Base()

	return base.String()
}

// GetFourLetterLanguageCode returns the Language in four letters ( for example nl-NL )
func GetFourLetterLanguageCode(ctx context.Context) string {
	lang := getLanguage(ctx)
	base, _ := lang.Base()
	region, _ := lang.Region()

	return fmt.Sprintf("%s-%s", base.String(), region.String())
}

// SetLanguage returns a new Context that carries the language.
func SetLanguage(parent context.Context, setLanguage string) context.Context {
	return context.WithValue(parent, LanguageKey, languageKey(setLanguage))
}
