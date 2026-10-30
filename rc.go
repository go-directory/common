package common

type ResultCode uint16

/*
New returns an instance of error circumscribing a concrete
[Error] instance bearing the result code in the receiver
alongside an optional text message.
*/
func (r ResultCode) New(msg ...string) Error {
	return newRes(uint16(r), msg...)
}

const (
	LDAPResultSuccess                      ResultCode = iota // 0
	LDAPResultOperationsError                                // 1
	LDAPResultProtocolError                                  // 2
	LDAPResultTimeLimitExceeded                              // 3
	LDAPResultSizeLimitExceeded                              // 4
	LDAPResultCompareFalse                                   // 5
	LDAPResultCompareTrue                                    // 6
	LDAPResultAuthMethodNotSupported                         // 7
	LDAPResultStrongAuthRequired                             // 8
	_                                                        // 9 (reserved)
	LDAPResultReferral                                       // 10
	LDAPResultAdminLimitExceeded                             // 11
	LDAPResultUnavailableCriticalExtension                   // 12
	LDAPResultConfidentialityRequired                        // 13
	LDAPResultSaslBindInProgress                             // 14
	_                                                        // 15 (reserved)
	LDAPResultNoSuchAttribute                                // 16
	LDAPResultUndefinedAttributeType                         // 17
	LDAPResultInappropriateMatching                          // 18
	LDAPResultConstraintViolation                            // 19
	LDAPResultAttributeOrValueExists                         // 20
	LDAPResultInvalidAttributeSyntax                         // 21
	_                                                        // 22 (reserved)
	_                                                        // 23 (reserved)
	_                                                        // 24 (reserved)
	_                                                        // 25 (reserved)
	_                                                        // 26 (reserved)
	_                                                        // 27 (reserved)
	_                                                        // 28 (reserved)
	_                                                        // 29 (reserved)
	_                                                        // 30 (reserved)
	_                                                        // 31 (reserved)
	LDAPResultNoSuchObject                                   // 32
	LDAPResultAliasProblem                                   // 33
	LDAPResultInvalidDNSyntax                                // 34
	_                                                        // 35 (reserved)
	LDAPResultAliasDereferencingProblem                      // 36
	_                                                        // 37 (reserved)
	_                                                        // 38 (reserved)
	_                                                        // 39 (reserved)
	_                                                        // 40 (reserved)
	_                                                        // 41 (reserved)
	_                                                        // 42 (reserved)
	_                                                        // 43 (reserved)
	_                                                        // 44 (reserved)
	_                                                        // 45 (reserved)
	_                                                        // 46 (reserved)
	_                                                        // 47 (reserved)
	LDAPResultInappropriateAuthentication                    // 48
	LDAPResultInvalidCredentials                             // 49
	LDAPResultInsufficientAccessRights                       // 50
	LDAPResultBusy                                           // 51
	LDAPResultUnavailable                                    // 52
	LDAPResultUnwillingToPerform                             // 53
	LDAPResultLoopDetect                                     // 54
	_                                                        // 55 (reserved)
	_                                                        // 56 (reserved)
	_                                                        // 57 (reserved)
	_                                                        // 58 (reserved)
	_                                                        // 59 (reserved)
	_                                                        // 60 (reserved)
	_                                                        // 61 (reserved)
	_                                                        // 62 (reserved)
	_                                                        // 63 (reserved)
	LDAPResultNamingViolation                                // 64
	LDAPResultObjectClassViolation                           // 65
	LDAPResultNotAllowedOnNonLeaf                            // 66
	LDAPResultNotAllowedOnRDN                                // 67
	LDAPResultEntryAlreadyExists                             // 68
	LDAPResultObjectClassModsProhibited                      // 69
	_                                                        // 70 (reserved)
	LDAPResultAffectsMultipleDSAs                            // 71
	_                                                        // 72 (reserved)
	_                                                        // 73 (reserved)
	_                                                        // 74 (reserved)
	_                                                        // 75 (reserved)
	_                                                        // 76 (reserved)
	_                                                        // 77 (reserved)
	_                                                        // 78 (reserved)
	_                                                        // 79 (reserved)
	LDAPResultOther                                          // 80
)

/*
Extended LDAP result codes.
*/
const (
	LDAPResultLCUPResOut ResultCode = iota +113       // 113 -- RFC 3928
	LDAPResultLCUPSecViolation			  // 114 -- RFC 3928
	LDAPResultLCUPInvalidData			  // 115 -- RFC 3928
	LDAPResultLCUPUnsupportedSchema			  // 116 -- RFC 3928
	LDAPResultLCUPReloadRequired			  // 117 -- RFC 3928
        LDAPResultCanceled                                // 118 -- RFC 3909
        LDAPResultNoSuchOperation                         // 119 -- RFC 3909
        LDAPResultTooLate                                 // 120 -- RFC 3909
        LDAPResultCannotCancel                            // 121 -- RFC 3909
        LDAPResultAssertionFailed                         // 122 -- RFC 4528
        LDAPAuthorizationDenied                           // 123 -- RFC 4370
)

/*
Topical errors.
*/
const (
	ErrorNetwork ResultCode = iota + 200 // 200 ; general network error
	ErrorClient                          // 201 ; general dUA (client) error
	ErrorSchema                          // 202 ; general schema error
	ErrorBackend                         // 203 ; general backend (dib) error
	ErrorSyntax                          // 204 ; general syntax error
	ErrorServer                          // 205 ; general dSA (server) error
	ErrorParse                           // 206 ; general parsing error
	ErrorConfig                          // 207 ; general configuration error
	ErrorGeneral                         // 208 ; general unclassified error
)

/*
Codec errors.
*/
const (
	ErrorASN1            ResultCode = iota + 300 // 300 ; general ASN.1 codec error
	ErrorFilterCompile                           // 301 ; filter compilation (encoding) error
	ErrorFilterDecompile                         // 302 ; filter decompilation (decoding) error
)

const LDAPResultESyncRefreshRequired ResultCode = 4096 // 4096 -- RFC 4533
