/*
	Copyright 2023 Loophole Labs

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

package function

//
//func ExportCmd(ch *cmdutil.Helper) *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "export <function> <output>",
//		Args:  cobra.ExactArgs(2),
//		Short: "export a compiled scale function to the given output path",
//		RunE: func(cmd *cobra.Command, args []string) error {
//			name := args[0]
//			output := args[1]
//			names := strings.Split(name, ":")
//			if len(names) != 2 {
//				name = fmt.Sprintf("%s:latest", name)
//			}
//
//			destination, err := storage.Default.Copy(name, output)
//			if err != nil {
//				return fmt.Errorf("failed to export scale function %s to %s: %w", name, destination, err)
//			}
//
//			if ch.Printer.Format() == printer.Human {
//				ch.Printer.Printf("Exported scale function %s to %s\n", printer.BoldGreen(name), printer.BoldBlue(destination))
//				return nil
//			}
//
//			return ch.Printer.PrintResource(map[string]string{
//				"destination": destination,
//			})
//		},
//	}
//
//	return cmd
//}
