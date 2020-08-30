package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type myTheme struct{}

// return bundled font resource
func (myTheme) TextFont() fyne.Resource     { return resourceFangZhengHeiTiJianTiTtf }
func (myTheme) TextBoldFont() fyne.Resource { return resourceFangZhengHeiTiJianTiTtf }

func (myTheme) BackgroundColor() color.Color      { return theme.LightTheme().BackgroundColor() }
func (myTheme) ButtonColor() color.Color          { return theme.LightTheme().ButtonColor() }
func (myTheme) DisabledButtonColor() color.Color  { return theme.LightTheme().DisabledButtonColor() }
func (myTheme) IconColor() color.Color            { return theme.LightTheme().IconColor() }
func (myTheme) DisabledIconColor() color.Color    { return theme.LightTheme().DisabledIconColor() }
func (myTheme) HyperlinkColor() color.Color       { return theme.LightTheme().HyperlinkColor() }
func (myTheme) TextColor() color.Color            { return theme.LightTheme().TextColor() }
func (myTheme) DisabledTextColor() color.Color    { return theme.LightTheme().DisabledTextColor() }
func (myTheme) HoverColor() color.Color           { return theme.LightTheme().HoverColor() }
func (myTheme) PlaceHolderColor() color.Color     { return theme.LightTheme().PlaceHolderColor() }
func (myTheme) PrimaryColor() color.Color         { return theme.LightTheme().PrimaryColor() }
func (myTheme) FocusColor() color.Color           { return theme.LightTheme().FocusColor() }
func (myTheme) ScrollBarColor() color.Color       { return theme.LightTheme().ScrollBarColor() }
func (myTheme) ShadowColor() color.Color          { return theme.LightTheme().ShadowColor() }
func (myTheme) TextSize() int                     { return theme.LightTheme().TextSize() }
func (myTheme) TextItalicFont() fyne.Resource     { return theme.LightTheme().TextItalicFont() }
func (myTheme) TextBoldItalicFont() fyne.Resource { return theme.LightTheme().TextBoldItalicFont() }
func (myTheme) TextMonospaceFont() fyne.Resource  { return theme.LightTheme().TextMonospaceFont() }
func (myTheme) Padding() int                      { return theme.LightTheme().Padding() }
func (myTheme) IconInlineSize() int               { return theme.LightTheme().IconInlineSize() }
func (myTheme) ScrollBarSize() int                { return theme.LightTheme().ScrollBarSize() }
func (myTheme) ScrollBarSmallSize() int           { return theme.LightTheme().ScrollBarSmallSize() }
