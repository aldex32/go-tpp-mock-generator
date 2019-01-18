# Go TPP mock generator

This is a Go project which I used to generate AWS DynamoDB scripts with predefined data.

*Remarks:* Currently it generates only AIS Consent mock items.

## Generate AIS Consents for TPP tests

This command will generate the output.json file to eventually be uploaded to a DynamoBD table:
```console
go run generator.go
```

## Eventually upload them in the DynamoDB mock table

```console
aws --region ${AWS_REGION} dynamodb batch-write-item --request-items "file://output.json"
```
