package ws_test

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"services/types"
	"strings"
	"testing"
	"time"
)

var addr = "127.0.0.1:8000"
var addressFile = "/Users/paul/GolandProjects/LFS-Services/testdata/2019_ADDRESS_FILE_FOR_CASPA.csv"
var remoteURL = "http://localhost:8000/imports/address"

func TestWS(t *testing.T) {

	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Error().Err(err).Msg("dial error")
		t.FailNow()
	}

	var client = &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	values := map[string]io.Reader{
		"lfsFile":  mustOpen(addressFile),
		"fileName": strings.NewReader("addressFile"),
	}

	err = Upload(client, remoteURL, values)
	if err != nil {
		panic(err)
	}

	message := types.WSMessage{
		Filename:     "addressFile",
		Percentage:   0,
		Status:       0,
		ErrorMessage: "",
	}

	c.EnableWriteCompression(true)

	for {
		err = c.WriteJSON(&message)
		if err != nil {
			log.Error().Err(err).Msg("write error")
			t.FailNow()
		}

		err = c.ReadJSON(&message)
		if err != nil {
			log.Error().Err(err).Msg("write error")
			t.FailNow()
		}

		log.Printf("Received status response: %d percentage: %f", message.Status, message.Percentage)

		if message.Status == types.UploadError {
			log.Printf("received an error from upload status: %s", message.ErrorMessage)
			break
		}
		if message.Status == types.UploadFinished {
			log.Printf("Done")
			break
		}
		time.Sleep(2 * time.Second)
	}

	log.Info().
		Str("fileName", message.Filename).
		Float64("percentage", message.Percentage).
		Int("status", message.Status).
		Msg("recieved message")

	_ = c.Close()

}

func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}

		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}
