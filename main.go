package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charleshuang3/subtool/lib/sub"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "subtool",
		Short: "A tool to process subtitle files",
	}

	var inputFile string
	var deeplKeyFile string
	var unwantFile string

	removeDescriptiveCmd := &cobra.Command{
		Use:   "remove-descriptive-subtitle",
		Short: "Remove descriptive subtitles from a file",
		Run: func(cmd *cobra.Command, args []string) {
			err := sub.RemoveDescriptiveSubtitles(inputFile, os.Stdout)
			if err != nil {
				log.Fatalln("Error removing descriptive subtitles: ", err)
			}
		},
	}

	analyzeRepeatCmd := &cobra.Command{
		Use:   "analyze-repeat",
		Short: "Analyze repeat of subtitles",
		Run: func(cmd *cobra.Command, args []string) {
			err := sub.AnalyzeRepeatSubtitles(inputFile, os.Stdout)
			if err != nil {
				log.Fatalln("Error analyzing subtitle repeats:", err)
			}
		},
	}

	removeUnwantCmd := &cobra.Command{
		Use:   "remove-unwant",
		Short: "Remove unwant subtitles from a file",
		Run: func(cmd *cobra.Command, args []string) {
			err := sub.RemoveUnwantSubtitles(inputFile, unwantFile, os.Stdout)
			if err != nil {
				log.Fatalln("Error removing unwant subtitles: ", err)
			}
		},
	}

	translateCmd := &cobra.Command{
		Use:   "translate",
		Short: "Translate subtitles using Deepl",
		Run: func(cmd *cobra.Command, args []string) {
			err := sub.TranslateSubtitles(inputFile, deeplKeyFile, os.Stdout)
			if err != nil {
				log.Fatalln("Error translating subtitles:", err)
			}
		},
	}

	removeDescriptiveCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	removeDescriptiveCmd.MarkFlagRequired("input")

	removeUnwantCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	removeUnwantCmd.MarkFlagRequired("input")
	removeUnwantCmd.Flags().StringVar(&unwantFile, "unwant", "", "File contains unwanted sub")
	removeUnwantCmd.MarkFlagRequired("unwant")

	translateCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	translateCmd.MarkFlagRequired("input")
	translateCmd.Flags().StringVar(&deeplKeyFile, "deepl-key", "", "Deepl key file path")
	translateCmd.MarkFlagRequired("deepl-key")

	rootCmd.AddCommand(removeDescriptiveCmd, analyzeRepeatCmd, removeUnwantCmd, translateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
