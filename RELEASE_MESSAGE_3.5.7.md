## Release 3.5.7 – Terraform Provider Cidaas

### Summary

This release ships **registration field** computed-field fixes, **app** guest-login and MFA drift fixes, **notification-srv** template-type behavior and documentation, **CI** lint configuration, and **registry documentation** updates (guide, examples, regenerated `docs/`).

### Enhancements

- **Documentation:** New [Notification service (notification-srv)](docs/guides/notification_srv.md) guide (`templates/guides/notification_srv.md.tmpl`): two-stack model (notification-srv vs legacy templates-srv), `notifications_context_path`, scopes, use cases, limitations, and casing notes. README links to the guide; provider index example version **3.5.7**.
- **Examples:** `cidaas_notification_template`, `cidaas_notifications_template_group`, graph datasources for templates and template groups; notification template type examples aligned with API lowercase conventions.
- **Registry:** `go generate ./...` refreshed resource/data source pages for notification-srv and related resources.

### Bug fixes

- **Registration field:** Computed `base_data_type` no longer stays unknown after apply for GROUPING; empty string when the API omits it.
- **App – `allow_guest_login_groups`:** `group_type` read from API into state.
- **App – MFA:** `time_interval_in_seconds` normalized to null when API sends 0 or omits the field.
- **`cidaas_notification_template_type`:** Default **`owner`** for custom types; **`communication_methods`** validated case-insensitively and normalized in plan/state to match notification-srv (lowercase), avoiding apply inconsistencies and API `35002`-style rejections.

### CI

- **Lint:** `.ci/lint/configs/golang/.golangci-standard.yml` for shared GitLab template.
- **Lint:** lint:diff job fix.

### Pre-release verification

Maintainers: run `go test ./...`, `go vet ./...`, `golangci-lint run` (or CI), and `go generate ./...` with a clean `git diff` on `docs/` before tagging. See **Pre-release review checklist** in `DEVELOPER_GUIDE.md`.
