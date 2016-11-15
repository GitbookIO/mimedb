package autoload

import (
	"mime"

	"github.com/GitbookIO/mimedb"
)

func init() {
	for ext, mimeType := range mimedb.DB {
		if err := mime.AddExtensionType("."+ext, mimeType.ContentType); err != nil {
			panic(err)
		}
	}
}
