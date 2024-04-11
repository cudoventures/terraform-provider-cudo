resource "cudo_storage_disk" "my-storage-disk" {
  data_center_id = "gb-bournemouth-1"
  id             = "my-disk"
  size_gib       = 100
}