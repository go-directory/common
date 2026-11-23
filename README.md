# common

Package common provides common types and features used in myriad [go-directory](https://github.com/go-directory) packages. This package is designed for exclusive use in that suite only.

Error constants were ported from [vjeantet/ldapserver](https://github.com/vjeantet/ldapserver) and reformatted. See the `LICENSE.ldapserver-constants` file in the repository root.

# About

This package extends the following:

  - Error constructor, helpers
  - LDAP result codes (conveniently bound to error construction)
  - General-use Key/Value storage type
  - An atomic counter
  - GeneralizedTime/Go Time helpers
  - General-use logging facility
