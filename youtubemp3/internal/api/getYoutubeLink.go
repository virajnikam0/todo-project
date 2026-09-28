package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/virajnikam0/todo-project/youtubemp3/internal/types"
)

func GetYoutubeLink(w http.ResponseWriter, r *http.Request){

	var userYoutubeLink *url.URL

	// verify the request structure
	ytStructure := types.YoutubeStructure{}
	if err := json.NewDecoder(r.Body).Decode(&ytStructure); err != nil {
		fmt.Println("While decoding request body " , err)
	}

	// verify the link is good
	// if ytUrl,err := url.ParseRequestURI(ytStructure.YoutubeLink); err == nil {
	// 	fmt.Println("link is not valid ", err)
	// 	userYoutubeLink = ytUrl 
	// }	

	// download the youtube mp4 and convert to the mp3 format

	if ConversionFormat(ytStructure.YoutubeLink) {
		http.Error(w,"error during conversion",http.StatusMethodNotAllowed)
	}

	fmt.Println("user link after verification" , userYoutubeLink.Path)
	


}

func ConversionFormat(userYoutubeLink string)(bool){
	
	doesVideoDownloadInLocal := downloadVideo(userYoutubeLink)
	doesItConvertMp4ToMp3 := convertToMP3( "sources/video.mp4",
    "sources/audio.mp3",)

	if doesItConvertMp4ToMp3 && doesVideoDownloadInLocal {
		return true
	}
	return false

}

func downloadVideo(link string)(bool){
	cmd := exec.Command(
		"../config/yt-dlp.exe",
		"-f","bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]",
		"--merge-output-format", "mp4",
   		 "-o", "C://Users//shubh//Documents//VN_study//youtubemp3//sources//video.mp4",
		 link,
	)

	fmt.Println("my link ",link)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run();err!=nil{
		fmt.Println("Error while downloading the video")	
		fmt.Println("and the error is , ", err)
		return false
	}
		return true
}

func convertToMP3(input string, output string) (bool) {
	cmd := exec.Command(
		"ffmpeg",
		"-i", input,
		"-vn",
		"-codec:a", "libmp3lame",
		"-q:a", "0",
		"-y",
		output,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run();err!=nil{
		fmt.Println("Error while downloading the video")
		return false
	}
		return true
}
