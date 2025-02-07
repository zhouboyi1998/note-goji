package controller

import (
	"encoding/json"
	"net/http"
	"note-goji/src/repository"
)

// One 根据id查询命令
func One(w http.ResponseWriter, r *http.Request) {
	command := repository.One(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(command)
}

// List 查询命令列表
func List(w http.ResponseWriter, r *http.Request) {
	commandArray := repository.List(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commandArray)
}

// Insert 新增命令
func Insert(w http.ResponseWriter, r *http.Request) {
	result, commandName := repository.Insert(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result":  result,
		"command": commandName,
	})
}

// InsertBatch 批量新增命令
func InsertBatch(w http.ResponseWriter, r *http.Request) {
	result := repository.InsertBatch(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Update 修改命令
func Update(w http.ResponseWriter, r *http.Request) {
	result := repository.Update(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// UpdateBatch 批量修改命令
func UpdateBatch(w http.ResponseWriter, r *http.Request) {
	result := repository.UpdateBatch(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Delete 删除命令
func Delete(w http.ResponseWriter, r *http.Request) {
	result, objectId := repository.Delete(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteBatch 批量删除命令
func DeleteBatch(w http.ResponseWriter, r *http.Request) {
	result, objectIds := repository.DeleteBatch(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": result,
		"_ids":   objectIds,
	})
}

// Select 查询命令
func Select(w http.ResponseWriter, r *http.Request) {
	command := repository.Select(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(command)
}

// NameList 查询命令名称列表
func NameList(w http.ResponseWriter, r *http.Request) {
	nameArray := repository.NameList(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nameArray)
}
