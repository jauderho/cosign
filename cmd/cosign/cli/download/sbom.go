//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package download

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/peterbourgon/ff/v3/ffcli"

	"github.com/sigstore/cosign/cmd/cosign/cli/options"
	"github.com/sigstore/cosign/internal/oci/remote"
)

func SBOM() *ffcli.Command {
	var (
		flagset = flag.NewFlagSet("cosign download sbom", flag.ExitOnError)
		regOpts options.RegistryOpts
	)
	options.ApplyRegistryFlags(&regOpts, flagset)
	return &ffcli.Command{
		Name:       "sbom",
		ShortUsage: "cosign download sbom <image uri>",
		ShortHelp:  "Download SBOMs from the supplied container image",
		FlagSet:    flagset,
		Exec: func(ctx context.Context, args []string) error {
			if len(args) != 1 {
				return flag.ErrHelp
			}
			_, err := SBOMCmd(ctx, regOpts, args[0], os.Stdout)
			return err
		},
	}
}

func SBOMCmd(ctx context.Context, regOpts options.RegistryOpts, imageRef string, out io.Writer) ([]string, error) {
	ref, err := name.ParseReference(imageRef)
	if err != nil {
		return nil, err
	}

	remoteOpts := regOpts.GetRegistryClientOpts(ctx)

	dstRef, err := remote.SBOMTag(ref, remote.WithRemoteOptions(remoteOpts...))
	if err != nil {
		return nil, err
	}
	img, err := remote.Signatures(dstRef, remote.WithRemoteOptions(remoteOpts...))
	if err != nil {
		return nil, err
	}
	sigs, err := img.Get()
	if err != nil {
		return nil, err
	}
	sboms := make([]string, 0, len(sigs))
	for _, l := range sigs {
		mt, err := l.MediaType()
		if err != nil {
			return nil, err
		}
		fmt.Fprintf(os.Stderr, "Found SBOM of media type: %s\n", mt)
		sbom, err := l.Payload()
		if err != nil {
			return nil, err
		}
		sboms = append(sboms, string(sbom))
		fmt.Fprintln(out, string(sbom))
	}
	return sboms, nil
}
