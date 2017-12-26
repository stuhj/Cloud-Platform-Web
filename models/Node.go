package models



type Node struct {
	Created_at			string
	Updated_at			string
	Hypervisor_hostname string
	Vcpus				int64
	Memory_mb			int64
	Local_gb			int64
	Vcpus_used			int64
	Memory_mb_used		int64
	Local_gb_used		int64
	Running_vms			int64
	Host_ip				string
}
