package main

import (
	"fmt"
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
	var repeatFile string

	removeDescriptiveCmd := &cobra.Command{
		Use:   "remove-descriptive-subtitle",
		Short: "Remove descriptive subtitles from a file",
		Run: func(cmd *cobra.Command, args []string) {
			err := sub.RemoveDescriptiveSubtitles(inputFile)
			if err != nil {
				fmt.Println("Error removing descriptive subtitles:", err)
			} else {
				fmt.Println("Successfully cleaned descriptive subtitles from:", inputFile)
			}
		},
	}

	analyzeRepeatCmd := &cobra.Command{
		Use:   "analyze-repeat",
		Short: "Analyze repeat of subtitles",
		Run: func(cmd *cobra.Command, args []string) {
			err := sub.AnalyzeRepeatSubtitles(inputFile, os.Stdout)
			if err != nil {
				fmt.Println("Error analyzing subtitle repeats:", err)
			}
		},
	}

	removeUselessCmd := &cobra.Command{
		Use:   "remove-useless-subtitle",
		Short: "Remove useless subtitles from a file",
		Run: func(cmd *cobra.Command, args []string) {
			err := sub.RemoveUselessSubtitles(inputFile, repeatFile)
			if err != nil {
				fmt.Println("Error removing useless subtitles:", err)
			} else {
				fmt.Println("Successfully removed useless subtitles from:", inputFile)
			}
		},
	}

	translateCmd := &cobra.Command{
		Use:   "translate",
		Short: "Translate subtitles using Deepl",
		Run: func(cmd *cobra.Command, args []string) {
			err := sub.TranslateSubtitles(inputFile, deeplKeyFile)
			if err != nil {
				fmt.Println("Error translating subtitles:", err)
			} else {
				fmt.Println("Successfully translated subtitles in:", inputFile)
			}
		},
	}

	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	rootCmd.MarkPersistentFlagRequired("input")

	removeUselessCmd.Flags().StringVar(&repeatFile, "rm", "", "Repeat file path")
	removeUselessCmd.MarkFlagRequired("rm")

	translateCmd.Flags().StringVar(&deeplKeyFile, "deepl-key", "", "Deepl key file path")
	translateCmd.MarkFlagRequired("deepl-key")

	rootCmd.AddCommand(removeDescriptiveCmd, analyzeRepeatCmd, removeUselessCmd, translateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
