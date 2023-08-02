Feature: playbook

  As a user
  I want to issue certificates using playbook and perform installation

  Background:
    And the default aruba exit timeout is 180 seconds

  @TPP
  Scenario Outline: Run playbook for TPP with extended configuration with PEM, PKCS12 and PEM installations
    Given I have playbook with TPP connection details
    And I have playbook with certificateTasks block
    And I have playbook with task named "myCertificateInstallation"
    And task named "myCertificateInstallation" has setenvvars "thumbprint,serial"
    And task named "myCertificateInstallation" has renewBefore with value "31d"
    And task named "myCertificateInstallation" has request
    And task named "myCertificateInstallation" has request with "chain" value "root-first"
    And task named "myCertificateInstallation" has request with "csr" value "service"
    And task named "myCertificateInstallation" has request with "fields" value "custom="Foo",cfList="item1",cfListMulti="tier2|tier3|tier4""
    And task named "myCertificateInstallation" has request with "sanDns" value "test.com,test2.com"
    And task named "myCertificateInstallation" has request with "sanEmail" value "test@test.com,test2@test.com"
    And task named "myCertificateInstallation" has request with "fetchPrivateKey" value "true"
    And task named "myCertificateInstallation" has request with "sanIP" value "127.0.0.1,192.168.1.2"
    # m = Microsoft
    And task named "myCertificateInstallation" has request with "issuerHint" value "m"
    And task named "myCertificateInstallation" has request with "validDays" value "30"
    And task named "myCertificateInstallation" has request with "keyType" value "rsa"
    And task named "myCertificateInstallation" has request with "keySize" value "4096"
    # "origin" is the full name for adding to meta information to certificate request
    And task named "myCertificateInstallation" has request with "appInfo" value "Venafi VCert CLI"
    And task named "myCertificateInstallation" has request with "sanUpn" value "test,test2"
    And task named "myCertificateInstallation" has request with "sanUri" value "uri.test.com,foo.venafi.com"
    And task named "myCertificateInstallation" has request with "keyPassword" value "Passcode123!"
    And task named "myCertificateInstallation" has request with default TPP zone
    And task named "myCertificateInstallation" has request with Location instance "devops-instance", workload prefixed by "workload", tlsAddress "wwww.example.com:443" and replace "false"
    And task named "myCertificateInstallation" request has subject
    And task named "myCertificateInstallation" request has subject with "country" value "US"
    And task named "myCertificateInstallation" request has subject with "locality" value "Salt Lake City"
    And task named "myCertificateInstallation" request has subject with "province" value "Utah"
    And task named "myCertificateInstallation" request has subject with "organization" value "Venafi Inc"
    And task named "myCertificateInstallation" request has subject with "orgUnits" value "engineering,marketing"
    And task named "myCertificateInstallation" request has subject random CommonName
    And task named "myCertificateInstallation" has request with nickname based on commonName
    And task named "myCertificateInstallation" has installations
    And task named "myCertificateInstallation" has installation format PEM with file name "cert.cer", chain name "chain.cer" and key name "key.pem" with installation and validation and uses backup
    And task named "myCertificateInstallation" has installation format JKS with cert name "cert.jks", jksAlias "venafi" and jksPassword "foobar123" with installation
    And task named "myCertificateInstallation" has installation format PKCS12 with cert name "cert.p12" with validation
    And I created playbook named "<config-file>" with previous content
    And I run `vcert run -f <config-file> --force-renew`
    Then the output should contain "Successfully executed installation validation actions"
    And the output should contain "playbook run finished"
    And a file named "cert.cer" should exist
    And a file named "chain.cer" should exist
    And a file named "key.pem" should exist
    And a file named "cert.jks" should exist
    And a file named "cert.jks" should exist
    And a file named "cert.p12" should exist
    And playbook generated private key in "cert.cer" and certificate in "key.pem" should have the same modulus
    And playbook generated "cert.p12" should be PKCS#12 archive with password "Passcode123!"
#    And private key in "cert.cer" and certificate in "key.pem" should have the same modulus
    And "cert.p12" should be PKCS#12 archive with password "Passcode123!"
    # And "cert.jks" should be jks archive with password "foobar123" # TODO: solve this case
    And I uninstall file named "cert.cer"
    And I uninstall file named "chain.cer"
    And I uninstall file named "key.pem"
    And I uninstall file named "cert.jks"
    And I uninstall file named "cert.p12"

    Examples:
    | config-file      |
    | playbook-tpp.yml |

