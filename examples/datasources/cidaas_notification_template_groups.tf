# Example: list template groups via notification-srv graph API (POST graph/templategroups/).
data "cidaas_notification_template_groups" "all" {
  graph_filter = jsonencode({
    size = 50
  })
}
