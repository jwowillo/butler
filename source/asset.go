package source

import (
	"bytes"
	"path/filepath"

	"github.com/jwowillo/butler/gen"
)

// sourcesAsset returns all gen.Sources for all static assets. This includes
// non-templated files and images.
func sourcesAsset() []gen.Source {
	as, err := AssetDir("asset")
	if err != nil {
		return nil
	}
	ss := make([]gen.Source, 0, len(as))
	for _, a := range as {
		x := a
		ss = append(ss, func() (*gen.Page, error) {
			bs, err := Asset(filepath.Join("asset", x))
			if err != nil {
				return nil, err
			}
			return gen.NewPage("/"+x, bytes.NewBuffer(bs)), nil
		})
	}
	return ss
}
