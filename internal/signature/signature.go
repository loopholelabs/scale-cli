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

package signature

import (
	"context"
	"errors"
	"github.com/loopholelabs/scale-cli/pkg/client"
	"github.com/loopholelabs/scale-cli/pkg/client/models"
	"github.com/loopholelabs/scale-cli/pkg/client/registry"
	"github.com/loopholelabs/scale/scalefile"
	"strings"
)

func GetRemoteGoSignature(client *client.ScaleAPIV1, ctx context.Context, namespace string, name string, version string) (*scalefile.Dependency, error) {
	var payload *models.ModelsGetSignatureResponse
	if namespace != "" {
		res, err := client.Registry.GetRegistrySignatureNamespaceNameVersion(registry.NewGetRegistrySignatureNamespaceNameVersionParamsWithContext(ctx).WithNamespace(namespace).WithName(name).WithVersion(version))
		if err != nil {
			return nil, err
		}
		payload = res.GetPayload()
	} else {
		res, err := client.Registry.GetRegistrySignatureNameVersion(registry.NewGetRegistrySignatureNameVersionParamsWithContext(ctx).WithName(name).WithVersion(version))
		if err != nil {
			return nil, err
		}
		payload = res.GetPayload()
	}

	if payload.Go == "" {
		return nil, errors.New("no published version found for go signature")
	}

	dependencyString := strings.Split(payload.Go, "@")

	return &scalefile.Dependency{
		Name:    dependencyString[0],
		Version: dependencyString[1],
	}, nil
}
