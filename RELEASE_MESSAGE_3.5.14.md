## Release 3.5.14 – Terraform Provider Cidaas

### Summary

Template group **locale copy** moves from **`cidaas_notifications_template_group`** to new **`cidaas_notifications_template_group_locale`**. Upgrade any configuration that used **`copy_from_group_id`** or **`copy_locale_mappings`** on the group.

> **Note:** A draft **3.5.13** release note described a required **`locales`** argument on the group; that design was **not** released. Use **3.5.14** and this document instead.

### Breaking changes

#### `cidaas_notifications_template_group`

- **Removed:** `copy_from_group_id`, `copy_locale_mappings`.
- **Unchanged:** `group_id`, `tg_type`, `description`, `default_locale`, `comm_setting_*`, `owner`.
- Group **create** does not send API `copy`. notification-srv may still seed locales from `default` per server rules.

#### What you must change

1. Delete `copy_from_group_id` and `copy_locale_mappings` from every `cidaas_notifications_template_group` block.
2. Add one **`cidaas_notifications_template_group_locale`** per locale you manage (see migration example below).
3. Set **`tg_type`** correctly:
   - **`cidaas`** — platform / system notification groups.
   - **`developer`** — groups used with **`cidaas_notification_template`** and custom **`cidaas_notification_template_type`** (same `group_id`).
   - **`reminder`** — reminder groups.
4. If templates already exist in the tenant, **import** locale resources (`{group_id}/{locale}`) or apply once (create is idempotent when locale exists in `templatefilters`).

Full steps: [Migration: Template group locale copy (3.5.14)](docs/guides/migration-notifications-template-group-locales.md).

### Enhancements

#### `cidaas_notifications_template_group_locale` (new)

- **Create:** `PUT` with `copy.locale[]` (`copy_from_group_id` + `copy_from_locale` → `locale`).
- **Read:** `GET …/templatefilters` — removes from state if locale missing.
- **Destroy:** bulk-delete templates for that locale; blocked if `locale` is still the group's `default_locale`.
- **Import:** `terraform import … {group_id}/{locale}`

### Migration example

```hcl
resource "cidaas_notifications_template_group" "my_group" {
  group_id       = "my_group"
  tg_type        = "developer" # use "cidaas" for platform groups; "developer" for custom template groups
  description    = "My template group (min 10 characters)."
  default_locale = "en"
}

resource "cidaas_notifications_template_group_locale" "en" {
  group_id           = cidaas_notifications_template_group.my_group.group_id
  locale             = "en"
  copy_from_group_id = "default"
  copy_from_locale   = "en"
}

resource "cidaas_notifications_template_group_locale" "de" {
  group_id           = cidaas_notifications_template_group.my_group.group_id
  locale             = "de"
  copy_from_group_id = "default"
  copy_from_locale   = "en"
}
```

### Import existing locales

```bash
terraform import 'cidaas_notifications_template_group_locale.en' 'my_group/en'
terraform import 'cidaas_notifications_template_group_locale.de' 'my_group/de'
```

### Pre-release verification

Maintainers: `go test ./...`, `go generate ./...`, verify `CHANGELOG.md` matches this file.
