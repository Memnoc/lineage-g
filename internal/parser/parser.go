// Package parser
package parser

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Parser struct {
	recipes     map[string]*Recipe
	connections map[string]*Connection
}

func New() *Parser {
	return &Parser{
		recipes:     make(map[string]*Recipe),
		connections: make(map[string]*Connection),
	}
}

// LoadDirectory TODO: Using WalkDir would be better here
func (p *Parser) LoadDirectory(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		if strings.HasSuffix(path, ".recipe.json") {
			return p.loadRecipeFile(path)
		}
		if strings.HasSuffix(path, ".connection.json") {
			return p.loadConnectionFile(path)
		}
		return nil
	})
}

func (p *Parser) loadRecipeFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var recipe Recipe
	if err := json.Unmarshal(data, &recipe); err != nil {
		return err
	}

	p.recipes[filepath.Base(path)] = &recipe
	return nil
}

func (p *Parser) loadConnectionFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var conn Connection
	if err := json.Unmarshal(data, &conn); err != nil {
		return err
	}

	p.connections[filepath.Base(path)] = &conn
	return nil
}

func (p *Parser) Process() []*ProcessedRecipe {
	var processed []*ProcessedRecipe

	for _, recipe := range p.recipes {
		proc := p.processRecipe(recipe)
		processed = append(processed, proc)
	}

	return processed
}

func (p *Parser) processRecipe(recipe *Recipe) *ProcessedRecipe {
	proc := &ProcessedRecipe{
		Name:        recipe.Name,
		Description: recipe.Description,
		Trigger: Step{
			System:     recipe.Code.Provider,
			Action:     recipe.Code.Name,
			StepNumber: recipe.Code.Number,
		},
	}

	/*
	* Recipe Actions
	 */
	for _, block := range recipe.Code.Block {
		proc.Actions = append(proc.Actions, Step{
			System:     block.Provider,
			Action:     block.Name,
			StepNumber: block.Number,
		})
	}

	/*
	* Unique Systems
	 */
	systemSet := make(map[string]bool)
	systemSet[proc.Trigger.System] = true
	for _, action := range proc.Actions {
		systemSet[action.System] = true
	}
	for sys := range systemSet {
		proc.Systems = append(proc.Systems, sys)
	}

	/*
	* Connections with cross-reference
	 */for _, cfg := range recipe.Config {
		connInfo := ConnectionInfo{
			Provider:  cfg.Provider,
			IsBuiltIn: cfg.AccountID == nil,
		}

		if cfg.AccountID != nil {
			// key for connections look-up logic
			if conn, exists := p.connections[cfg.AccountID.ZipName]; exists {
				connInfo.ConnectionName = conn.Name
			} else {
				connInfo.ConnectionName = cfg.AccountID.Name
			}
		} else {
			connInfo.ConnectionName = "Built-in"
		}

		proc.Connections = append(proc.Connections, connInfo)
	}

	/*
	* Recipe Flow
	 */
	allSteps := append([]Step{proc.Trigger}, proc.Actions...)
	for i := 1; i < len(allSteps); i++ {
		proc.Flow = append(proc.Flow, SystemFlow{
			From:       allSteps[i-1].System,
			To:         allSteps[i].System,
			ActionType: allSteps[i].Action,
		})
	}

	return proc
}
