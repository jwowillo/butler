const recipes = [
  {{ range $i, $r := . }}
  {
    "path": "{{ $r.Path }}",
    "name": "{{ $r.Name }}",
    "description": "{{ $r.Description }}",
    "ingredients": [
      {{ range $j, $item := $r.Ingredients }}
      "{{ $item }}"{{ if ne (inc $j) (len $r.Ingredients) }},{{ end }}
      {{ end }}
    ],
    "steps": [
      {{ range $j, $item := $r.Steps }}
      "{{ $item }}"{{ if ne (inc $j) (len $r.Steps) }},{{ end }}
      {{ end }}
    ],
    "notes": [
      {{ range $j, $item := $r.Notes }}
      "{{ $item }}"{{ if ne (inc $j) (len $r.Notes) }},{{ end }}
      {{ end }}
    ]
  }{{ if ne (inc $i) (len $) }},{{ end }}
  {{ end }}
];
