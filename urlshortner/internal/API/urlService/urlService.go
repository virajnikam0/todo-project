package urlservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	sendandrecive "github.com/virajnikam0/todo-project/urlshortner/internal/sendAndReceieve"
	"github.com/virajnikam0/todo-project/urlshortner/internal/types"
)

func GetLink(w http.ResponseWriter, r *http.Request) {
	var urlStuctureForComingLink types.UrlStructure

	// get the link structure
	if err := json.NewDecoder(r.Body).Decode(&urlStuctureForComingLink); err != nil {
		fmt.Println("getting error during the fetching data from frontend : error -> ", err)
	}

	// that unique string
	uniqueCode := types.UniqueStringGenaration()
	types.LinkAndUniqueValue[urlStuctureForComingLink.Link] = uniqueCode

	w.Header().Set("Content-Type", "application/json")

	sendData := sendandrecive.SendDataSucusses{}
	sendData.StatusCode = http.StatusOK
	sendData.AnyMessage = "adding link and generating the code successfuly"
	sendData.Data = struct {
		Recievelink string
		Uniquecode  string
		Randnum     int
	}{
		Recievelink: urlStuctureForComingLink.Link,
		Uniquecode:  types.LinkAndUniqueValue[urlStuctureForComingLink.Link],
		Randnum:     123,
	}

	json.NewEncoder(w).Encode(&sendData)

	fmt.Println(sendData)

}
