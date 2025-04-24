Feature: JKS format output

  As user, I need VCert to output my certificate, private key, and chain certificates in the JKS format
  required by my application so that I don't have to use OpenSSL to combine the individual PEM files generated
  by VCert into a PKCS#12 keystore and then use the Java KeyTool to convert the PKCS#12 keystore to a Java keystore.

  - User requests JKS by specifying "jks" after the -format switch

  - User must use the -file switch to specify the name of the keystore file when they specify -format jks
    (i.e. neither the -cert-file, -key-file, nor -chain-file switches may appear on the command line,
    and console output as a base64 encoded blob will not be supported).

  - User always will be request for a password of at least 6 characters when he request a JKS file so it isn't possible to use the -no-prompt switch

  - User specifies the password for the JKS file and the key entry, conformed by the private key, the certificate and the chain certificates, using, preferable,
    the -jks-password or alternatively the -key-password switch

  - User must use the -jks-alias switch to provide the alias for the jks entry that will be conformed by the private key, the certificate and the chain certificates

  - JKS format is not allowed for the enroll or renew actions when -csr is "file"

  - JKS format is not allowed for the enroll or renew action when -csr is "local" (or not specified)
    and the -no-pickup switch is used

  - JKS format is only allowed for the pickup action when the private key is stored in the Venafi Platform

  
  Background:
    And the default aruba exit timeout is 180 seconds

  @FAKE
  Scenario: where it outputs error if JKS format is specified, but STDOUT output is used (default output)
    When I enroll random certificate in test-mode with -no-prompt -format jks
      Then it should fail with "JKS format requires certificate, private key, and chain to be written to a single file; specify using --file"
    When I retrieve the certificate in test-mode with -pickup-id xxx -key-password xxx -format jks
      Then it should fail with "JKS format requires certificate, private key, and chain to be written to a single file; specify using --file"
    When I renew the certificate in TPP with flags -id xxx -no-prompt -format jks
      Then it should fail with "JKS format requires certificate, private key, and chain to be written to a single file; specify using --file"

  @FAKE
  Scenario: where JKS format is specified, but a short password is used
    When I enroll random certificate in test-mode with -no-prompt -format jks -key-password 1234 -jks-password 123456 -file all.jks
      Then it should fail with "JKS format requires passwords that are at least 6 characters long"
    When I enroll random certificate in test-mode with -no-prompt -format jks -key-password 123456 -jks-password 1234 -file all.jks
      Then it should fail with "JKS format requires passwords that are at least 6 characters long"
    When I enroll random certificate in test-mode with -no-prompt -format jks -key-password 1234 -jks-password 1234 -file all.jks
          Then it should fail with "JKS format requires passwords that are at least 6 characters long"

  @FAKE
  Scenario: where JKS format is specified and a password is used but the jks alias is not provided
    When I enroll random certificate in test-mode with -no-prompt -format jks -key-password 123456 -file all.jks
      Then it should fail with "The --jks-alias parameter is required with --format jks"

  @FAKE
  Scenario: where JKS format is specified but a password is not provided
      When I enroll random certificate in test-mode with -no-prompt -format jks -file all.jks -jks-alias alias
        Then it should fail with "JKS format requires passwords that are at least 6 characters long"

  @FAKE
  Scenario: where JKS format is specified but a key-password is not provided
      When I enroll random certificate in test-mode with -no-prompt -format jks -file all.jks -jks-password 123456 -jks-alias alias
        Then it should fail with "JKS format requires passwords that are at least 6 characters long"

  @FAKE
  Scenario: where JKS format is not specified but the jks password is provided
    When I enroll random certificate in test-mode with -no-prompt -format pkcs12 -jks-password 123456 -file all.jks
      Then it should fail with "The --jks-password parameter may only be used with --format jks"

  @FAKE
  Scenario: where JKS format is not specified but the jks alias is provided
    When I enroll random certificate in test-mode with -no-prompt -format pkcs12 -jks-alias alias -file all.jks
      Then it should fail with "The --jks-alias parameter may only be used with --format jks"

  @FAKE
  Scenario: where all objects are written to one JKS archive
    When I enroll random certificate in test-mode with -no-prompt -format jks -file all.jks -key-password 123456 -jks-password 123456 -jks-alias abc
    Then the exit status should be 0
    And "all.jks" should be jks archive with password "123456"

  @FAKE
  Scenario: where all objects are written to one JKS archive
      When I enroll random certificate in test-mode with -no-prompt -format jks -file all.jks -key-password 123456 -jks-password 123456 -jks-alias abc -key-type ecdsa
      Then the exit status should be 0
      And "all.jks" should be jks archive with password "123456"

  Scenario Outline: where all objects are written to one JKS archive with key-password and providing the jks-password
    When I enroll random certificate in <endpoint> with -format jks -file all.jks -key-password 123abc -jks-password 123456 -jks-alias abc
    Then the exit status should be 0
    And "all.jks" should be jks archive with password "123456"

    @FAKE
    Examples:
       | endpoint  |
       | test-mode |

    @TPP
    Examples:
      | endpoint  |
      | TPP       |

    @VAAS
    Examples:
      | endpoint  |
      | Cloud     |

  Scenario Outline: where all objects are written to one JKS archive with key-password and providing the jks-password and the key-type is ecdsa
    When I enroll random certificate in <endpoint> with -format jks -file all.jks -key-password 123abc -jks-password 123456 -jks-alias abc key-type ecdsa
    Then the exit status should be 0
    And "all.jks" should be jks archive with password "123456"

    @FAKE
    Examples:
      | endpoint  |
      | test-mode |

    @TPP
    Examples:
      | endpoint  |
      | TPP       |

    @VAAS
    Examples:
      | endpoint  |
      | Cloud     |


  Scenario Outline: where it outputs error when trying to pickup local-generated certificate and output it in JKS format
    When I enroll random certificate using <endpoint> with -no-prompt -no-pickup
    And I retrieve the certificate using <endpoint> using the same Pickup ID with -timeout 180 -no-prompt -file all.jks -format jks
    And it should fail with "key password must be provided"

    @FAKE
    Examples:
      | endpoint  |
      | test-mode |

    @TPP
    Examples:
      | endpoint  |
      | TPP       |

    @VAAS
    Examples:
      | endpoint  |
      | Cloud     |

  Scenario Outline: where it outputs error when trying to enroll certificate in -csr file: mode and output it in JKS format
    Given I generate random CSR with -no-prompt -csr-file csr.pem -key-file k.pem
    When I enroll certificate using <endpoint> with -no-prompt -csr file:csr.pem -file all.jks -format jks
    And it should fail with "The --csr \"file\" option may not be used with the enroll or renew actions when --format is \"jks\""

    @FAKE
    Examples:
      | endpoint  |
      | test-mode |

    @TPP
    Examples:
      | endpoint  |
      | TPP       |

    @VAAS
    Examples:
      | endpoint  |
      | Cloud     |

  Scenario Outline: where it outputs error when trying to enroll certificate in -csr local (by default), -no-pickup and output it in JKS format
    When I enroll random certificate using <endpoint> with -no-prompt -file all.jks -format jks -no-pickup
    And it should fail with "The --csr \"local\" option may not be used with the enroll or renew actions when --format is \"jks\" and --no-pickup is specified"

    @FAKE
    Examples:
      | endpoint  |
      | test-mode |

    @TPP
    Examples:
      | endpoint  |
      | TPP       |

    @VAAS
    Examples:
      | endpoint  |
      | Cloud     |

  Scenario Outline: where it outputs error when trying to enroll certificate in -csr local (specified), -no-pickup and output it in JKS format
    When I enroll random certificate using <endpoint> with -no-prompt -file all.jks -format jks -no-pickup -csr local
    And it should fail with "The --csr \"local\" option may not be used with the enroll or renew actions when --format is \"jks\" and --no-pickup is specified"

    @FAKE
    Examples:
      | endpoint  |
      | test-mode |

    @TPP
    Examples:
      | endpoint  |
      | TPP       |

    @VAAS
    Examples:
      | endpoint  |
      | Cloud     |

  Scenario Outline: where it pickups up service-generated certificate and outputs it in JKS format
    When I enroll random certificate using <endpoint> with -no-prompt -no-pickup -csr service
    And I retrieve the certificate using <endpoint> using the same Pickup ID and using a dummy password with -timeout 180 -file all.jks -format jks -jks-alias abc
#    And "all.jks" should be JKS archive with password "dummy password" # currently, we don't have JKS steps

    @FAKE
    Examples:
      | endpoint  |
      | test-mode |

    @TPP
    Examples:
      | endpoint  |
      | TPP       |

# TODO: Now VaaS supports CSR, but we need to verify this behavior for this test
#    @VAAS
#    Examples:
#      | endpoint  |
#      | Cloud     | # -csr service is not supported by Cloud

