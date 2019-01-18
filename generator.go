package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type ConsentItem struct {
	ConsentId                 string
	Iban                      string
	Status                    string
	RecurringIndicator        bool
	FrequencyPerDay           string
	LastUpdated               string
	RequestId                 string
	TppAuthorizationNumber    string
	TransactionStatusDetailed string
	ValidUntil                string
}

func main() {

	consentItemTemplate, err := template.ParseFiles("templates/consentItem.tpl")
	if err != nil {
		panic(err)
	}

	// Execute consent item template and store each output in the array
	var consentItems []string
	for _, element := range getConsents() {
		var tpl bytes.Buffer
		if err := consentItemTemplate.Execute(&tpl, element); err != nil {
			panic(err)
		}
		consentItems = append(consentItems, tpl.String())
	}

	// Execute dynamodb template and store the output in the buffer
	var output bytes.Buffer
	dynamoDbTemplate, err := template.New("dynamodbInput.tpl").Funcs(template.FuncMap{"addSeparator": addSeparator}).ParseFiles("templates/dynamodbInput.tpl")
	if err != nil {
		panic(err)
	}
	if err := dynamoDbTemplate.Execute(&output, consentItems); err != nil {
		panic(err)
	}

	saveToFile(output)
}

/** This function returns a function which determines if the comma separator must be added. Will be used in the template to avoid the last comma in the list of PutRequest items */
func addSeparator() func() bool {
	firstInvocation := true

	return func() bool {
		if firstInvocation {
			firstInvocation = false

			return false
		}

		return true
	}
}

func getUuid() string {
	uuid, _ := exec.Command("uuidgen").Output()
	return strings.ToLower(strings.TrimSuffix(string(uuid), "\n"))
}

func getConsents() []ConsentItem {
	var validConsent = ConsentItem{
		"00000000-0000-0000-0000-000000000001", "NL66MOYO0000000001", "valid", true, "24", "2019-01-15T15:06:59.920Z", getUuid(), "12345678901", "SUCCESS", "2019-04-15",
	}
	var rejectedCancelledConsent = ConsentItem{
		"00000000-0000-0000-0000-000000000002", "NL66MOYO0000000002", "rejected", false, "4", "2019-01-15T15:00:00.001Z", getUuid(), "12345678902", "CANCELLED", "9999-12-31",
	}
	var rejectedExpiredConsent = ConsentItem{
		"00000000-0000-0000-0000-000000000003", "NL66MOYO0000000003", "rejected", false, "4", "2019-01-15T15:00:00.001Z", getUuid(), "12345678903", "EXPIRED", "9999-12-31",
	}
	var rejectedFailureConsent = ConsentItem{
		"00000000-0000-0000-0000-000000000004", "NL66MOYO0000000004", "rejected", false, "4", "2019-01-15T15:00:00.001Z", getUuid(), "12345678904", "FAILURE", "9999-12-31",
	}
	var expiredConsent = ConsentItem{
		"00000000-0000-0000-0000-000000000005", "NL66MOYO0000000005", "expired", true, "24", "2019-01-15T15:00:00.001Z", getUuid(), "12345678905", "SUCCESS", "2019-01-15",
	}
	var revokedByPsuConsent = ConsentItem{
		"00000000-0000-0000-0000-000000000006", "NL66MOYO0000000006", "revokedByPsu", false, "24", "2019-01-15T15:00:00.001Z", getUuid(), "12345678906", "SUCCESS", "2019-01-30",
	}
	var terminatedByTppConsent = ConsentItem{
		"00000000-0000-0000-0000-000000000007", "NL66MOYO0000000007", "terminatedByTpp", false, "24", "2019-01-15T15:00:00.001Z", getUuid(), "12345678907", "SUCCESS", "2019-01-30",
	}

	return []ConsentItem{
		validConsent,
		rejectedCancelledConsent,
		rejectedExpiredConsent,
		rejectedFailureConsent,
		expiredConsent,
		revokedByPsuConsent,
		terminatedByTppConsent,
	}
}

func saveToFile(output bytes.Buffer) {
	// Format the output in pretty json
	var prettyOutput bytes.Buffer
	json.Indent(&prettyOutput, output.Bytes(), "", "\t")

	// Save the output to file
	outputFile, err := os.Create("output.json")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(outputFile, bytes.NewReader(prettyOutput.Bytes()))
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()
}
