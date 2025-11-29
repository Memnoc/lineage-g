== Send an email when non business domain is added to the Customer Contact List

*Systems:* email, google_sheets

*Connections:*
- google_sheets: My Google Sheets account (Custom)
- email: Built-in (Built-in)

#grid(
  columns: 2,
  gutter: 2em,
  align: center,
  box(fill: rgb("#e0f2ff"), inset: 1em, radius: 0.5em)[*Google Sheets*],
  box(fill: rgb("#bae6fd"), inset: 1em, radius: 0.5em)[*Email*],
)

*Data Flow:* Google Sheets → Email

#table(
  columns: (auto, 1fr, 1fr),
  [*Step*], [*System*], [*Action*],
  [0], [google_sheets], [Updated Spreadsheet Row V4 2],
  [1], [email], [Send Mail],
)

#pagebreak()

