// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/CZL-AI/czlai-go"
	"github.com/CZL-AI/czlai-go/internal/testutil"
	"github.com/CZL-AI/czlai-go/option"
)

func TestUploadImageNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := czlai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	err := client.UploadImage.New(context.TODO(), czlai.UploadImageNewParams{
		Image:      czlai.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		IsToCloud:  czlai.F(int64(0)),
		UploadType: czlai.F(czlai.UploadImageNewParamsUploadType1),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
