package session

import "regexp"

var (
	sessionLine      = regexp.MustCompile(`(?i)>[^<]*session[^<]*id[^<]*<`)
	visibleLine      = regexp.MustCompile(`(?i)<table class="[^"]*visible[^"]*">`)
	idColumnLine     = regexp.MustCompile(`(?i)<th class="text-center">[^<]*\bid\b[^<]*<`)
	marketplaceMatch = regexp.MustCompile(`(?i)>[^<]*marketplace[^<]*<`)
	sessionIDLine    = regexp.MustCompile(`<!-- react-text: [0-9]+ -->[1-9][0-9]*`)
	sessionIDExpr    = regexp.MustCompile(`[1-9][0-9]*$`)
	fullyPickedLine  = regexp.MustCompile(`(?i)<!-- react-text: [0-9]+ -->fully.?picked\b`)
	orderNumberLine  = regexp.MustCompile(`[0-9-]+</figcaption>`)
	orderNumberExpr  = regexp.MustCompile(`[0-9-]+`)
)
