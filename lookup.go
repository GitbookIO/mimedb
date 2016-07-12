package mimedb

// ContentType returns the content-type for a given extension, an empty string if non found
func ContentType(extension string) string {
	entry := lookupEntry(extension)
	if entry == nil {
		return ""
	}
	return entry.ContentType
}

// IsCompressible if the extension usually belongs to a file that can be compressed
// this often true for text like files (JSON, HTML, CSS, JS, ...)
func IsCompressible(extension string) bool {
	entry := lookupEntry(extension)
	if entry == nil {
		return false
	}
	return entry.Compressible
}

func lookupEntry(extension string) *mimeEntry {
	// Trim dot
	entry, ok := DB[trimDot(extension)]
	if !ok {
		return nil
	}
	return &entry
}

func trimDot(str string) string {
	if len(str) < 2 {
		return str
	}
	if str[0] == '.' {
		return str[1:]
	}
	return str
}
