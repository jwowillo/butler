package source

import (
	"bytes"
	"encoding/xml"
	"path/filepath"

	"github.com/jwowillo/butler/gen"
)

// sourceSitemap returns a gen.Source containing a sitemap for the site at the
// provided URL and with all the pages provided as gen.Sources.
func sourceSitemap(u string, ss []gen.Source) gen.Source {
	return func() (*gen.Page, error) {
		sm := sitemap{}
		sm.Version = "http://www.sitemaps.org/schemas/sitemap/0.9"
		for _, s := range filter(".html", ss) {
			p, err := s()
			if err != nil {
				return nil, err
			}
			if p.Path == "index.html" {
				sm.URLs = append(sm.URLs, url{Loc: u})
			} else if filepath.Base(p.Path) == "index.html" {
				path := filepath.Dir(p.Path) + "/"
				sm.URLs = append(
					sm.URLs,
					url{Loc: u + path},
				)
			} else {
				path := u + "/" + p.Path
				sm.URLs = append(sm.URLs, url{Loc: path})
			}
		}
		buf := bytes.NewBufferString(xml.Header)
		enc := xml.NewEncoder(buf)
		enc.Indent("", "    ")
		if err := enc.Encode(sm); err != nil {
			return nil, err
		}
		return gen.NewPage("sitemap.xml", buf), nil
	}
}

// sitemap is an XML struct which defines what a sitemap looks like in XML.
type sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Version string   `xml:"xmlns,attr"`
	URLs    []url    `xml:"url"`
}

// url is an XML struct which defines what a URL looks like in XML.
type url struct {
	Loc string `xml:"loc"`
}
