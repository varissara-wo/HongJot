# Hongjot

HongJot is a mobile app that helps spenders to track their expenses. Spenders can take a picture of their receipts and the app will extract the expense information from the image. The app will then store the expense information in a database and provide a summary of the expenses to the spender.

# HongJot Diagram

```mermaid
sequenceDiagram
	Mobile App ->> HongJot API: Upload Image to S3 Bucket
	activate Mobile App
	HongJot API ->> S3 Bucket: Store Image
	HongJot API ->> Mobile App: success
	deactivate Mobile App
	S3 Bucket ->> Lambda Function: (Event Trigger)
	Lambda Function ->> Amazon Textract: Extract Text from Image
	Amazon Textract ->> Lambda Function: Slip Expense info
	Lambda Function ->> HongJot API: Send Slip Expense info
	HongJot API ->> PostgreSQL Database: Store Expense info
	Mobile App ->> HongJot API: Get Expense Summary (Spender Requests Summary) Retrieve Expense Summary
	activate Mobile App
	HongJot API ->> PostgreSQL Database: Get Expense Summary Data
	PostgreSQL Database ->> HongJot API: Return Summary Data
	HongJot API ->> Mobile App: Show Expense Summary Data
	deactivate Mobile App
```
