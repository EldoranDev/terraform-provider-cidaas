# Example: discover template rows via notification-srv graph API (POST graph/templates/).
data "cidaas_notification_templates" "by_group" {
  graph_filter = jsonencode({
    size = 20
    filters = [
      {
        searchfield        = "groupId"
        fieldOperator      = "STRING_SEARCH_OPTIONS"
        searchfieldOptions = { value = "default", match = "EXACT" }
      }
    ]
  })
}
