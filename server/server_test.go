package main

import (
  "github.com/pact-foundation/pact-go/dsl"
  "github.com/pact-foundation/pact-go/types"
  "testing"
)

func TestAllTheThings(t *testing.T) {
  pact := dsl.Pact{
    Port: 6666,
    Consumer: "MyReactFrontEnd",
    Provider: "GolangAPI",
  }
  err := pact.VerifyProvider(types.VerifyRequest{
    ProviderBaseURL: "http://localhost:1323",
    PactURLs: []string{"/home/vagrant/dev/react_for_pact/pacts/myreactfrontend-golangapi.json"},
  })
  if err != nil {
    t.Fatal("Error", err)
  }
}
