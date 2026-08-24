package common

type ResultCode uint16

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
New returns an instance of error circumscribing a concrete
[Error] instance bearing the result code in the receiver
alongside an optional text message.
*/
func (r ResultCode) New(msg ...string) error {
	return newRes(uint16(r), msg...)
}

const (
	LDAPResultCanceled        ResultCode = iota + 118 // 118
	LDAPResultNoSuchOperation                         // 119
	LDAPResultTooLate                                 // 120
	LDAPResultCannotCancel                            // 121
)

const (
	ErrorNetwork         ResultCode = iota + 200 // 200
	ErrorFilterCompile                           // 201
	ErrorFilterDecompile                         // 202
	ErrorDebugging                               // 203
)
