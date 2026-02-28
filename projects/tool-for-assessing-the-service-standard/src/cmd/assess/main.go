package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/engine"
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/output"
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
	"github.com/spf13/cobra"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	var (
		serviceURL    string
		format        string
		outFile       string
		automatedOnly bool
		verbose       bool
	)

	cmd := &cobra.Command{
		Use:   "assess <service-name>",
		Short: "Assess a government service against the GDS Service Standard",
		Long: `assess evaluates a government service in real time against the 14 points
of the GDS Service Standard and produces a graded report.

Example:
  assess "Apply for a passport" --url https://www.gov.uk/apply-renew-passport`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			serviceName := args[0]

			if serviceURL == "" {
				return fmt.Errorf("--url is required")
			}

			if !strings.HasPrefix(serviceURL, "http://") && !strings.HasPrefix(serviceURL, "https://") {
				serviceURL = "https://" + serviceURL
			}

			opts := engine.Options{
				AutomatedOnly: automatedOnly,
				Verbose:       verbose,
			}

			eng := engine.New()

			if verbose {
				fmt.Fprintf(os.Stderr, "Assessing %q at %s …\n", serviceName, serviceURL)
			}

			result, err := eng.Assess(serviceName, serviceURL, opts)
			if err != nil {
				return fmt.Errorf("assessment failed: %w", err)
			}

			var dest *os.File
			if outFile != "" {
				dest, err = os.Create(outFile)
				if err != nil {
					return fmt.Errorf("creating output file: %w", err)
				}
				defer dest.Close()
			} else {
				dest = os.Stdout
			}

			switch output.Format(format) {
			case output.FormatJSON:
				if err := output.WriteJSON(dest, result); err != nil {
					return fmt.Errorf("writing JSON: %w", err)
				}
			case output.FormatText:
				if err := output.WriteText(dest, result); err != nil {
					return fmt.Errorf("writing text: %w", err)
				}
			default:
				return fmt.Errorf("unknown format %q; choose json or text", format)
			}

			if outFile != "" && verbose {
				fmt.Fprintf(os.Stderr, "Result written to %s\n", outFile)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&serviceURL, "url", "", "Entry-point URL of the service to assess (required)")
	cmd.Flags().StringVarP(&format, "format", "f", "text", "Output format: text or json")
	cmd.Flags().StringVarP(&outFile, "output", "o", "", "Write result to this file instead of stdout")
	cmd.Flags().BoolVar(&automatedOnly, "automated-only", false, "Only include fully automated checks in the report")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print progress to stderr")

	cmd.AddCommand(newSaveResultCmd())

	return cmd
}

// savedResult is used when writing results to a timestamped JSON file.
type savedResult struct {
	*types.AssessmentResult
	SavedAt string `json:"saved_at"`
}

func newSaveResultCmd() *cobra.Command {
	var (
		serviceURL    string
		outDir        string
		automatedOnly bool
	)

	cmd := &cobra.Command{
		Use:   "save <service-name>",
		Short: "Assess a service and save the result to a timestamped JSON file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			serviceName := args[0]

			if serviceURL == "" {
				return fmt.Errorf("--url is required")
			}

			if !strings.HasPrefix(serviceURL, "http://") && !strings.HasPrefix(serviceURL, "https://") {
				serviceURL = "https://" + serviceURL
			}

			eng := engine.New()
			result, err := eng.Assess(serviceName, serviceURL, engine.Options{AutomatedOnly: automatedOnly})
			if err != nil {
				return fmt.Errorf("assessment failed: %w", err)
			}

			if err := os.MkdirAll(outDir, 0755); err != nil {
				return fmt.Errorf("creating output directory: %w", err)
			}

			slug := strings.ToLower(strings.ReplaceAll(serviceName, " ", "-"))
			timestamp := time.Now().UTC().Format("20060102-150405")
			fileName := fmt.Sprintf("%s-%s.json", slug, timestamp)
			filePath := filepath.Join(outDir, fileName)

			f, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("creating result file: %w", err)
			}
			defer f.Close()

			wrapped := savedResult{
				AssessmentResult: result,
				SavedAt:          time.Now().UTC().Format(time.RFC3339),
			}

			enc := json.NewEncoder(f)
			enc.SetIndent("", "  ")
			if err := enc.Encode(wrapped); err != nil {
				return fmt.Errorf("encoding result: %w", err)
			}

			fmt.Fprintf(os.Stdout, "Result saved to: %s\n", filePath)
			fmt.Fprintf(os.Stdout, "Overall grade:  %s (%.1f%%)\n", result.OverallGrade, result.OverallScore)
			return nil
		},
	}

	cmd.Flags().StringVar(&serviceURL, "url", "", "Entry-point URL of the service to assess (required)")
	cmd.Flags().StringVar(&outDir, "dir", "results", "Directory to save result files")
	cmd.Flags().BoolVar(&automatedOnly, "automated-only", false, "Only include fully automated checks")

	return cmd
}
