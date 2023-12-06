package img

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"

	"golang.org/x/image/webp"
)

// GetDimension get dimension of picture
func GetDimension(fheader *multipart.FileHeader) (width, height int, err error) {
	src, err := fheader.Open()
	if err != nil {
		return 0, 0, err
	}
	defer src.Close()

	decode, _, err := image.DecodeConfig(src)
	if err != nil {
		return 0, 0, err
	}

	return decode.Width, decode.Height, nil
}

// GetFileContentType is return content type
func GetFileContentType(r io.Reader) (string, error) {
	/*
		Only the first 512 bytes are used to sniff the content type.
	*/
	buffer := make([]byte, 512)

	if _, err := r.Read(buffer); err != nil {
		return "", err
	}

	/*
		Use the net/http package's handy DectectContentType function. Always returns a valid

		content-type by returning "application/octet-stream" if no others seemed to match.
	*/
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// GetDimension get dimension of picture
func GetDimensionWebp(fheader *multipart.FileHeader) (width, height int, err error) {
	src, err := fheader.Open()
	if err != nil {
		return 0, 0, err
	}
	defer src.Close()

	decode, err := webp.DecodeConfig(src)
	if err != nil {
		return 0, 0, err
	}

	return decode.Width, decode.Height, nil
}
