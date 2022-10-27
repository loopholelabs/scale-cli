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

package cmd

//var listCmd = &cobra.Command{
//	Use:   "list",
//	Args:  cobra.ExactArgs(0),
//	Short: "list lists all the scale functions that are available",
//	Long:  `List lists all the scale functions that are available.`,
//	Run: func(cmd *cobra.Command, args []string) {
//		logger, _ := config.Init(cmd, false)
//		scaleFuncEntries, err := storage.Default.List()
//		if err != nil {
//			logger.Fatal().Err(err).Msg("error listing scale functions")
//		}
//		middleware := cmd.Flag("middleware").Value.String() == "true"
//		if cmd.Flag("json").Value.String() == "true" {
//			err = json.NewList(scaleFuncEntries, middleware)
//			if err != nil {
//				logger.Fatal().Err(err).Msg("error displaying scale function list as json")
//			}
//		} else {
//			err = ui.NewList(scaleFuncEntries, middleware)
//			if err != nil {
//				logger.Fatal().Err(err).Msg("error displaying scale function list")
//			}
//		}
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(listCmd)
//	listCmd.Flags().BoolP("middleware", "m", false, "only list middleware functions")
//}
