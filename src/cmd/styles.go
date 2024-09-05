package main

import "github.com/charmbracelet/lipgloss"

var titleStyle = lipgloss.NewStyle().
	Width(100).
	Bold(true).
	Foreground(lipgloss.Color("208")).
	Padding(1, 0).
	Align(lipgloss.Center)

var filesStyle = lipgloss.NewStyle().
	Width(100).
	Foreground(lipgloss.Color("7")).
	Padding(1, 1).
	Align(lipgloss.Left)

var hoveringStyle = lipgloss.NewStyle().
	Width(100).
	Bold(true).
	Foreground(lipgloss.Color("7")).
	Background(lipgloss.Color("166")).
	Padding(1, 2).
	Align(lipgloss.Left)

var selectdStyle = lipgloss.NewStyle().
	Width(100).
	Bold(true).
	Foreground(lipgloss.Color("172")).
	Padding(1, 1).
	Align(lipgloss.Left)

var activeBlockStyle = lipgloss.NewStyle().
	Width(100).
	Foreground(lipgloss.Color("7")).
	Background(lipgloss.Color("233")).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("7")).
	Padding(1, 1).
	Align(lipgloss.Left)
