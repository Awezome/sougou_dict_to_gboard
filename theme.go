package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

var ui = theme.LightTheme()

type BaseTheme struct{}

func (BaseTheme) TextFont() fyne.Resource     { return resourceFangZhengHeiTiJianTiTtf }
func (BaseTheme) TextBoldFont() fyne.Resource { return resourceFangZhengHeiTiJianTiTtf }

func (BaseTheme) BackgroundColor() color.Color      { return ui.BackgroundColor() }
func (BaseTheme) ButtonColor() color.Color          { return ui.ButtonColor() }
func (BaseTheme) DisabledButtonColor() color.Color  { return ui.DisabledButtonColor() }
func (BaseTheme) IconColor() color.Color            { return ui.IconColor() }
func (BaseTheme) DisabledIconColor() color.Color    { return ui.DisabledIconColor() }
func (BaseTheme) HyperlinkColor() color.Color       { return ui.HyperlinkColor() }
func (BaseTheme) TextColor() color.Color            { return ui.TextColor() }
func (BaseTheme) DisabledTextColor() color.Color    { return ui.DisabledTextColor() }
func (BaseTheme) HoverColor() color.Color           { return ui.HoverColor() }
func (BaseTheme) PlaceHolderColor() color.Color     { return ui.PlaceHolderColor() }
func (BaseTheme) PrimaryColor() color.Color         { return ui.PrimaryColor() }
func (BaseTheme) FocusColor() color.Color           { return ui.FocusColor() }
func (BaseTheme) ScrollBarColor() color.Color       { return ui.ScrollBarColor() }
func (BaseTheme) ShadowColor() color.Color          { return ui.ShadowColor() }
func (BaseTheme) TextSize() int                     { return ui.TextSize() }
func (BaseTheme) TextItalicFont() fyne.Resource     { return ui.TextItalicFont() }
func (BaseTheme) TextBoldItalicFont() fyne.Resource { return ui.TextBoldItalicFont() }
func (BaseTheme) TextMonospaceFont() fyne.Resource  { return ui.TextMonospaceFont() }
func (BaseTheme) Padding() int                      { return ui.Padding() }
func (BaseTheme) IconInlineSize() int               { return ui.IconInlineSize() }
func (BaseTheme) ScrollBarSize() int                { return ui.ScrollBarSize() }
func (BaseTheme) ScrollBarSmallSize() int           { return ui.ScrollBarSmallSize() }
