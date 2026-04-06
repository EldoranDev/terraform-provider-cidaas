# Development vs Master – Comparison & Release Notes

## Branch comparison

**development** is ahead of **master** by a few commits (CI/lint fixes). **master** already includes the 3.5.6-related fixes (registration_field `base_data_type`, app `group_type` and MFA `time_interval`).

### Commits only on development (not on master)

| Commit     | Description |
|-----------|-------------|
| `5ba2b46` | ci(): Fix lint:diff failing |
| `25ee7ac` | ci: add .ci/lint config for golangci-lint (required by shared CI template) |

So the only **code/behavior difference** between the two branches right now is **none**; the extra commits on development are **CI/config only** (lint config path and lint:diff job fix).

---

## Changes in development since last release baseline

All changes that are in **development** (and mostly merged to master) since the pre–3.5.5 state:

### 3.5.5 (already in CHANGELOG)

- **Registration field data source:** List endpoint now uses `fieldsetup-srv/graph/fields`.
- **Notification template type:** New resource `cidaas_notification_template_type`; template type service in client.
- **Bug fixes:** Template type helper (NewHTTPClient / MakeRequest context), hosted page `content` attribute removed.
- **Upgrade note:** Remove `content` from `cidaas_hosted_page` config before upgrading.

### 3.5.6 (already in CHANGELOG)

- **Registration field data source:** `enabled` attribute exposed.
- **Registration field resource:** Remote field settings for GROUPING (RemoteFieldSettings, ApiClientSetup, APIAccessSetup, APIKEY/TOTP/BASIC_AUTH/OAuth2).
- **Bug fixes:** `is_group` removed; `Language` removed from LocaleText; test updates.

### Post–3.5.6 (in development and master)

- **Registration field:** `base_data_type` for GROUPING is empty string (not unknown); fallbacks so computed `base_data_type` is never unknown after apply.
- **App:** `allow_guest_login_groups` now includes `group_type` in state; MFA `time_interval_in_seconds` 0/nil treated as null to avoid drift.

### CI (only on development for now)

- **Lint:** Added `.ci/lint/configs/golang/.golangci-standard.yml` so shared GitLab template finds the config.
- **Lint:** Fix for lint:diff failing in pipeline.

---

## Release notes to be documented

Use this for the **next release** (e.g. 3.5.7 or “3.5.6 + patch”) if you tag from **development** and want one place that lists everything since 3.5.4.

### Enhancements

- **Registration field (GROUPING):** `base_data_type` is now correctly set to empty string when the API returns none; computed value is never unknown after apply.
- **App – allow_guest_login_groups:** `group_type` is now populated in state from the API.
- **App – MFA:** `time_interval_in_seconds` is stored as null when the API returns 0 or omits it, avoiding unnecessary plan drift.

### Bug fixes

- **Registration field:** Avoid “unknown value after apply” for computed `base_data_type` (especially for GROUPING fields).
- **App:** MFA time interval drift when API returns 0 or empty.

### CI / tooling (no user impact)

- Added `.ci/lint/configs/golang/.golangci-standard.yml` for shared GitLab CI template.
- Fixed lint:diff job in pipeline.

---

## Suggested CHANGELOG entry (if releasing 3.5.7)

```markdown
### 3.5.7

#### Bug Fixes

- **Registration field:** Fixed computed `base_data_type` sometimes being unknown after apply (e.g. for GROUPING). It is now set to empty string when the API returns none, with fallbacks so state is always known.
- **App:** `allow_guest_login_groups` now correctly populates `group_type` in state from the API.
- **App:** MFA `time_interval_in_seconds` is now stored as null when the API returns 0 or omits it, preventing unnecessary plan drift.

#### CI

- Added `.ci/lint/configs/golang/.golangci-standard.yml` so the shared GitLab lint template finds the config.
- Fixed lint:diff job in the pipeline.
```

You can paste the “Suggested CHANGELOG entry” into `CHANGELOG.md` when you cut the release.
