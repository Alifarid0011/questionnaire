package validation

import (
	"log"
	"sync"

	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
)

// ResourceProvider holds the singleton instances for validation and translation.
type ResourceProvider struct {
	validate   *validator.Validate
	translator ut.Translator
	once       sync.Once
}

// Provider is the globally accessible singleton instance of the ResourceProvider.
var Provider = &ResourceProvider{}

// setup performs the thread-safe, one-time configuration of validator and translator resources.
func (r *ResourceProvider) setup() {
	r.validate = validator.New()

	faLocale := fa.New()
	uni := ut.New(faLocale, faLocale)

	tr, found := uni.GetTranslator("fa")
	if !found {
		log.Fatal("failed to get translator for 'fa' locale")
	}
	r.translator = tr

	if err := fa_translations.RegisterDefaultTranslations(r.validate, r.translator); err != nil {
		log.Fatalf("failed to register translations: %v", err)
	}
}

// Initialize ensures the internal setup runs only once.
func (r *ResourceProvider) Initialize() {
	r.once.Do(r.setup)
}

// Validator returns the configured singleton *validator.Validate instance.
// It ensures initialization occurs lazily on first access.
func (r *ResourceProvider) Validator() *validator.Validate {
	r.Initialize()
	return r.validate
}

// Translator returns the configured singleton ut.Translator instance.
// It ensures initialization occurs lazily on first access.
func (r *ResourceProvider) Translator() ut.Translator {
	r.Initialize()
	return r.translator
}
