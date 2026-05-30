package resources

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestParseNotificationsTemplateGroupLocaleImportID(t *testing.T) {
	t.Parallel()
	g, l, ok := parseNotificationsTemplateGroupLocaleImportID("my_group/ar")
	if !ok || g != "my_group" || l != "ar" {
		t.Fatalf("got %q %q %v", g, l, ok)
	}
	_, _, ok = parseNotificationsTemplateGroupLocaleImportID("invalid")
	if ok {
		t.Fatal("expected invalid import id")
	}
}

func TestLocaleModelFromPlan_id(t *testing.T) {
	t.Parallel()
	m := localeModelFromPlan(notificationsTemplateGroupLocaleModel{
		GroupID:         types.StringValue("g1"),
		Locale:          types.StringValue("de-DE"),
		CopyFromGroupID: types.StringValue("default"),
		CopyFromLocale:  types.StringValue("en"),
	})
	if m.ID.ValueString() != "g1/de-DE" {
		t.Fatalf("id %q", m.ID.ValueString())
	}
}
