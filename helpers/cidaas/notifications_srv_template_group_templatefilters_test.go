package cidaas

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseTemplateFiltersLocales(t *testing.T) {
	t.Parallel()
	raw := json.RawMessage(`{"locales":["en","de","de-DE","en-US"]}`)
	got, err := ParseTemplateFiltersLocales(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 || got[0] != "en" || got[2] != "de-DE" {
		t.Fatalf("got %v", got)
	}
}

func TestParseTemplateFiltersLocales_empty(t *testing.T) {
	t.Parallel()
	got, err := ParseTemplateFiltersLocales(nil)
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Fatalf("got %v want nil", got)
	}
}

func TestIsNotificationSrvTemplatesAlreadyExistError(t *testing.T) {
	t.Parallel()
	if !IsNotificationSrvTemplatesAlreadyExistError(fmt.Errorf("already templates found for the locales [ar]")) {
		t.Fatal("expected match")
	}
	if IsNotificationSrvTemplatesAlreadyExistError(fmt.Errorf("other")) {
		t.Fatal("expected no match")
	}
}
