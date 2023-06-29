package main

import (
	"encoding/json"
	"net/http"
)

const (
	requestId      = "ae1882ba-f60a-7629-ce1a-6618c482de3e"
	accessor       = "N5O3Ceobw5mvZ9YGRFLsXYPf"
	creationTime   = "1687975599"
	creationTTL    = "768h"
	displayName    = "token"
	entityId       = "0660dce5-4f2c-926a-8b15-158901557d9d"
	expireTime     = "2023-07-31T10:05:56.405304+01:00"
	explicitMaxTtl = "0s"
	id             = "hvs.CAESICO3rMp55tMIChbr0JFuYcezpTZQUmBa2A8SntV6emxHGh4KHGh2cy5UTTZZdzdWbjJWVXlzYXFIUk5TT1ByRnc"
	issueTime      = "2023-06-29T10:05:56.405307+01:00"
	numUses        = 0
	orphan         = true
	path           = "auth/token/create"
	renewable      = true
	ttl            = "767h59m37s"
	tokenType      = "service"
)

func Lookup() http.HandlerFunc {

	type Data struct {
		Accessor       string      `json:"accessor"`
		CreationTime   string      `json:"creation_time"`
		CreationTTL    string      `json:"creation_ttl"`
		DisplayName    string      `json:"display_name"`
		EntityId       string      `json:"entity_id"`
		ExpireTime     string      `json:"expire_time"`
		ExplicitMaxTtl string      `json:"explicit_max_ttl"`
		Id             string      `json:"id"`
		IssueTime      string      `json:"issue_time"`
		Meta           interface{} `json:"meta"`
		NumUses        int         `json:"num_uses"`
		Orphan         bool        `json:"orphan"`
		Path           string      `json:"path"`
		Policies       []string    `json:"policies"`
		Renewable      bool        `json:"renewable"`
		Ttl            string      `json:"ttl"`
		Type           string      `json:"type"`
	}

	type lookupResponse struct {
		RequestId     string      `json:"request_id"`
		LeaseId       string      `json:"lease_id"`
		LeaseDuration int         `json:"lease_duration"`
		Renewable     bool        `json:"renewable"`
		Data          Data        `json:"data"`
		Warnings      interface{} `json:"warnings"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

		resp := lookupResponse{
			RequestId:     "31bdcd4d-41ce-d784-6b11-b93615f93ddc",
			LeaseId:       "",
			LeaseDuration: 0,
			Renewable:     false,
			Data: Data{
				Accessor:       accessor,
				CreationTime:   creationTime,
				CreationTTL:    creationTTL,
				DisplayName:    displayName,
				EntityId:       entityId,
				ExpireTime:     expireTime,
				ExplicitMaxTtl: explicitMaxTtl,
				Id:             id,
				IssueTime:      issueTime,
				Meta:           nil,
				NumUses:        numUses,
				Orphan:         orphan,
				Path:           path,
				Policies:       []string{"default", "policy1", "policy2"},
				Renewable:      renewable,
				Ttl:            ttl,
				Type:           tokenType,
			},
			Warnings: nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
