package typst

import (
	"fmt"
	"strings"

	"github.com/Memnoc/lineage/internal/parser"
)

/*
* WARNING: empty struct for now
*/
type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(recipes []*parser.ProcessedRecipe) string {
	var b strings.Builder

	// WARNING: newer Fletcher API breakes the whole app for now LMAO
	// b.WriteString("#import \"@preview/fletcher:0.4.5\" as fletcher: node, edge\n\n")
	// b.WriteString("= Lineage Recipe Visualization\n\n")

	for _, recipe := range recipes {
		g.generateRecipe(&b, recipe)
	}

	return b.String()
}

func (g *Generator) generateRecipe(b *strings.Builder, recipe *parser.ProcessedRecipe) {
	b.WriteString(fmt.Sprintf("== %s\n\n", recipe.Name))

	if recipe.Description != "" {
		b.WriteString(fmt.Sprintf("_%s_\n\n", recipe.Description))
	}

	/*
	* Systems representation
	*/
	b.WriteString("*Systems:* ")
	b.WriteString(strings.Join(recipe.Systems, ", "))
	b.WriteString("\n\n")

	/*
	*Systems Connections
	*/
	b.WriteString("*Connections:*\n")
	for _, conn := range recipe.Connections {
		status := "Custom"
		if conn.IsBuiltIn {
			status = "Built-in"
		}
		b.WriteString(fmt.Sprintf("- %s: %s (%s)\n",
			conn.Provider, conn.ConnectionName, status))
	}
	b.WriteString("\n")

	g.generateDiagram(b, recipe)

	g.generateTable(b, recipe)

	b.WriteString("#pagebreak()\n\n")
}

func (g *Generator) generateDiagram(b *strings.Builder, recipe *parser.ProcessedRecipe) {
	b.WriteString("#grid(\n")
	b.WriteString("  columns: " + fmt.Sprintf("%d", len(recipe.Actions)+1) + ",\n")
	b.WriteString("  gutter: 2em,\n")
	b.WriteString("  align: center,\n")

	/*
	*Trigger box
	*/
	b.WriteString(fmt.Sprintf("  box(fill: rgb(\"#e0f2ff\"), inset: 1em, radius: 0.5em)[*%s*],\n",
		formatName(recipe.Trigger.System)))

	/*
	*Action boxes
	*/
	for _, action := range recipe.Actions {
		b.WriteString(fmt.Sprintf("  box(fill: rgb(\"#bae6fd\"), inset: 1em, radius: 0.5em)[*%s*],\n",
			formatName(action.System)))
	}

	b.WriteString(")\n\n")

	/*
	* WARNING: Flow arrows as text is much easier for now
	*/
	b.WriteString("*Data Flow:* ")
	prevSystem := recipe.Trigger.System
	for _, flow := range recipe.Flow {
		b.WriteString(fmt.Sprintf("%s → ", formatName(prevSystem)))
		prevSystem = flow.To
	}
	b.WriteString(fmt.Sprintf("%s\n\n", formatName(prevSystem)))
}

/*
* Table syntax
* WARNING: not sure how this expands with longer recipes
* */
func (g *Generator) generateTable(b *strings.Builder, recipe *parser.ProcessedRecipe) {
	b.WriteString("#table(\n")
	b.WriteString("  columns: (auto, 1fr, 1fr),\n")
	b.WriteString("  [*Step*], [*System*], [*Action*],\n")

	b.WriteString(fmt.Sprintf("  [0], [%s], [%s],\n",
		recipe.Trigger.System, formatName(recipe.Trigger.Action)))

	for _, action := range recipe.Actions {
		b.WriteString(fmt.Sprintf("  [%d], [%s], [%s],\n",
			action.StepNumber, action.System, formatName(action.Action)))
	}

	b.WriteString(")\n\n")
}

func formatName(s string) string {
	parts := strings.Split(s, "_")
	var formatted []string
	for _, part := range parts {
		if len(part) > 0 {
			formatted = append(formatted, strings.Title(part))
		}
	}
	return strings.Join(formatted, " ")
}
