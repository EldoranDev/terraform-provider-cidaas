---
page_title: "Migration: Classic templates vs notification template types"
description: |-
  How cidaas_template / cidaas_template_group relate to cidaas_notification_template_type,
  and how to plan a move to the notification template-type model.
---

# Migration guide: Classic templates and template groups → notification template types

This guide explains how the **classic** Terraform resources `cidaas_template` and `cidaas_template_group` relate to `cidaas_notification_template_type` (added in provider **3.5.5**), and how to approach a migration in Terraform.

It does **not** replace Cidaas product or release documentation. Confirm timelines, deprecation, and UI/API behavior with your Cidaas account team or release notes for your tenant version.

## Concepts: three different layers

| Layer | Terraform resource | What it manages |
|--------|----------------------|-----------------|
| **Template group** | `cidaas_template_group` | A **group id** used to scope communication templates to apps (e.g. `default`, custom group ids). Referenced from `cidaas_app` as `template_group_id`. |
| **Template instance** | `cidaas_template` | **Localized content** for a channel: `template_key`, `locale`, `template_type` (EMAIL, SMS, …), `content`, `subject`, etc. One row per key + locale + type combination. |
| **Template type (notification)** | `cidaas_notification_template_type` | **Type definition** in the notification service: which **communication methods** are allowed, which **system/custom/context attributes** exist, **processing** / **usage** / **verification** types, optional **template_group_ids** binding. This is **metadata/schema**, not the email/SMS body text. |

So:

- **`cidaas_notification_template_type` is not a renamed `cidaas_template`.**  
  It configures **what** a template family allows (attributes, methods, groups), not the **per-locale body** of an email or SMS.
- **`cidaas_template` remains the resource for actual template content** in the classic flow, unless your Cidaas rollout has moved that responsibility to another surface (only your product docs can state that).

You will often **run classic templates and notification template types in parallel** during a transition: types define the contract; templates still hold content until your project switches fully.

## When to use `cidaas_notification_template_type`

Use it when you need to:

- Define or adjust **custom** template types (`category = "custom"`) with full CRUD in Terraform.
- Tune **system** template types (`category = "cidaas"`): import existing types and, where allowed, update **`custom_attributes`** only (system types are pre-provisioned; they cannot be created from scratch in Terraform).

Required provider scopes (same family as templates): `cidaas:templates_read`, `cidaas:templates_write`, `cidaas:templates_delete`.

Examples: `examples/resources/cidaas_notification_template_type/resource.tf`.

## Mapping from classic resources

| Classic (`cidaas_template`) | Notification template type |
|-----------------------------|----------------------------|
| `template_key` | Same **key** identifies the family (e.g. `VERIFY_USER`). Template type keys must match provider rules (e.g. uppercase letters, digits, `_`, `-`). |
| `group_id` (system templates) | Optional list `template_group_ids` on `cidaas_notification_template_type` to tie the type to groups such as `default` or your custom `cidaas_template_group.group_id`. |
| `template_type` (EMAIL/SMS/…) | `communication_methods` on the template type (set of allowed channels). |
| `processing_type`, `usage_type`, `verification_type` | `processing_types`, `usage_types`, `verification_types` on the template type. |
| N/A (content lives on `cidaas_template`) | **Not represented** on `cidaas_notification_template_type`; content stays on `cidaas_template` until you migrate off that API. |

`cidaas_template_group` is still the resource for **creating/naming groups**; template types can reference those same ids via `template_group_ids`.

## Suggested migration steps (Terraform)

1. **Inventory**  
   List all `cidaas_template` and `cidaas_template_group` resources and note: `template_key`, `group_id`, locales, `template_type`, and system vs custom flags.

2. **Align keys**  
   Ensure each logical `template_key` you want in the notification model matches the naming rules for `cidaas_notification_template_type.template_key`.

3. **Import or create template types**  
   - **System** types: `terraform import cidaas_notification_template_type.<name> <template_key>` then manage only what the provider allows (typically `custom_attributes`).  
   - **Custom** types: add `cidaas_notification_template_type` blocks with `category = "custom"` and the desired `communication_methods`, attributes, and `template_group_ids`.

4. **Keep classic templates during transition**  
   Leave `cidaas_template` (and groups) in place until Cidaas and your runbooks say content is fully handled elsewhere. Remove or shrink them only after validation (plans, sends, regression tests).

5. **Apps**  
   If you use `template_group_id` on `cidaas_app`, keep groups consistent with `template_group_ids` on the template types you rely on.

6. **Validate**  
   Run `terraform plan` in non-production, test verification and notification flows, then roll out.

## Provider version

Use a provider version that includes `cidaas_notification_template_type` (≥ **3.5.5**). See [CHANGELOG](https://github.com/Cidaas/terraform-provider-cidaas/blob/master/CHANGELOG.md) for fixes and enhancements per release.

## Further reading

- Registry: [cidaas_template](https://registry.terraform.io/providers/Cidaas/cidaas/latest/docs/resources/template), [cidaas_template_group](https://registry.terraform.io/providers/Cidaas/cidaas/latest/docs/resources/template_group). The `cidaas_notification_template_type` page appears in the registry after provider docs are published for your release; until then, use the schema in the provider source (`internal/resources/resource_notification_template_type.go`) and the examples below.
- Examples: [`examples/resources/cidaas_notification_template_type/`](https://github.com/Cidaas/terraform-provider-cidaas/tree/master/examples/resources/cidaas_notification_template_type).
