package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spinnaker/kleat/api/client/config"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"

	"github.com/micnncim/kleat-format/pkg/version"
)

func main() {
	log.SetFlags(0)

	if err := newCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}

type runner struct {
	write   bool
	check   bool
	quiet   bool
	version bool
}

func newCommand() *cobra.Command {
	r := &runner{}

	cmd := &cobra.Command{
		Use: "kleat-format /path/to/halconfig",
		Run: func(cmd *cobra.Command, args []string) {
			if r.version {
				fmt.Printf("%s (%s)\n", version.Version, version.Revision)
				return
			}

			if len(args) == 0 {
				log.Fatal("accepts 1 arg(s), received 0")
			}

			halPath := args[0]
			if err := r.run(halPath); err != nil {
				if !r.quiet {
					log.Println(err)
				}
				os.Exit(1)
			}
		},
	}

	cmd.Flags().BoolVarP(&r.write, "write", "w", false, "If true, write result to source halconfig instead of stdout")
	cmd.Flags().BoolVar(&r.check, "check", false, "If true, only check whether there is diff between source halconfig and formatted one")
	cmd.Flags().BoolVarP(&r.quiet, "quiet", "q", false, "If true, suppress printing logs")
	cmd.Flags().BoolVar(&r.version, "version", false, "If true, print version information")

	return cmd
}

func (r *runner) run(halPath string) error {
	data, err := ioutil.ReadFile(halPath)
	if err != nil {
		return err
	}

	newData, err := format(data)
	if err != nil {
		return err
	}

	switch {
	case r.write:
		if err := ioutil.WriteFile(halPath, newData, 0666); err != nil {
			return err
		}

	case r.check:
		if string(data) != string(newData) {
			return fmt.Errorf("%s is not formatted", halPath)
		}

	default:
		fmt.Printf(string(data))
	}

	return nil
}

func format(data []byte) ([]byte, error) {
	hal := &config.Hal{}
	if err := unmarshalProto(data, hal); err != nil {
		return nil, err
	}

	return marshalProto(hal)
}

func marshalProto(m proto.Message) ([]byte, error) {
	json, err := protojson.Marshal(m)
	if err != nil {
		return nil, err
	}

	return yaml.JSONToYAML(json)
}

func unmarshalProto(b []byte, m proto.Message) error {
	json, err := yaml.YAMLToJSON(b)
	if err != nil {
		return err
	}

	return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(json, m)
}
