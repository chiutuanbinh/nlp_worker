package nlp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	nerTypeMapper["LOC"] = LOCATION
	nerTypeMapper["PER"] = PERSON
	nerTypeMapper["ORG"] = ORGANIZATION
}

func Tokenize(input string) []string {
	request, err := http.NewRequest("POST", nlpTokenizeURL, bytes.NewBuffer([]byte(input)))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println(string(body))
	var res []string
	json.Unmarshal([]byte(body), &res)
	log.Printf("%+v\n", res[1])
	return res

}

func NLPExtract(input string) NLPResp {
	request, err := http.NewRequest("POST", nlpNERUrl, bytes.NewBuffer([]byte(input)))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
		return empty
	}

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return empty
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return empty
	}

	var nlpResp = NLPResp{Phrases: make([]PhrasesT, 0), NamedEntities: make([]NamedEntitiesT, 0)}
	var nerBuffer bytes.Buffer
	var lastNeType = "0"
	var ok bool
	var res [][]string

	json.Unmarshal([]byte(body), &res)
	// log.Printf("%v\n", nlpResp)
	for i := range res {
		entry := res[i]
		if entry[1] == NounPhrase {
			nlpResp.Phrases = append(nlpResp.Phrases, PhrasesT{Type: NounPhrase, Text: entry[0]})
		}
		firstNERTagChar := entry[3][0]
		if nerBuffer.Len() == 0 {
			if firstNERTagChar == 'B' {
				nerBuffer.WriteString(entry[0])
				lastNeType, ok = nerTypeMapper[entry[3][2:]]
				// log.Println(entry[3][2:])
				if !ok {
					lastNeType = UNKNOWN
				}
			}
		} else {
			if firstNERTagChar == 'B' {
				nlpResp.NamedEntities = append(nlpResp.NamedEntities, NamedEntitiesT{Text: nerBuffer.String(), Type: lastNeType})
				nerBuffer.Reset()
				nerBuffer.WriteString(entry[0])
				lastNeType, ok = nerTypeMapper[entry[3][2:]]
				if !ok {
					lastNeType = UNKNOWN
				}
			} else if firstNERTagChar == 'I' {
				nerBuffer.WriteString(" " + entry[0])
			}
		}

	}
	if nerBuffer.Len() != 0 {
		nlpResp.NamedEntities = append(nlpResp.NamedEntities, NamedEntitiesT{Text: nerBuffer.String(), Type: lastNeType})
	}
	return nlpResp
}
