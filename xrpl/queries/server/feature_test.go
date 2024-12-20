package server

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/queries/server/types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
)

func TestFeatureAllResponse(t *testing.T) {
	r := FeatureAllResponse{
		Features: map[string]types.FeatureStatus{
			"feature1": {Enabled: true, Name: "feature1", Supported: true},
			"feature2": {Enabled: false, Name: "feature2", Supported: false},
		},
	}

	s := `{
	"features": {
		"feature1": {
			"enabled": true,
			"name": "feature1",
			"supported": true
		},
		"feature2": {
			"enabled": false,
			"name": "feature2",
			"supported": false
		}
	}
}`

	if err := testutil.SerializeAndDeserialize(t, r, s); err != nil {
		t.Fatal(err)
	}
}

func TestFeatureOneRequest(t *testing.T) {
	r := FeatureOneRequest{
		Feature: "feature1",
	}

	s := `{
	"feature": "feature1"
}`

	if err := testutil.Serialize(t, r, s); err != nil {
		t.Fatal(err)
	}
}

func TestFeatureOneResponse(t *testing.T) {
	r := FeatureResponse{
		"feature1": {Enabled: true, Name: "feature1", Supported: true},
	}

	s := `{
	"feature1": {
		"enabled": true,
		"name": "feature1",
		"supported": true
	}
}`

	if err := testutil.SerializeAndDeserialize(t, r, s); err != nil {
		t.Fatal(err)
	}
}
