/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package auth

import (
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

const (
	DefaultEndpoint = "auth.scale.sh"
	DefaultBasePath = "/"
	OAuthClientID   = "dgmOd1GERZf7l6uLGX1Y"
)

var (
	DefaultScheme = []string{"https"}
)

func Expired(raw string) (bool, error) {
	standardClaims := jwt.Claims{}
	token, err := jwt.ParseSigned(raw)
	if err != nil {
		return false, err
	}
	err = token.UnsafeClaimsWithoutVerification(&standardClaims)
	if err != nil {
		return false, err
	}
	err = standardClaims.Validate(jwt.Expected{
		Time: time.Now().UTC(),
	})
	if err != nil {
		return true, nil
	}
	return false, nil
}
