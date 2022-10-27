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

//var buildCmd = &cobra.Command{
//	Use:   "build",
//	Short: "build a scale function from a scalefile using the scale build service",
//	Long: `Build a scale function from a scalefile using the scale build service.
//
//The scalefile is a YAML file that describes the scale function and its dependencies.
//The scale build service will build the scale function and return the compiled module.`,
//	Run: func(cmd *cobra.Command, args []string) {
//		logger, t := config.Init(cmd, true)
//
//		scalefilePath := cmd.Flag("scalefile").Value.String()
//		if scalefilePath == "" {
//			logger.Fatal().Msg("scalefile path is required")
//		}
//		logger.Debug().Msgf("build called with scalefile '%s'", scalefilePath)
//		scaleFile, err := scalefile.Read(scalefilePath)
//		if err != nil {
//			logger.Fatal().Err(err).Msg("error reading scalefile")
//		}
//
//		directory := path.Dir(scalefilePath)
//		sourcePath := path.Join(directory, scaleFile.Source)
//
//		start := time.Now()
//		source, err := os.ReadFile(sourcePath)
//		if err != nil {
//			logger.Fatal().Err(err).Msgf("error while reading scale function source %s", sourcePath)
//		}
//		logger.Debug().Msgf("read scale function source %s in %s", sourcePath, time.Since(start))
//
//		scaleFunc := build.Build(source, t["access_token"], scaleFile, new(tls.Config), logger)
//		name := cmd.Flag("name").Value.String()
//		if name != "" {
//			scaleFunc.ScaleFile.Name = name
//			names := strings.Split(name, ":")
//			if len(names) == 2 {
//				scaleFunc.ScaleFile.Name = names[0]
//				scaleFunc.Tag = names[1]
//				name = fmt.Sprintf("%s:%s", scaleFunc.ScaleFile.Name, scaleFunc.Tag)
//			} else {
//				scaleFunc.Tag = "latest"
//				name = fmt.Sprintf("%s:%s", name, scaleFunc.Tag)
//			}
//			err = storage.Default.Put(name, scaleFunc)
//		} else {
//			err = storage.Default.Put(fmt.Sprintf("%s:%s", scaleFunc.ScaleFile.Name, scaleFunc.Tag), scaleFunc)
//		}
//		if err != nil {
//			logger.Fatal().Err(err).Msg("error while storing scale function")
//		}
//
//		logger.Info().Msg("Scale Function Compilation Completed")
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(buildCmd)
//	buildCmd.Flags().StringP("scalefile", "s", "scalefile", "the scalefile to use")
//	buildCmd.Flags().StringP("name", "n", "", "the (optional) name of this scale function")
//}
