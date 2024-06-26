---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "geoserver_gwc_wms_layer Resource - terraform-provider-geoserver"
subcategory: ""
description: |-
  
---

# geoserver_gwc_wms_layer (Resource)



## Example Usage

```terraform
# Example using a well known WMS service
resource "geoserver_gwc_wms_layer" "ign_orthophotos" {
    name = "ign_orthophotos"
    blobstore_id = geoserver_gwc_S3_blobstore.s3_blobstore.blobstore_id
    wms_url="https://wxs.ign.fr/essentiels/geoportail/r/wms?SERVICE=WMS"
    wms_layer="ORTHOIMAGERY.ORTHOPHOTOS"
    mime_formats=["image/png","image/jpeg"]
    grid_subsets=[geoserver_gwc_gridset.gridset_3857.name]
    metatile_height=4
    metatile_width=4
    expire_duration_clients=3600
}

# Example using a layer provided with a GeoServer instance
# The geoserver URL and the layer name are provided with variables
resource "geoserver_gwc_wms_layer" "nexsis_fdp_osm" {
    name = "nexsis_fdp_osm"
    blobstore_id = geoserver_gwc_S3_blobstore.s3_blobstore.blobstore_id
    wms_url=var.geoserver_wms
    wms_layer=var.fdp_osm_layer
    mime_formats=["image/png","image/jpeg"]
    grid_subsets=[geoserver_gwc_gridset.gridset_3857.name]
    metatile_height=4
    metatile_width=4
    expire_duration_clients=3600
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `blobstore_id` (String) Blobstore to use for storing the tiles.
- `grid_subset` (Block Set, Min: 1) List of the grids supported for this cache. (see [below for nested schema](#nestedblock--grid_subset))
- `metatile_height` (Number) Height of the meta tiles.
- `metatile_width` (Number) Width of the meta tiles.
- `mime_formats` (List of String) List of the mime formats upported for this cache.
- `name` (String) Name of the cache layer. Use as resource id.
- `wms_layer` (String) Identifier of the layer to cache.
- `wms_url` (String) URL of the server hosting the WMS.

### Optional

- `allow_cache_bypass` (Boolean) Allow bypass of the cache. Default to false.
- `backend_timeout` (Number) Timeout of the backend. Default to 120.
- `background_color` (String) Background color when requesting maps.
- `enabled` (Boolean) Is the layer cache enabled? Default to true.
- `expire_duration_cache` (Number) Server cache expire duration. Default to 0.
- `expire_duration_clients` (Number) Client cache expire duration. Default to 0.
- `gutter_size` (Number) Size of the gutter to use for the meta-tiles. Default to 0.
- `transparent` (Boolean) Request tiles with transparent background? Default to true.
- `vendor_parameters` (String) Additional vendor parameters to the service.
- `wms_version` (String) WMS version to use when requesting service. Default to 1.3.0

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--grid_subset"></a>
### Nested Schema for `grid_subset`

Required:

- `name` (String)

Optional:

- `max_cached_level` (Number)
- `min_cached_level` (Number)


