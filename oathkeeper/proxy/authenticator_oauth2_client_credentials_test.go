/*
 * Copyright © 2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author       Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright  2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license  	   Apache-2.0
 */

package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/justpark/auth/oathkeeper/helper"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAuthenticatorOAuth2ClientCredentials(t *testing.T) {
	_, err := NewAuthenticatorOAuth2ClientCredentials("")
	require.Error(t, err)
	_, err = NewAuthenticatorOAuth2ClientCredentials("oauth2/token")
	require.Error(t, err)
}

func TestAuthenticatorOAuth2ClientCredentials(t *testing.T) {
	h := httprouter.New()
	h.POST("/oauth2/token", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h := herodot.NewJSONWriter(logrus.New())
		u, p, ok := r.BasicAuth()
		if !ok || u != "client" || p != "secret" {
			h.WriteError(w, r, helper.ErrUnauthorized)
			return
		}
		h.Write(w, r, map[string]interface{}{"access_token": "foo-token"})
	})
	ts := httptest.NewServer(h)

	a, err := NewAuthenticatorOAuth2ClientCredentials(ts.URL + "/oauth2/token")
	require.NoError(t, err)
	//     "client",
	// "secret",
	//,
	//[]string{"foo-scope"},)
	assert.NotEmpty(t, a.GetID())

	authOk := &http.Request{Header: http.Header{}}
	authOk.SetBasicAuth("client", "secret")

	authInvalid := &http.Request{Header: http.Header{}}
	authInvalid.SetBasicAuth("foo", "bar")

	for k, tc := range []struct {
		r             *http.Request
		config        json.RawMessage
		expectErr     error
		expectSession *AuthenticationSession
	}{
		{
			r:         &http.Request{Header: http.Header{}},
			expectErr: ErrAuthenticatorNotResponsible,
		},
		{
			r:         authInvalid,
			expectErr: helper.ErrUnauthorized,
		},
		{
			r:             authOk,
			expectErr:     nil,
			expectSession: &AuthenticationSession{Subject: "client"},
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			session, err := a.Authenticate(tc.r, tc.config, nil)

			if tc.expectErr != nil {
				require.EqualError(t, errors.Cause(err), tc.expectErr.Error())
			} else {
				require.NoError(t, err)
			}

			if tc.expectSession != nil {
				assert.EqualValues(t, tc.expectSession, session)
			}
		})
	}
}
