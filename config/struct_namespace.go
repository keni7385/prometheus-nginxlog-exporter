package config

import (
	"errors"
	"sort"
)

// NamespaceConfig is a struct describing single metric namespaces
type NamespaceConfig struct {
	Name           string            `hcl:",key"`
	SourceFiles    []string          `hcl:"source_files" yaml:"source_files"`
	Format         string            `hcl:"format"`
	Labels         map[string]string `hcl:"labels"`
	RelabelConfigs []RelabelConfig   `hcl:"relabel" yaml:"relabel_configs"`

	OrderedLabelNames  []string
	OrderedLabelValues []string

	UpstreamSecondsHistBucket  []float64 `hcl:"upstream_seconds_hist_bucket" yaml:"upstream_seconds_hist_bucket"`
	ResponseSecondsHistBucket  []float64 `hcl:"response_seconds_hist_bucket" yaml:"response_seconds_hist_bucket"`
}

// StabilityWarnings tests if the NamespaceConfig uses any configuration settings
// that are not yet declared "stable"
func (c *NamespaceConfig) StabilityWarnings() error {
	if len(c.RelabelConfigs) > 0 {
		return errors.New("you are using the 'relabel' configuration parameter")
	}

	return nil
}

// MustCompile compiles the configuration (mostly regular expressions that are used
// in configuration variables) for later use
func (c *NamespaceConfig) MustCompile() {
	err := c.Compile()
	if err != nil {
		panic(err)
	}
}

// Compile compiles the configuration (mostly regular expressions that are used
// in configuration variables) for later use
func (c *NamespaceConfig) Compile() error {
	for i := range c.RelabelConfigs {
		if err := c.RelabelConfigs[i].Compile(); err != nil {
			return nil
		}
	}

	c.OrderLabels()

	return nil
}

// OrderLabels builds two lists of label keys and values, ordered by label name
func (c *NamespaceConfig) OrderLabels() {
	keys := make([]string, 0, len(c.Labels))
	values := make([]string, len(c.Labels))

	for k := range c.Labels {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for i, k := range keys {
		values[i] = c.Labels[k]
	}

	c.OrderedLabelNames = keys
	c.OrderedLabelValues = values
}
