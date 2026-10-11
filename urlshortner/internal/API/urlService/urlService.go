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
	types.UniqueValueAndLink[uniqueCode] = urlStuctureForComingLink.Link

	// localhost link
	ShortUrl := ("http://localhost:8080/makebig/") + (types.LinkAndUniqueValue[urlStuctureForComingLink.Link])

	w.Header().Set("Content-Type", "application/json")

	sendData := sendandrecive.SendDataSucusses{}
	sendData.StatusCode = http.StatusOK
	sendData.AnyMessage = "adding link and generating the code successfuly"
	sendData.Data = struct {
		Recievelink string
		ShortUrl    string
		Randnum     int
	}{
		Recievelink: urlStuctureForComingLink.Link,
		ShortUrl:    ShortUrl,
		Randnum:     123,
	}

	json.NewEncoder(w).Encode(&sendData)

	fmt.Println(sendData)

}

func GetShortLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// take request data
	var shortLink types.ShortUrlStructure
	shortLink.ShortLink = r.PathValue("code")
	// check with link to key
	getLongLink := types.UniqueValueAndLink[shortLink.ShortLink]
	if getLongLink != "" {
		// redirect the long link
		http.Redirect(w, r, getLongLink, http.StatusPermanentRedirect)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(&getLongLink)
		sendData := sendandrecive.SendDataSucusses{}
		sendData.StatusCode = http.StatusBadRequest
		sendData.AnyMessage = "Code find Matching URL"
		sendData.Data = struct {
			Msg     string
			Randnum int
		}{
			Msg:     "long URL found",
			Randnum: 123,
		}
		json.NewEncoder(w).Encode(&sendData)
	} else {
		sendData := sendandrecive.SendDataFail{}
		sendData.StatusCode = http.StatusBadRequest
		sendData.AnyMessage = "sended code is not valid or wrong"
		sendData.Data = struct {
			Msg     string
			Randnum int
		}{
			Msg:     "Failed to get long URL",
			Randnum: 123,
		}
		json.NewEncoder(w).Encode(&sendData)
		fmt.Println("Short url is not valide or may wrong")
	}
}
