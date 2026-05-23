---
page_title: "Migration: Template group locale copy (3.5.14)"
description: |-
  How to migrate from copy_from_group_id / copy_locale_mappings on cidaas_notifications_template_group
  to cidaas_notifications_template_group_locale per locale.
---

# Migration: Template group locale copy (provider 3.5.14)

Provider **3.5.14** splits **locale copy** out of **`cidaas_notifications_template_group`** into **`cidaas_notifications_template_group_locale`**. This guide replaces any draft that added a required **`locales`** set on the group resource (that approach was not released).

See [CHANGELOG](https://github.com/Cidaas/terraform-provider-cidaas/blob/master/CHANGELOG.md) and [Notification service (notification-srv)](notification_srv.md).

## What changed

| Before (≤ 3.5.13) | After (3.5.14) |
| --- | --- |
| `copy_from_group_id` on the group | Removed from **`cidaas_notifications_template_group`** |
| `copy_locale_mappings` on the group | One **`cidaas_notifications_template_group_locale`** per target locale |
| Optional bulk locale control on group | Per-locale **Create** (copy), **Read** (`templatefilters`), **Destroy** (bulk-delete templates) |

Group **Create** no longer sends API `copy`. notification-srv may still create templates/locales from `default` when the group is created; manage only the locales you want with locale resources.

## Choose `tg_type` correctly

On **`cidaas_notifications_template_group`**, **`tg_type`** (API `tgType`) must match how the group is used:

| `tg_type` | Use when |
| --- | --- |
| **`cidaas`** | Standard Cidaas platform notification groups (welcome, password reset, verification, etc.). Typical for tenant-wide system mail tied to built-in flows. |
| **`developer`** | Custom developer templates: you define **`cidaas_notification_template_type`** (`category = "custom"`) and **`cidaas_notification_template`** content for that group. Use the **same** `group_id` on template types/templates as on the group. |
| **`reminder`** | Reminder / follow-up notification groups. |

**Rule:** If you manage **custom** templates with **`cidaas_notification_template`** / **`cidaas_notification_template_type`**, set **`tg_type = "developer"`** on the group (not `cidaas`). Platform groups use **`tg_type = "cidaas"`**.

## Step-by-step migration

### 1. Update the group resource

Remove **`copy_from_group_id`** and **`copy_locale_mappings`**. Keep metadata only:

```hcl
resource "cidaas_notifications_template_group" "my_group" {
  group_id       = "my_group"
  tg_type        = "cidaas" # or "developer" / "reminder" — see table above
  description    = "My template group (min 10 characters)."
  default_locale = "en"
}
```

### 2. Add one locale resource per locale

Each block maps to API `copy.locale[]` with one `{ from, to }` pair. Every **`locale`** must exist as templates in the group after apply.

```hcl
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

Discover codes on the source group (often **`default`**): `GET …/templategroups/default/templatefilters` (e.g. `en`, `de`, `de-DE`, `en-US` — tenant-specific).

### 3. Import existing locales (recommended if templates already exist)

If the group already has templates in the tenant, **import** so plan does not show four creates:

```bash
terraform import 'cidaas_notifications_template_group_locale.en' 'my_group/en'
terraform import 'cidaas_notifications_template_group_locale.de' 'my_group/de'
```

Import id format: **`{group_id}/{locale}`**.

Then run **`terraform plan`** — expect **no changes** if templatefilters match.

Without import, plan shows **`+ create`** for each new locale resource; **apply** adopts existing locales when copy is skipped.

### 4. Apply order

1. **`cidaas_notifications_template_group`** (or refresh/import if it already exists).
2. **`cidaas_notifications_template_group_locale`** resources (apply or import).

Use **`depends_on = [cidaas_notifications_template_group.my_group]`** if needed; referencing **`group_id`** is usually enough.

### 5. Destroying locales

- **Destroy** a locale resource → bulk-deletes all templates for that locale (`DELETE …/templates?groupId=…&locale=…`).
- You **cannot** destroy the locale that equals the group's **`default_locale`** until you change **`default_locale`** on the group first.
- Destroy **locale** resources before destroying the **group**.

### 6. Extra locales in the tenant (not in Terraform)

The provider **does not** auto-remove locales that exist in `templatefilters` but have no locale resource. To remove them: import a locale resource and destroy it, or delete in Admin.

## Before / after example

**Before:**

```hcl
resource "cidaas_notifications_template_group" "my_group" {
  group_id       = "my_group"
  tg_type        = "cidaas"
  description    = "..."
  default_locale = "en"
  copy_from_group_id = "default"
  copy_locale_mappings = [
    { from = "en", to = "en" },
    { from = "en", to = "de" },
  ]
}
```

**After:**

```hcl
resource "cidaas_notifications_template_group" "my_group" {
  group_id       = "my_group"
  tg_type        = "cidaas"
  description    = "..."
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

## Related guides

- [Migration: Classic templates → notification template types](migration-template-to-notification-template-type.md) — `cidaas_template` / template types vs notification-srv (orthogonal to locale copy).
- [Notification service (notification-srv)](notification_srv.md) — day-to-day workflows.
