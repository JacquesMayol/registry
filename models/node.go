package models

import (
	"fmt"
	u "registry/utils"
)

type Node struct {
	// gorm.Model
	ID            int64   `gorm:"primaryKey;autoIncrement;" json:"id"`
	ParentID      int64   `json:"parent_id"`
	SelfID        int64   `json:"self_id"`
	Key           string  `json:"key"`
	Value         string  `json:"value"`
	Level         int64   `json:"level"`
	Description   string  `json:"description"`
	ExtParams    *string  `gorm:"column:ext_params" json:"ext_params"`
}

// TableName overrides the table name used by User to `profiles`
func (Node) TableName() string {
	return "registry"
}

type Sibling struct {
	ID int64
}


// Validate
// This struct function validate the required parameters sent through the http request body
// returns message and true if the requirement is met
func (node *Node) Validate() (map[string]interface{}, bool) {

	if node.Key <= "" {
		return u.Message(false, "Key is not recognized"), false
	}

	if node.Value <= "" {
		return u.Message(false, "Value is not recognized"), false
	}

	// All the required parameters are present
	return u.Message(true, "success"), true
}

func (node *Node) SetSibling() {
	parent := GetNode(node.ParentID)

	var sibling int64
	db.Raw("select MAX(id) from \"public\".\"get_dict\"(?)", parent.Key).Scan(&sibling)

	sibling += 1
	node.SelfID = sibling
}

func (node *Node) Create() map[string]interface{} {

	if resp, ok := node.Validate(); !ok {
		fmt.Println("Node validate is not OK")
		return resp
	}

	node.SetSibling()
	fmt.Printf("%+v\n", node)

	GetDB().Create(node)

	resp := u.Message(true, "success")
	resp["node"] = node
	return resp
}

func (node *Node) Update() map[string]interface{} {

	if resp, ok := node.Validate(); !ok {
		return resp
	}

	db.Save(node)

	// GetNode(node.ID)

	resp := u.Message(true, "success")
	resp["node"] = node
	return resp
}

// GetNode
// Get node by Node.ID
func GetNode(id int64) *Node {

	node := &Node{}
	err := GetDB().Table("registry").Where("id = ?", id).First(node).Error

	if err != nil {
		return nil
	}

	return node
}

// GetNodes
// Get list of all exists nodes
func GetNodes() []*Node {
	nodes := make([]*Node, 0)
	err := GetDB().Table("registry").Find(&nodes).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nodes
}

func GetChild(parent uint) []*Node {
	nodes := make([]*Node, 0)
	err := GetDB().Table("registry").Where("parent_id = ?", parent).Find(&nodes).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nodes
}

func GetBranch(parent uint) []*Node {
	nodes := make([]*Node, 0)
	err := GetDB().Table("registry").Where("parent_id = ?", parent).Find(&nodes).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nodes
}
