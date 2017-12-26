package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"reflect"
	"cloud-web/models"
)

type NodeController struct {
	beego.Controller
}


func (c *NodeController)Get(){
	connurl := fmt.Sprintf("root:0421@tcp(192.168.1.20:3306)/nova?charset=utf8")
	db, err := sql.Open("mysql", connurl)
	var nodes []models.Node
	if err != nil{
		beego.Error(err.Error())
	}else{
		defer db.Close()
		query := "select "
		t := reflect.TypeOf(models.Node{})
		num := t.NumField()
		beego.Info(num)
		for i := 0; i < num; i++{
			query += t.Field(i).Name
			if i != num - 1{
				query += ", "
			}
		}
		query += " from compute_nodes;"
		beego.Info(query)
		r, err := db.Query(query)
		if err != nil {
			beego.Error(err.Error())
		}
		defer r.Close()
		for r.Next(){
			node := models.Node{}
			r.Scan(&node.Created_at, &node.Updated_at, &node.Hypervisor_hostname, &node.Vcpus, &node.Memory_mb, &node.Local_gb,
				&node.Vcpus_used, &node.Memory_mb_used, &node.Local_gb_used,&node.Running_vms,&node.Host_ip)
			nodes = append(nodes, node)
		}

	}
	beego.Info(nodes)
	c.Data["Nodes"] = nodes
	c.TplName = "nodes.html"
}
