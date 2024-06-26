---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "geoserver_wms_store Resource - terraform-provider-geoserver"
subcategory: ""
description: |-
  
---

# geoserver_wms_store (Resource)



## Example Usage

```terraform
resource "geoserver_wms_store" "geoplateforme" {
  workspace_name = geoserver_workspace.ign.name
  name           = "geoplateforme"

  default = false
  enabled = true

  capabilities_url = "https://data.geopf.fr/wms-r/wms?SERVICE=WMS&amp;Version=1.3.0&amp;Request=GetCapabilities"

  max_connections = 10
  read_timeout = 20
  connection_timeout = 10
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `capabilities_url` (String) URL of the remote WMS server capability URL.
- `name` (String) Name of the WMS store. Used to compute the id of the resource.
- `workspace_name` (String) Name of the workspace owning the WMS store. Used to compute the id of the resource.

### Optional

- `connection_timeout` (Number) Number of seconds before considering a connection request in timeout. Default value is 30.
- `default` (Boolean) Mark the WMS store as default. Default value is false.
- `description` (String) Description of the WMS store. Default value is empty.
- `disable_connection_on_failure` (Boolean) Don't try to connect to remote server if failure occurs. Default value is false.
- `enabled` (Boolean) Mark the WMS store as enabled. Default value is true.
- `max_connections` (Number) Number of maximum parallel connections allowed to the remote server. Default value is 6
- `read_timeout` (Number) Number of seconds before considering a read request in timeout. Default value is 60.

### Read-Only

- `id` (String) The ID of this resource.


