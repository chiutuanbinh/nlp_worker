package nlp

type NamedEntitiesT struct {
	Text string
	Type string
}

type PhrasesT struct {
	Text string
	Type string
}

type NLPResp struct {
	Phrases       []PhrasesT
	NamedEntities []NamedEntitiesT
}
