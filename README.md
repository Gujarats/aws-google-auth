# Aws google auth
using google SSO and then assume role what a trick !

## Install

You can choose the binary file ready to download on this [release](Gujarats/aws-google-auth/releases).
Or you can run this command : 

``` shell
$ curl https://raw.githubusercontent.com/Gujarats/aws-google-auth/master/downloader.sh | sh
```

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

Sorry this is just lazy way to do  it :D


## TODO 

 - Rewrite  `aws-google-auth` in go 

Seems pretty straight forward but takes time.
