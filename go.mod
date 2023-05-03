module github.com/sigstore/cosign

go 1.16

require (
	cloud.google.com/go/storage v1.29.0
	cuelang.org/go v0.4.1
	github.com/ThalesIgnite/crypto11 v1.2.5
	github.com/cyberphone/json-canonicalization v0.0.0-20220623050100-57a0ce2678a7
	github.com/go-openapi/runtime v0.26.0
	github.com/go-openapi/strfmt v0.21.7
	github.com/go-openapi/swag v0.22.3
	github.com/go-piv/piv-go v1.11.0
	github.com/google/certificate-transparency-go v1.1.2
	github.com/google/go-cmp v0.5.9
	github.com/google/go-containerregistry v0.14.0
	github.com/google/go-containerregistry/pkg/authn/k8schain v0.0.0-20220114205711-890d5b362eb8
	github.com/google/go-github/v42 v42.0.0
	github.com/in-toto/in-toto-golang v0.8.0
	github.com/manifoldco/promptui v0.9.0
	github.com/miekg/pkcs11 v1.1.1
	github.com/open-policy-agent/opa v0.35.0
	github.com/pkg/errors v0.9.1
	github.com/secure-systems-lab/go-securesystemslib v0.5.0
	github.com/sigstore/fulcio v0.1.2-0.20220114150912-86a2036f9bc7
	github.com/sigstore/rekor v1.1.1
	github.com/sigstore/sigstore v1.6.3
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.2
	github.com/theupdateframework/go-tuf v0.5.2
	github.com/xanzy/go-gitlab v0.54.3
	golang.org/x/oauth2 v0.7.0
	golang.org/x/term v0.7.0
	google.golang.org/api v0.119.0
	k8s.io/api v0.26.1
	k8s.io/apimachinery v0.26.1
	k8s.io/client-go v0.26.1
	k8s.io/utils v0.0.0-20221128185143-99ec85e7a448
	knative.dev/pkg v0.0.0-20220114141842-0a429cba1c73
)

require (
	github.com/bytecodealliance/wasmtime-go v0.32.0 // indirect
	github.com/spiffe/go-spiffe/v2 v2.1.3
	go.opentelemetry.io/contrib v1.3.0 // indirect
)
