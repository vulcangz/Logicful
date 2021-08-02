package searcher

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/logicful/service/storage"
	"io/ioutil"
	"strings"
)

type JsonIndex struct {
	Content string `json:"content"`
}

func Index(key string, id string, data interface{}) (bleve.Index, error) {
	files := []string{
		"store",
		"index_meta.json",
	}
	noIndex := false
	var path = "/tmp/" + key
	for i := range files {
		result, err := storage.DownloadToFile("search-indexes", key+files[i], path, files[i])
		if err != nil {
			return nil, err
		}
		if !result {
			noIndex = true
			break
		}
	}
	var index bleve.Index
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	json := JsonIndex{
		Content: string(serialized),
	}
	if noIndex {
		mapping := bleve.NewIndexMapping()
		index, err := bleve.New(path, mapping)
		if err != nil {
			return nil, err
		}
		err = index.Index(key, json)
		if err != nil {
			return nil, err
		}
	} else {
		index, err := bleve.Open(path)
		if err != nil {
			return nil, err
		}
		err = index.Index(id, json)
		if err != nil {
			return nil, err
		}
	}
	for i := range files {
		bytes, err := ioutil.ReadFile(path + "/" + files[i])
		if err != nil {
			return nil, err
		}
		_, err = storage.SetBytes(bytes, key+files[i], "search-indexes", "private")
		if err != nil {
			return nil, err
		}
	}

	return index, nil
}

func Search(key string, q string) ([]string, error) {
	files := []string{
		"store",
		"index_meta.json",
	}
	noIndex := false
	var path = "/tmp/" + key
	for i := range files {
		result, err := storage.DownloadToFile("search-indexes", key+files[i], path, files[i])
		if err != nil {
			return nil, err
		}
		if !result {
			noIndex = true
			break
		}
	}

	if noIndex {
		return nil, nil
	}

	index, err := bleve.Open(path)
	if err != nil {
		return nil, err
	}

	query := bleve.NewFuzzyQuery(strings.ToLower(q))
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	count, _ := index.DocCount()
	println(count)
	searchResults, err := index.Search(search)
	if err != nil {
		return nil, err
	}
	var result []string
	for i := range searchResults.Hits {
		doc := searchResults.Hits[i].Fields
		for s := range doc {
			println(s)
			println(fmt.Sprint(doc[s]))
		}
		result = append(result, fmt.Sprint(doc[""]))
	}
	return result, nil
}
