package common

// var AliasModels = []string{"models", "model"}
// var AliasDatasets = []string{"datasets", "dataset", "data"}
// var AliasExperiments = []string{"experiments", "experiment", "expr"}
// var AliasTasks = []string{"tasks", "task"}

var resourceMap = map[string]string{
	"models": "models",
	"model":  "models",

	"datasets": "datasets",
	"dataset":  "datasets",
	"data":     "datasets",

	"experiments": "experiments",
	"experiment":  "experiments",
	"exp":         "experiments",

	"tasks": "tasks",
	"task":  "tasks",
}

func GetResourceType(input string) string {
	result, ok := resourceMap[input]
	if !ok {
		result = input
	}

	return result
}
