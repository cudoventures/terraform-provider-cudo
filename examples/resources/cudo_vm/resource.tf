resource "cudo_storage_disk" "my-storage-disk" {
  data_center_id = "gb-bournemouth-1"
  id             = "storage-disk-1"
  size_gib       = 100
}

# pick a specific data center and machine type
resource "cudo_vm" "my-vm" {
  depends_on     = [cudo_storage_disk.my-storage-disk]
  id             = "instance-1"
  machine_type   = "intel-broadwell"
  data_center_id = "gb-bournemouth-1"
  memory_gib     = 2
  vcpus          = 1
  boot_disk = {
    image_id = "debian-12"
    size_gib = 10
  }
  storage_disks = [
    {
      disk_id = "storage-disk-1"
    }
  ]
  ssh_key_source = "project"
  start_script   = <<EOF
                     touch /multiline-script.txt
                     echo  $PWD > /current-dir.txt
                     EOF
}
