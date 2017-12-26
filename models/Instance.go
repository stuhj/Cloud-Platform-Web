package models



type Instance struct{
	Vm_state			string
	Uuid				string
	Display_name		string
	Host				string
	Availability_zone	string
	Memory_mb			int64
	Vcpus				int64
	Created_at 			string
	Updated_at			string
	Deleted_at			string
	User_id				string
	Project_id			string
	Image_ref			string
	Kernel_id			string
	Power_state			int64
}

func InitInstance(ins *Instance){
	ins.Display_name = "none"
	ins.Uuid = "none"
	ins.Vm_state = "none"
	ins.Power_state = -1
	ins.Host = "none"

}