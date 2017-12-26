package controllers

import (
	"github.com/astaxie/beego"
	"database/sql"
	"reflect"
	"fmt"
	"cloud-web/models"
)

type InstanceController struct{
	beego.Controller
}


func(c *InstanceController)Get(){
	connurl := fmt.Sprintf("root:0421@tcp(192.168.1.20:3306)/nova?charset=utf8")
	db, err := sql.Open("mysql", connurl)
	var instances []models.Instance

	if err != nil{
		beego.Error(err.Error())
	}else{
		defer db.Close()
		query := "select "
		t := reflect.TypeOf(models.Instance{})
		num := t.NumField()
		beego.Info(num)
		for i := 0; i < num; i++{
			query += t.Field(i).Name
			if i != num - 1{
				query += ", "
			}
		}
		query += " from instances order by id desc;"
		beego.Info(query)
		r, err := db.Query(query)
		if err != nil {
			beego.Error(err.Error())
		}
		defer r.Close()
		for r.Next(){
			instance := models.Instance{}
			r.Scan(&instance.Vm_state,
				&instance.Uuid,&instance.Display_name, &instance.Host,
				&instance.Availability_zone,
				&instance.Memory_mb, &instance.Vcpus,
				&instance.Created_at, &instance.Updated_at, &instance.Deleted_at,
				&instance.User_id, &instance.Project_id, &instance.Image_ref, &instance.Kernel_id,
					&instance.Power_state)
			instances = append(instances, instance)
		}

		for j := 0; j < len(instances); j++{
			v := reflect.ValueOf(&instances[j]).Elem()
			None := reflect.ValueOf("--")
			for i := 0; i < t.NumField(); i++{
				if v.Field(i).String() == ""{
					beego.Info(t.Field(i).Name)
					if v.Field(i).CanSet(){
						v.Field(i).Set(None)
						fmt.Println(v.Field(i))
					}else{
						beego.Info("can not set")
					}
				}
			}
		}
		}



	for _, ins := range instances{
		beego.Info(ins)
	}
	c.Data["Instances"] = instances
	c.TplName = "instances.html"
}
/*
type Instance struct{
	Created_at 			string
	Updated_at			string
	Deleted_at			string
	User_id				string
	Project_id			string
	Image_ref			string
	Kernel_id			string
	Power_state			string
	Vm_state			string
	Memory_mb			int64
	Vcpus				int64
	Hostname			string
	host				string
	Availability_zone	string
	Uuid				string
}
*/