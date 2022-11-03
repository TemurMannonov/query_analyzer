package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/TemurMannonov/query_analyzer/models"
	mockdb "github.com/TemurMannonov/query_analyzer/storage/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetQueries(t *testing.T) {
	resp := models.GetQueriesResponse{
		Queries: []*models.Query{
			{
				QueryID:       123,
				Query:         "SELECT * FROM blogs",
				Calls:         5,
				TotalExecTime: 0.034,
				MinExecTime:   0.008,
				MaxExecTime:   0.010,
				MeanExecTime:  0.015,
			},
			{
				QueryID:       453,
				Query:         "DELETE FROM blogs WHERE id=$1",
				Calls:         4,
				TotalExecTime: 0.0114,
				MinExecTime:   0.0009,
				MaxExecTime:   0.0030,
				MeanExecTime:  0.0025,
			},
		},
		Count: 2,
	}

	testCases := []struct {
		name          string
		query         string
		buildStubs    func(strg *mockdb.MockDBRepositoryI)
		checkResponse func(t *testing.T, recoder *http.Response)
	}{
		{
			name:  "success case",
			query: "?limit=10&page=1&type=select&sort_by_time=asc",
			buildStubs: func(strg *mockdb.MockDBRepositoryI) {
				strg.EXPECT().
					GetList(&models.GetQueriesRequest{
						Limit:      10,
						Page:       1,
						Type:       "select",
						SortByTime: "asc",
					}).Times(1).Return(&resp, nil)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, http.StatusOK, response.StatusCode)
				requireBodyMatch(t, response.Body, &resp)
			},
		},
		{
			name:  "with incorrect type",
			query: "?limit=10&page=1&type=alter",
			buildStubs: func(strg *mockdb.MockDBRepositoryI) {
				strg.EXPECT().GetList(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, http.StatusBadRequest, response.StatusCode)
			},
		},
		{
			name:  "with incorrect sort",
			query: "?limit=10&page=1&sort_by_time=ascending",
			buildStubs: func(strg *mockdb.MockDBRepositoryI) {
				strg.EXPECT().GetList(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, http.StatusBadRequest, response.StatusCode)
			},
		},
		{
			name:  "with incorrect limit",
			query: "?limit=a&page=1",
			buildStubs: func(strg *mockdb.MockDBRepositoryI) {
				strg.EXPECT().GetList(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, http.StatusBadRequest, response.StatusCode)
			},
		},
		{
			name:  "with incorrect page",
			query: "?limit=10&page=1a",
			buildStubs: func(strg *mockdb.MockDBRepositoryI) {
				strg.EXPECT().GetList(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, http.StatusBadRequest, response.StatusCode)
			},
		},
		{
			name:  "internal server error",
			query: "?limit=10&page=1",
			buildStubs: func(strg *mockdb.MockDBRepositoryI) {
				strg.EXPECT().GetList(gomock.Any()).Times(1).Return(nil, errors.New("db error"))
			},
			checkResponse: func(t *testing.T, response *http.Response) {
				require.Equal(t, http.StatusInternalServerError, response.StatusCode)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			strg := mockdb.NewMockDBRepositoryI(ctrl)
			tc.buildStubs(strg)

			server := newTestServer(t, strg)

			url := fmt.Sprintf("/queries%s", tc.query)
			fmt.Println(url)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			resp, err := server.Router.Test(request, 3)
			assert.NoError(t, err)
			tc.checkResponse(t, resp)
		})
	}
}

func requireBodyMatch(t *testing.T, body io.ReadCloser, resp *models.GetQueriesResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotResponse *models.GetQueriesResponse
	err = json.Unmarshal(data, &gotResponse)
	require.NoError(t, err)
	require.Equal(t, resp, gotResponse)
}
