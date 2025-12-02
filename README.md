# Workato Recipe Visualizer

A desktop application that transforms Workato recipe exports into clear, interactive visual diagrams showing system connections and data flow patterns.

## Purpose

When working with complex Workato integrations, understanding the connections between different systems can be challenging. Recipe JSON files contain all the workflow logic but are difficult to parse manually. This tool provides an instant visual overview of:

- **System Connections**: See which applications are connected
- **Data Flow**: Understand how data moves between systems
- **Workflow Logic**: Visualize triggers, actions, and dependencies
- **Integration Patterns**: Quickly identify common connection patterns

## Key Features

- **Drag & Drop Import**: Upload Workato recipe export files (.recipe.json)
- **Visual Diagrams**: Generate interactive connection diagrams
- **System Overview**: See all connected systems at a glance
- **Multi-Recipe Analysis**: Analyze multiple recipes simultaneously
- **Connection Details**: View account names and configuration details
- **Export Options**: Save diagrams for documentation

## How It Works

1. **Export** your Workato recipes as JSON files
2. **Import** the files into the application
3. **Visualize** system connections instantly
4. **Explore** data flow and integration patterns

## Supported File Types

- `.recipe.json` - Contains workflow logic and system connections
- `.connection.json` - Contains account details and configuration

## What can you use this for

- **Integration Documentation**: Generate visual documentation for integrations
- **System Architecture Review**: Understand data flow across systems
- **Troubleshooting**: Quickly identify connection points and dependencies
- **Onboarding**: Help new team members understand existing integrations
- **Compliance**: Document data flow for security and compliance reviews

## Example Output

The tool converts complex JSON manifests like this:

```json
{
  "name": "Sync Customer Data",
  "code": {
    "provider": "salesforce",
    "block": [{ "provider": "google_sheets" }, { "provider": "slack" }]
  }
}
```

Into clear visual diagrams showing:

```
Salesforce → Google Sheets → Slack
   ↓              ↓           ↓
[Trigger]     [Update]    [Notify]
```

## Basic Workato export looks like

```bash
workato-export.zip
├── recipe1.recipe.json
├── recipe2.recipe.json
├── connection1.connection.json
└── connection2.connection.json
```

## Getting Started (TBD)

```bash
1. make build
2. ./bin/lineage testdata/fixtures
3. typst compile recipes.typ
```

```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

## 📄 License

## MIT

**Note**: This tool is designed for Workato recipe analysis and is not affiliated with Workato Inc.
```
