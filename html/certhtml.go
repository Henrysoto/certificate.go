package html

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"certificate.go/cert"
)

type HtmlSaver struct {
	OutputDir string
}

func New(outputdir string) (*HtmlSaver, error) {
	var h *HtmlSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return h, err
	}

	h = &HtmlSaver{
		OutputDir: outputdir,
	}
	return h, nil
}

func (h *HtmlSaver) Save(c cert.Cert) error {
	t, err := template.New("Certificate").Parse(tpl)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v.html", c.LabelTitle)
	path := path.Join(h.OutputDir, filename)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, c)
	if err != nil {
		return err
	}

	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

var tpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{.LabelTitle}}</title>
		<style>
			body {
				text-align: center;
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
			}

			h1 {
				font-size: 3em;
			}
		</style>
	</head>
	<body>
		<h1>{{.LabelCompletion}}</h1>
		<h2>{{.LabelPresented}}</h2>
		<h1>{{.Name}}</h1>
		<h2>{{.LabelParticipation}}</h2>
		<p>
			<em>{{.LabelDate}}</em>
		</p>
	</body>
</html>
`
