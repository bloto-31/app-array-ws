package main

import (
	model "app-array-ws/model"
	//"fmt"

	//"testing"
	"encoding/json"
	"errors"
	"log"

	"net/http" 
	"github.com/gin-gonic/gin"
)

var applicationModel *model.Application
var err error

func getApplicationModel(appModel string) (*model.Application, error) {
	
	 if(appModel == "") {
		return nil, errors.New("empty model") 
	 }
	 var a model.Application
	 err := json.Unmarshal([]byte(appModel), &a)
	 if(err != nil) {
		log.Fatal(`Deserialisation of failed`, appModel)
		return nil, errors.New(appModel) 
	 } 
	 return &a, nil 
}

func getApplication(c *gin.Context) {
	 c.IndentedJSON(http.StatusOK, applicationModel)
}
	
func putApplication(c *gin.Context) {
	 var newAppliModel model.Application
	 // Call BindJSON to bind the received JSON to newAppliModel
	 if err := c.BindJSON(&newAppliModel); err != nil {
        return
	 }

	 c.IndentedJSON(http.StatusOK, newAppliModel)

}

func main() {
	
	 
  	 log.SetPrefix("ws: ")
	 log.SetFlags(0)
		
	 const m string = `{"id":"FOApp","type":"application","components":[{"id":"Database","type":"component","tags":{"group":"core","type":"database"},"provides":[{"id":"raw data","kind":6}]},{"id":"EventBus","type":"component","tags":{"group":"core"},"commands":{"start":{"type":"javascript","steps":["StartComponent"]},"stop":{"type":"javascript","steps":["StopComponent"]}},"provides":[{"id":"raw events","kind":6}]},{"id":"Cache","type":"component","tags":{"group":"core"},"consumes":["raw events","raw data"]},{"id":"PositionService","type":"component","tags":{"group":"TradePosition"},"provides":[{"id":"/api/Position","object":"Position","kind":2,"protocol":"REST"}],"consumes":["raw events","raw data"]},{"id":"Spreadsheet","type":"component","tags":{"group":"TradePosition"},"consumes":["/api/Position"]}]}`
	 applicationModel, err = getApplicationModel(m)
	 if(err != nil) { log.Fatal(err) }  
	 //fmt.Println(applicationModel)
	 //fmt.Println(applicationModel.Id)  
     
     /*
	 router := gin.Default()
	 router.GET("/application", getApplication)
	 router.Run("localhost:8080")
	 */
			
	 router := gin.Default()
	 
	 router.GET("/application", getApplication)
	 router.PUT("/application", putApplication)

	 router.Run("localhost:8080")

}
	
/*	
	func testApplicationUnmarshal(t *testing.T) {
			const m string = `{"id":"FOApp","type":"application","components":[{"id":"Database","type":"component","tags":{"group":"core","type":"database"},"provides":[{"id":"raw data","kind":6}]},{"id":"EventBus","type":"component","tags":{"group":"core"},"commands":{"start":{"type":"javascript","steps":["StartComponent"]},"stop":{"type":"javascript","steps":["StopComponent"]}},"provides":[{"id":"raw events","kind":6}]},{"id":"Cache","type":"component","tags":{"group":"core"},"consumes":["raw events","raw data"]},{"id":"PositionService","type":"component","tags":{"group":"TradePosition"},"provides":[{"id":"/api/Position","object":"Position","kind":2,"protocol":"REST"}],"consumes":["raw events","raw data"]},{"id":"Spreadsheet","type":"component","tags":{"group":"TradePosition"},"consumes":["/api/Position"]}]}`
			var a model.Application
			err := json.Unmarshal([]byte(m), &a)
			if err != nil {
				t.Errorf(`Deserialisation of %v failed`, m)
			}
	}
*/

/*
func main2() {

  	log.SetPrefix("ws: ")
    log.SetFlags(0)

	const m string = `{"id":"FOApp","type":"application","components":[{"id":"Database","type":"component","tags":{"group":"core","type":"database"},"provides":[{"id":"raw data","kind":6}]},{"id":"EventBus","type":"component","tags":{"group":"core"},"commands":{"start":{"type":"javascript","steps":["StartComponent"]},"stop":{"type":"javascript","steps":["StopComponent"]}},"provides":[{"id":"raw events","kind":6}]},{"id":"Cache","type":"component","tags":{"group":"core"},"consumes":["raw events","raw data"]},{"id":"PositionService","type":"component","tags":{"group":"TradePosition"},"provides":[{"id":"/api/Position","object":"Position","kind":2,"protocol":"REST"}],"consumes":["raw events","raw data"]},{"id":"Spreadsheet","type":"component","tags":{"group":"TradePosition"},"consumes":["/api/Position"]}]}`
	var a model.Application
	err := json.Unmarshal([]byte(m), &a)
	if err != nil {
		log.Fatal(`Deserialisation of failed`, m)
	} 
	fmt.Println(a)

	

	 //f := new (model.Application)
	 //fmt.Println(f) 
}
*/