# monorepo

This is my personal playground project to try out new things and experiment.

If you come across this don't take it too seriously.

## Services

- products
- reviews



## Verify container images with cosign
For simplicity we are using the cosign public signatures. But its really cool.

We are able to not only sign the image but also add attestation that the image has been scanned aswell.
```bash
# verifies the image has been signed by cosign
COSIGN_EXPERIMENTAL=1 cosign verify ghcr.io/braderz/monorepo/reviews:pr-42

# verifies the scanning of the image performed by trivy
COSIGN_EXPERIMENTAL=1 cosign verify-attestation --type vuln ghcr.io/braderz/monorepo/reviews:pr-42
```