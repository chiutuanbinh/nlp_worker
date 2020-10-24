package nlp

import (
	"nlp_worker/util"
	"time"
)

var nlpDomain = util.Config.NLP.Domain
var nlpTokenizeURL = nlpDomain + "/tokenize"
var nlpNERUrl = nlpDomain + "/ner"
var timeout = 10 * time.Second

const (
	Noun              = "N"
	Verb              = "V"
	None              = "O"
	NounPhrase        = "Np"
	VerbPhrase        = "Vp"
	Punctuation       = "CH"
	BeginNounPhrase   = "B-NP"
	InNounPhrase      = "I-NP"
	BeginVerbPhrase   = "B-VP"
	InVerbPhrase      = "I-VP"
	BeginLocation     = "B-LOC"
	InLocation        = "I-LOC"
	BeginOrganization = "B-ORG"
	InOrganization    = "I-ORG"
	BeginPerson       = "B-PER"
	InPerson          = "I-PER"
)

var empty = NLPResp{}

var nerTypeMapper = make(map[string]string)

const (
	LOCATION     = "LOCATION"
	PERSON       = "PERSON"
	ORGANIZATION = "ORGANIZATION"
	UNKNOWN      = "UNKNOWN"
)
