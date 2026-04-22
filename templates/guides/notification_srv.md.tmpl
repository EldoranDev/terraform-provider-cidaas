---
page_title: "Notification service (notification-srv)"
description: |-
  How to use the cidaas Terraform provider with notification-srv: template groups, template types, templates, graph datasources, limitations, and legacy templates-srv.
---

# Notification service (notification-srv)

This guide describes **notification-srv** resources and datasources in the cidaas Terraform provider, how they differ from **legacy templates-srv**, and how to complete common workflows.

## Two template stacks

| Stack | Provider resources | Backend |
| --- | --- | --- |
| **Notification service** | `cidaas_notifications_template_group`, `cidaas_notification_template_type`, `cidaas_notification_template`, graph datasources | **notification-srv** under `/{notifications_context_path}/…` |
| **Legacy** | `cidaas_template_group`, `cidaas_template` | **templates-srv** (separate API paths) |

- Set optional provider argument **`notifications_context_path`** (default: `notifications-srv`) so all notification-srv clients use the same URL prefix. Legacy `cidaas_template` / `cidaas_template_group` **do not** use this setting.
- Prefer the notification-srv resources for **new** infrastructure as code.

## Authentication and scopes

- The provider uses **client credentials** from environment variables `TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID` and `TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET` (non-interactive app).
- Template APIs typically require scopes such as **`cidaas:templates_read`**, **`cidaas:templates_write`**, and **`cidaas:templates_delete`** (exact enforcement is on notification-srv and your tenant). Grant these on the client used by Terraform.

## Use cases and building blocks

### Custom template type, then templates in a group

1. Define **`cidaas_notification_template_type`** with `category = "custom"`, `template_key`, `description`, and `communication_methods` (see [Casing](#api-casing-and-terraform-state) below).
2. Create **`cidaas_notification_template`** rows per `group_id`, `template_key`, `communication_method`, `locale`, and `message_format`.
3. Use **`depends_on`** if Terraform cannot infer order.

See [examples/resources/cidaas_notification_template_type/resource.tf](https://github.com/Cidaas/terraform-provider-cidaas/blob/master/examples/resources/cidaas_notification_template_type/resource.tf) and [examples/resources/cidaas_notification_template/resource.tf](https://github.com/Cidaas/terraform-provider-cidaas/blob/master/examples/resources/cidaas_notification_template/resource.tf).

### Create a template group and copy templates

- Use **`cidaas_notifications_template_group`** with `group_id`, `tg_type` (`cidaas`, `developer`, or `reminder`), `description` (10–600 characters), `default_locale`, and optionally `copy_from_group_id` / `copy_locale_mappings`.
- Configure **`comm_setting_*`** blocks when you need per-channel service setup ids; resolve ids with **`data.cidaas_notification_service_setups`** where applicable.

See [examples/resources/cidaas_notifications_template_group/resource.tf](https://github.com/Cidaas/terraform-provider-cidaas/blob/master/examples/resources/cidaas_notifications_template_group/resource.tf).

### Discover existing templates or groups (graph API)

- **`data.cidaas_notification_templates`** — POST `graph/templates/` with a JSON **`graph_filter`** body.
- **`data.cidaas_notification_template_groups`** — POST `graph/templategroups/` with a JSON **`graph_filter`** body.

See [examples/datasources/cidaas_notification_templates.tf](https://github.com/Cidaas/terraform-provider-cidaas/blob/master/examples/datasources/cidaas_notification_templates.tf) and [examples/datasources/cidaas_notification_template_groups.tf](https://github.com/Cidaas/terraform-provider-cidaas/blob/master/examples/datasources/cidaas_notification_template_groups.tf).

### System template types (category `cidaas`)

- **System** template types are **pre-provisioned**; you usually **`terraform import`** them and only change allowed fields (for example **`custom_attributes`**), depending on resource logic.
- You **cannot** create a new system template type from scratch via Terraform the same way as a custom type.

### Legacy templates-srv

- Use **`cidaas_template`** and **`cidaas_template_group`** only for existing configurations that still target templates-srv.
- Do not mix legacy resources with notification-srv paths for the same logical workflow without understanding both APIs.

## API casing and Terraform state

- **notification-srv** JSON uses **lowercase** communication methods (`email`, `sms`, `ivr`, `push`).
- For **`cidaas_notification_template_type`**, the provider accepts **case-insensitive** `communication_methods` in configuration and **normalizes** them to lowercase in plan/state so Terraform stays consistent with the API (no spurious diffs after apply).
- **`cidaas_notification_template`** already expects lowercase **`communication_method`** and **`message_format`** (`html`, `text`, `media`) to match the API.

## Known limitations

- **Admin-only or UI-only** flows may exist that are not exposed as Terraform resources; this provider only covers what is implemented in code against the public HTTP APIs.
- **Graph filters** are passed through as JSON strings; invalid filter shapes fail at runtime with API errors—validate against notification-srv graph filter rules.
- **Import** identifiers are resource-specific (for example template document **`id`** for `cidaas_notification_template`); see each resource’s documentation.
- **Destroy** behavior for system or protected objects follows API rules; some resources may warn instead of deleting.

## Related registry documentation

- Provider schema: **`base_url`**, **`notifications_context_path`**
- Resources: **`cidaas_notifications_template_group`**, **`cidaas_notification_template_type`**, **`cidaas_notification_template`**
- Data sources: **`cidaas_notification_templates`**, **`cidaas_notification_template_groups`**, **`cidaas_notification_service_setups`**

Run `go generate ./...` after changing schemas so the [Terraform Registry](https://registry.terraform.io/providers/cidaas/cidaas/latest/docs) docs stay in sync.
