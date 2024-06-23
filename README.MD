# MOA (My own Accountant)

MOA is your own accountant. Deploy and use it as is to keep records of your money flows.

# Written on
Frontend - React
Backend - Go
BD - MongoDB


# Database structure

```json
{   
    "Cash" : "CashCount_int"
    "Categories_array": [
        "CategoryName1_string" : "MoneySpend_int",
        "CategoryName2_string" : "MoneSpend_int" 
    ]
}
```