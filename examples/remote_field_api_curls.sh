#!/usr/bin/env bash
# Example cURL calls for fieldsetup-srv/fields with remoteFieldSettings (GROUPING).
# This file defines the canonical apiClientSetup / apiAccess shape for remote fields.
# It may differ from the webhook auth API (e.g. detail keys: apikeyDetails, totpDetails,
# basicAuthDetails, oAuthDetails). callOnce can be set to true or false.
# Replace the Bearer token and BASE_URL when needed (tokens are short-lived).

AUTH_HEADER='Authorization: Bearer '
BASE_URL='https://....cidaas.eu'

# --- 1. APIKEY ---
curl --location "${BASE_URL}/fieldsetup-srv/fields" \
  --header "$AUTH_HEADER" \
  --header 'Content-Type: application/json' \
  --data '{
    "parent_group_id": "DEFAULT",
    "is_group": true,
    "dataType": "GROUPING",
    "fieldKey": "remote_apikey_field",
    "fieldType": "CUSTOM",
    "scopes": ["profile"],
    "localeTexts": [{ "locale": "en-US", "name": "Remote API Key field" }],
    "remoteFieldSettings": {
      "apiClientSetup": {
        "communicationEP": "https://api.example.com/users?identifier={{sub}}",
        "httpMethod": "GET",
        "apiAccess": {
          "apiAccessType": "APIKEY",
          "apikeyDetails": {
            "apikey": "your-api-key",
            "apikey_placeholder": "X-Api-Key",
            "apikey_placement": "header"
          }
        }
      },
      "callOnce": false
    }
  }'

# --- 2. TOTP ---
curl --location "${BASE_URL}/fieldsetup-srv/fields" \
  --header "$AUTH_HEADER" \
  --header 'Content-Type: application/json' \
  --data '{
    "parent_group_id": "DEFAULT",
    "is_group": true,
    "dataType": "GROUPING",
    "fieldKey": "remote_totp_field",
    "fieldType": "CUSTOM",
    "scopes": ["profile"],
    "localeTexts": [{ "locale": "en-US", "name": "Remote TOTP field" }],
    "remoteFieldSettings": {
      "apiClientSetup": {
        "communicationEP": "https://api.example.com/users?identifier={{sub}}",
        "httpMethod": "GET",
        "apiAccess": {
          "apiAccessType": "TOTP",
          "totpDetails": {
            "totpkey": "base32-secret",
            "totp_placeholder": "X-Totp-Key",
            "totp_placement": "header"
          }
        }
      },
      "callOnce": true
    }
  }'

# --- 3. BASIC_AUTH ---
curl --location "${BASE_URL}/fieldsetup-srv/fields" \
  --header "$AUTH_HEADER" \
  --header 'Content-Type: application/json' \
  --data '{
    "parent_group_id": "DEFAULT",
    "is_group": true,
    "dataType": "GROUPING",
    "fieldKey": "remote_basicauth_field",
    "fieldType": "CUSTOM",
    "scopes": ["profile"],
    "localeTexts": [{ "locale": "en-US", "name": "Remote Basic Auth field" }],
    "remoteFieldSettings": {
      "apiClientSetup": {
        "communicationEP": "https://api.example.com/users?identifier={{sub}}",
        "httpMethod": "GET",
        "apiAccess": {
          "apiAccessType": "BASIC_AUTH",
          "basicAuthDetails": {
            "user": "api-user",
            "password": "api-password"
          }
        }
      },
      "callOnce": false
    }
  }'

# --- 4. CIDAAS_OAUTH2 ---
curl --location "${BASE_URL}/fieldsetup-srv/fields" \
  --header "$AUTH_HEADER" \
  --header 'Content-Type: application/json' \
  --data '{
    "parent_group_id": "DEFAULT",
    "is_group": true,
    "dataType": "GROUPING",
    "fieldKey": "remote_cidaas_oauth2_field",
    "fieldType": "CUSTOM",
    "scopes": ["profile"],
    "localeTexts": [{ "locale": "en-US", "name": "Remote Cidaas OAuth2 field" }],
    "remoteFieldSettings": {
      "apiClientSetup": {
        "communicationEP": "https://api.example.com/users?identifier={{sub}}",
        "httpMethod": "GET",
        "apiAccess": {
          "apiAccessType": "CIDAAS_OAUTH2",
          "oAuthDetails": {
            "client_id": "YOUR_CLIENT_ID",
            "req_scopes": "cidaas:users_read"
          }
        }
      },
      "callOnce": true
    }
  }'

# --- 5. GEN_OAUTH2 ---
curl --location "${BASE_URL}/fieldsetup-srv/fields" \
  --header "$AUTH_HEADER" \
  --header 'Content-Type: application/json' \
  --data '{
    "parent_group_id": "DEFAULT",
    "is_group": true,
    "dataType": "GROUPING",
    "fieldKey": "remote_gen_oauth2_field",
    "fieldType": "CUSTOM",
    "scopes": ["profile"],
    "localeTexts": [{ "locale": "en-US", "name": "Remote Gen OAuth2 field" }],
    "remoteFieldSettings": {
      "apiClientSetup": {
        "communicationEP": "https://api.example.com/users?identifier={{sub}}",
        "httpMethod": "GET",
        "apiAccess": {
          "apiAccessType": "GEN_OAUTH2",
          "oAuthDetails": {
            "client_id": "external-client-id",
            "client_secret": "external-client-secret",
            "wellknownUrl": "https://auth.example.com/.well-known/openid-configuration",
            "req_scopes": "openid profile"
          }
        }
      },
      "callOnce": true
    }
  }'
