package comments

import (
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
)

type Model struct {
	Id, Comment string
}

func Comments(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("Content-Type") == "text/json; charset=utf-8" {
		comments, err := json.Marshal(Model{Id: "foo", Comment: "It is comment"})
		if err != nil {
			log.Fatalf(err.Error())
		}
		w.Write([]byte(comments))
		return
	} else {
		w.Header().Set("Content-Type", "text/html")
		//w.Write([]byte("baar"))
	}
	//w.Write([]byte("<h2>foo bar</h2>"))
	io.WriteString(w, "<h2>Hello world </h2>")
	//encoPng(w)
}
func CommentHandler() http.Handler {
	return http.HandlerFunc(Comments)
}

func encoPng(w http.ResponseWriter) {
	const width, height = 256, 256

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	//f, err := os.Create("image.png")
	//if err != nil {
	//	log.Fatal(err)
	//}

	if err := png.Encode(w, img); err != nil {
		//f.Close()
		log.Fatal(err)
	}

	//if err := f.Close(); err != nil {
	//	log.Fatal(err)
	//}
}
