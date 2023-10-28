# Expense Reporting Program

This Go program demonstrates how to implement an expense reporting system for different types of expenses, including emails, SMS messages, and invalid entries. It utilizes Go's type assertion and interfaces to handle various expense types and calculate their costs.

## Program Components

1. **Expense Interface**: The `expense` interface defines a common behavior for different types of expenses. All expense types must implement the `cost()` method.

2. **Expense Types**:
   - **Email**: Represents an email with attributes such as subscription status, message body, and destination address.
   - **SMS**: Represents an SMS message with attributes including subscription status, message body, and the recipient's phone number.
   - **Invalid**: A placeholder for invalid or unsupported expense types.

3. **Expense Cost Calculation**: Each expense type has its own implementation of the `cost()` method, which calculates the cost based on specific criteria. The `cost()` method takes into account the subscription status and message length.

4. **Expense Reporting Function**: The `getReport` function takes an expense as an argument and uses type assertion to determine the underlying type (email, SMS, or invalid). It then retrieves the appropriate information (destination address or phone number) and calculates the cost using the specific expense's `cost()` method.

5. **Test Function**: The `test` function takes an expense and calls the `getReport` function to generate a report for that expense. It formats the report and displays the result.

## How to Use the Program

1. Define expenses of different types, such as emails and SMS messages, with specific attributes.
2. Call the `test` function with these expenses to generate expense reports, including destination addresses or phone numbers and their respective costs.
3. The program will output the reports for each expense, and it will handle invalid expense types gracefully.

This program showcases the use of type assertion to work with different types that implement a common interface, making it flexible and extensible for handling various expense types.


###### Inspired by codeCamp GoLang course