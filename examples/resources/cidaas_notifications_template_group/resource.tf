# Example: notification-srv template group (not legacy cidaas_template_group).
# Creating a group usually copies templates from a source group (default: "default").
resource "cidaas_notifications_template_group" "example_dev" {
  group_id           = "example_dev_group"
  tg_type            = "developer"
  description        = "Example developer template group created with Terraform (min 10 chars)."
  default_locale     = "en"
  copy_from_group_id = "default"
}
