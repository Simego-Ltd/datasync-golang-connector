package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type dataResponse struct {
	List        string                   `json:"list"`
	Filter      string                   `json:"filter"`
	OrderBy     string                   `json:"orderby"`
	CurrentPage int                      `json:"currentpage"`
	NextPage    int                      `json:"nextpage"`
	TotalPages  int                      `json:"totalpages"`
	Count       int                      `json:"count"`
	Data        []map[string]interface{} `json:"data"`
}
type schemaInfo struct {
	UpdateKeyColumn         string
	UpdateKeyColumnDataType string
	BlobNameColumnFormat    string
	BlobURLFormat           string `json:"BlobUrlFormat"`
	Columns                 []schemaColumnInfo
}
type schemaColumnInfo struct {
	Name      string
	DataType  string
	MaxLength int
	AllowNull bool
	Unique    bool
	System    bool
	ReadOnly  bool
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/list/list1/schema", handleList1Schema)
	mux.HandleFunc("/list/list1", handleList1Data)
	mux.HandleFunc("/list", handleLists)
	mux.HandleFunc("/", handleNotFound)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	fmt.Println("Server Ready ... ", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, mux))
}

func (d *dataResponse) AppendRow(row map[string]interface{}) {
	if d.Data == nil {
		d.Data = make([]map[string]interface{}, 0, 100)
	}
	d.Data = append(d.Data, row)
	d.Count = len(d.Data)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func handleLists(w http.ResponseWriter, r *http.Request) {

	lists := []string{"list1"}
	if b, err := json.Marshal(&lists); err == nil {
		w.Write(b)
	} else {
		http.Error(w, err.Error(), 500)
	}
}

func handleList1Schema(w http.ResponseWriter, r *http.Request) {

	lists := schemaInfo{
		UpdateKeyColumn:         "ID",
		UpdateKeyColumnDataType: "System.Int32",
		Columns: []schemaColumnInfo{
			{Name: "ID", DataType: "System.Int32", Unique: true},
			{Name: "Name", DataType: "System.String"},
		},
	}

	if b, err := json.Marshal(&lists); err == nil {
		w.Write(b)
	} else {
		http.Error(w, err.Error(), 500)
	}
}

func handleList1Data(w http.ResponseWriter, r *http.Request) {

	lists := dataResponse{List: "list1", CurrentPage: 1, NextPage: 0, TotalPages: 1}

	for i := 0; i < 10; i++ {
		row := map[string]interface{}{"ID": i, "Name": fmt.Sprintf("Item %d", i)}
		lists.AppendRow(row)
	}

	if b, err := json.Marshal(&lists); err == nil {
		w.Write(b)
	} else {
		http.Error(w, err.Error(), 500)
	}
}
