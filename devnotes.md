# Workato files to parse

## Where to find key information

Systems Connections & Workflow Logic -> `recipe.json`
Account Details -> `.connection.json` files

## Files content

1. Recipe files ending in `.recipe.json`

- _Contains:_ workflow logic, systems connections, data flow

2. Connection files ending in `.connection.json`

- _Contains:_ Account/connection details

## File Processing Strategy

### Primary layer -> `recipe.json`

90% of what I need is in:

- Systems types (**provider** fields)
- Workflow Steps (**code.block** array)
- Connection references (**config.account_id.zip_name** )

### Secondary layer

Enrichment layer

- Human readable connection names
- Additional connection metadata

### Key-value pairs sample

1. Recipe Metadata (TOP LEVEL)

`{
  "name": "Recipe name",
  "description": "What the recipe does",
  "version": 3
}`

2. Systems Connections (Most of the info here)

`"config": [
  {
    "provider": "google_sheets",  // System type
    "account_id": {
      "name": "My Google Sheets account"  // Connection name
    }
  },
  {
    "provider": "email",
    "account_id": null  // Built-in connector
  }
]`

3. Workflow Steps (Recipe Logic)

### Trigger

`{
  "number": 0,
  "provider": "google_sheets",        // Source system
  "name": "new_spreadsheet_row_v4",   // What triggers it
  "keyword": "trigger"
}`

### Actions (in the `block` array)

`{
  "number": 1,
  "provider": "email",           // Target system  
  "name": "send_mail",          // What action
  "keyword": "action"
}`

4. Data Flow (How data flows between systems)

`"input": {
  "data": {
    "col_1": "#{...output from previous step...}"  // Data mapping
  }
}`

## Basic parsing steps

1. We scan and load the files in a single directory walk, loading both file types into separate maps
2. The processing for each recipe is: cross-reference zip_name → look up in connections map → merge data
3. Transformation of the processed data into Typst markup

- 4a Load the recipe
  `"config": [
  {
    "provider": "google_sheets",
    "account_id": {
      "zip_name": "my_google_sheets_account.connection.json",  // ← Points to connection file
      "name": "My Google Sheets account"
    }
  }
]`

  **long_filename.connection.json**

- 4b look up the connection file

`{
  "name": "My Google Sheets account",     // ← Same name, more details here
  "provider": "google_sheets",
  "root_folder": false
}`

- 4c merge data: Combine recipe workflow info with connection details

5. Generate visualization from the combined data
