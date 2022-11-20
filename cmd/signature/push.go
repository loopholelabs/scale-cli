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
	"errors"
	"fmt"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/client/models"
	"github.com/loopholelabs/scale-cli/pkg/client/registry"
	"github.com/loopholelabs/scale/signature"
	"github.com/spf13/cobra"
	"path"
)

func PushCmd(ch *cmdutil.Helper) *cobra.Command {
	var namespace string
	var directory string
	var public bool

	cmd := &cobra.Command{
		Use:     "push <signature> [flags]",
		Args:    cobra.ExactArgs(1),
		Short:   "push a scale signature to the scale registry",
		PreRunE: cmdutil.CheckAuthentication(ch.Config),
		RunE: func(cmd *cobra.Command, args []string) error {
			signatureName := args[0]

			ctx := cmd.Context()
			client, err := ch.Client()
			if err != nil {
				return err
			}

			end := ch.Printer.PrintProgress(fmt.Sprintf("Pushing Signature %s...", signatureName))
			signatureDefinition, err := signature.Read(path.Join(directory, signatureName, "signature.yaml"))
			if err != nil {
				end()
				return fmt.Errorf("failed to read signature definition: %w", err)
			}
			if namespace == "" {
				res, err := client.Registry.PostRegistrySignature(registry.NewPostRegistrySignatureParamsWithContext(ctx).WithSignature(&models.ModelsCreateSignatureRequest{
					Name:    signatureName,
					Version: signatureDefinition.Version,
					Published: &models.ModelsPublishedVersions{
						Go: &models.ModelsPublishedVersion{
							Name:    signatureDefinition.PublishedVersions.Go.Name,
							Version: signatureDefinition.PublishedVersions.Go.Version,
						},
					},
					Public: public,
				}))
				end()
				if err != nil {
					return fmt.Errorf("failed to create signature: %w", err)
				}

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Successfully published scale signature %s/%s:%s\n", printer.BoldGreen(res.Payload.Namespace), printer.BoldGreen(res.Payload.Name), printer.BoldGreen(res.Payload.Version))
					return nil
				}
				return ch.Printer.PrintResource(map[string]string{
					"Name":      res.Payload.Name,
					"Namespace": res.Payload.Namespace,
					"Version":   res.Payload.Version,
					"Public":    fmt.Sprintf("%t", res.Payload.Public),
					"Go":        fmt.Sprintf("%s@%s", res.Payload.Published.Go.Name, res.Payload.Published.Go.Version),
				})
			} else {
				end()
				return errors.New("custom namespaces are not supported yet")
			}
		},
	}

	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "the namespace to push the signature to (defaults to the user's namespace)")
	cmd.Flags().StringVarP(&directory, "directory", "d", "signature", "the directory where the scale signatures are located")
	cmd.Flags().BoolVarP(&public, "public", "p", false, "make the signature public")
	return cmd
}
