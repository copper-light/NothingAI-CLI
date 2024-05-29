package cmd

import (
	"errors"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
)

func createResource(resourceType string, name string, flags *pflag.FlagSet) {
	resourceType = common.GetResourceType(resourceType)

	options := make(map[string]string)
	flags.Visit(func(f *pflag.Flag) {
		options[strings.Replace(f.Name, "-", "_", -1)] = f.Value.String()
	})

	body := map[string]any{}
	body["name"] = name

	for k, v := range options {
		body[k] = v
	}

	id, err := common.CreateResource(resourceType, body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else if id == -1 {
		fmt.Printf("Failed to create %s\n", resourceType)
	} else {
		fmt.Printf("\"%v\" %v is created\n", name, resourceType)
	}
}

var createCmd = &cobra.Command{
	Use:     "create [RESOURCE_TYPE] [NAME]",
	Short:   "Create a resource",
	GroupID: "basic",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a resource type and a name")
		} else if len(args) < 2 {
			return errors.New("requires a resource type")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]
		name := args[1]
		createResource(resourceType, name, cmd.Flags())
	},
}

var createModel = &cobra.Command{
	Use:     "model [MODEL_NAME] [flags]",
	Short:   "Create a model",
	Aliases: []string{"models"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createResource("model", name, cmd.Flags())
	},
}

var createDataset = &cobra.Command{
	Use:     "dataset [DATASET_NAME] [flags]",
	Short:   "Create a dataset",
	Aliases: []string{"data"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createResource("dataset", name, cmd.Flags())
	},
}

var createExperiment = &cobra.Command{
	Use:     "experiment [DATASET_NAME] [flags]",
	Short:   "Create a experiment",
	Aliases: []string{"exp"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createResource("experiment", name, cmd.Flags())
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.AddCommand(createModel)
	createModel.Flags().StringP("description", "", "", "Model description")
	createModel.Flags().StringP("model-type", "t", "", "Model type")
	createModel.Flags().StringP("source-type", "s", "", "Type storage of source file")
	createModel.Flags().StringP("python-version", "p", "", "Python-version")

	createCmd.AddCommand(createDataset)
	createDataset.Flags().StringP("description", "", "", "Dateset description")
	createDataset.Flags().StringP("dataset-type", "t", "", "Dateset type")

	createCmd.AddCommand(createExperiment)
	createExperiment.Flags().StringP("description", "", "", "Experiment description")
	createExperiment.Flags().StringP("dataset", "d", "", "Dateset id")
	createExperiment.Flags().StringP("model", "m", "", "model id")
	_ = createExperiment.MarkFlagRequired("model")
	_ = createExperiment.MarkFlagRequired("dataset")
}
