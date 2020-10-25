package main

import (
	"binhct/common/xtype"
	"bytes"
	"log"
	"nlp_worker/mongodb"
	"nlp_worker/nlp"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func nlpExtractor(wg *sync.WaitGroup, jobs <-chan string) {
	defer wg.Done()
	for job := range jobs {
		if job == "DONE" {
			break
		}
		article, err := mongodb.GetByID(job)
		if err != nil {
			continue
		}
		content := bytes.Buffer{}
		for _, c := range article.Content.Parts {

			var cx xtype.Paragraph
			cx.Content = c.(primitive.D)[0].Value.(string)
			// log.Printf("%+v\n", cx)
			content.WriteString(cx.Content)
		}
		if article.Nlp.NamedEntities == nil {
			article.Nlp.NamedEntities = make(map[string][]string)
		} else {
			continue
		}
		if article.Nlp.Phrases == nil {
			article.Nlp.Phrases = make(map[string][]string)
		} else {
			continue
		}
		nlpResp := nlp.NLPExtract(content.String())
		ne := article.Nlp.NamedEntities
		for _, n := range nlpResp.NamedEntities {
			_, ok := ne[n.Type]
			if !ok {
				ne[n.Type] = make([]string, 0)
			}
			skip := false
			for _, nex := range ne[n.Type] {
				if nex == n.Text {
					skip = true
				}

			}
			if !skip {
				ne[n.Type] = append(ne[n.Type], n.Text)
			}

		}

		phr := article.Nlp.Phrases
		for _, ph := range nlpResp.Phrases {
			if _, ok := phr[ph.Type]; !ok {
				phr[ph.Type] = make([]string, 0)
			}
			skip := false
			for _, phx := range phr[ph.Type] {
				if phx == ph.Text {
					skip = true
					break
				}
			}
			if !skip {
				phr[ph.Type] = append(phr[ph.Type], ph.Text)
			}
		}

		// log.Println(article)
		mongodb.Update(job, article)
		log.Println("FIN ", job)
	}
}
