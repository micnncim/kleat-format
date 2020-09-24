package main

import (
	"fmt"
	"io/ioutil"
	"log"

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

type options struct {
	write   bool
	check   bool
	version bool
}

func newCommand() *cobra.Command {
	opt := &options{}

	cmd := &cobra.Command{
		Use: "kleat-format PATH_TO_HALCONFIG",
		Run: func(cmd *cobra.Command, args []string) {
			if opt.version {
				fmt.Printf("%s (%s)\n", version.Version, version.Revision)
				return
			}

			if len(args) == 0 {
				log.Fatal("accepts 1 arg(s), received 0")
			}

			halPath := args[0]
			if err := opt.run(halPath); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&opt.write, "write", "w", false, "If true, write result to source halconfig instead of stdout")
	cmd.Flags().BoolVar(&opt.check, "check", false, "If true, only check whether there is diff between source halconfig and formatted one")
	cmd.Flags().BoolVar(&opt.version, "version", false, "If true, print version information")

	return cmd
}

func (o *options) run(halPath string) error {
	data, err := ioutil.ReadFile(halPath)
	if err != nil {
		log.Fatal(err)
	}

	newData, err := format(data)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case o.write:
		if err := ioutil.WriteFile(halPath, newData, 0666); err != nil {
			return err
		}

	case o.check:
		if string(data) != string(newData) {
			return fmt.Errorf("%s is not formatted", halPath)
		}

	default:
		fmt.Println(string(data))
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
