# Aws google auth
using google SSO and then assume role what a trick !

## Usage

Create config file in `$HOME/.aws-google-auth/config.yaml`

```yaml
---
profileParent: saml
profile: development
```

All the profiles above is store in `~/.aws/config`

## Dependencies

you must install :
 - [aws-google-auth](https://github.com/cevoaustralia/aws-google-auth)
 - [awsudo](https://github.com/makethunder/awsudo)

Sorry this is just lazy way to do  it :D


## TODO 

 - Rewrite  `aws-google-auth` in go
 - Rewrite `awsudo` in go

Seems pretty straight forward but takes time.
