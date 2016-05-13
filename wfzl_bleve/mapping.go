package main

import "github.com/blevesearch/bleve"

func buildIndexMapping() (*bleve.IndexMapping, error) {
	zlmapping := bleve.NewDocumentMapping()
	fulltext := bleve.NewTextFieldMapping()
	fulltext.Name = "fulltext"
	fulltext.Analyzer = "gojieba"

	zlmapping.AddFieldMappingsAt("F_ApplicationNo", fulltext)
	zlmapping.AddFieldMappingsAt("F_PatentName", fulltext)
	zlmapping.AddFieldMappingsAt("F_Inventor", fulltext)
	zlmapping.AddFieldMappingsAt("F_Applicant", fulltext)
	zlmapping.AddFieldMappingsAt("F_AreaCode", fulltext)
	zlmapping.AddFieldMappingsAt("F_ApplicantAddress", fulltext)
	zlmapping.AddFieldMappingsAt("W_WoAppCode", fulltext)
	zlmapping.AddFieldMappingsAt("F_Abstract", fulltext)
	zlmapping.AddFieldMappingsAt("F_Cite", fulltext)
	zlmapping.AddFieldMappingsAt("W_ReprintPatent", fulltext)
	zlmapping.AddFieldMappingsAt("F_SignoryItem", fulltext)
	zlmapping.AddFieldMappingsAt("F_PatentName", fulltext)
	indexMapping.AddDocumentMapping("zl", zlmapping)
	return indexMapping, nil
}
