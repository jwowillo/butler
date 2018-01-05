const recipes = {
  {{ range $i, $r := . }}
  '{{ $r.Name }}': {
    'path': '{{ $r.Path }}',
    'name': '{{ $r.Name }}',
    'description': '{{ $r.Description }}',
    'ingredients': [
      {{ range $j, $item := $r.Ingredients }}
      {
        'amount': new Fraction(
          {{ $item.Amount.Numerator }},
          {{ $item.Amount.Denominator }}
        ),
        'unit': '{{ $item.Unit }}',
        'item': '{{ $item.Item }}',
        'singularPhrase': '{{ $item.SingularPhrase }}',
        'pluralPhrase': '{{ $item.PluralPhrase }}',
        'fractionalPhrase': '{{ $item.FractionalPhrase }}'
      }{{ if ne (inc $j) (len $r.Ingredients) }},{{ end }}
      {{ end }}
    ],
    'steps': [
      {{ range $j, $item := $r.Steps }}
      '{{ $item }}'{{ if ne (inc $j) (len $r.Steps) }},{{ end }}
      {{ end }}
    ],
    'notes': [
      {{ range $j, $item := $r.Notes }}
      '{{ $item }}'{{ if ne (inc $j) (len $r.Notes) }},{{ end }}
      {{ end }}
    ]
  }{{ if ne (inc $i) (len $) }},{{ end }}
  {{ end }}
};
