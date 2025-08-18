resource "cudo_machine" "machine-1" {
  data_center_id = "no-kristiansand-1"
  id             = "machine-1"
  machine_type   = "sapphire-rapids-h100"
  os             = "ubuntu/jammy"
}
