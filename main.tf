provider "coveo" {
    organization_id = "${var.coveo_organization_id}"
    api_key =  "${var.coveo_api_key}"
}

resource "coveo_source" "my-source" {
    name = "my-push-source-from-terraform"
    type = "PUSH"
    visibility = "SHARED"
    push_enabled = false
}

output "source-id" {
    value = "${coveo_source.my-source.id}"
}

output "organisation-id" {
    value = "${var.coveo_organization_id}"
}