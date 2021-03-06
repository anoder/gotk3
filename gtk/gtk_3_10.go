// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// This file includes wrapers for symbols included since GTK 3.10, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.8, use
// 'go build -tags gtk_3_8'.  Otherwise, if no build tags are used, GTK 3.10
// is assumed and this file is built.
// +build !gtk_3_6,!gtk_3_8

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_3_10.go.h"
import "C"
import (
	"github.com/conformal/gotk3/gdk"
	"github.com/conformal/gotk3/glib"
	"runtime"
	"unsafe"
)

/*
 * Constants
 */

const (
	ALIGN_BASELINE Align = C.GTK_ALIGN_BASELINE
)

// RevealerTransitionType is a representation of GTK's GtkRevealerTransitionType.
type RevealerTransitionType int

const (
	REVEALER_TRANSITION_TYPE_NONE        RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_NONE
	REVEALER_TRANSITION_TYPE_CROSSFADE   RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_CROSSFADE
	REVEALER_TRANSITION_TYPE_SLIDE_RIGHT RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT
	REVEALER_TRANSITION_TYPE_SLIDE_LEFT  RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT
	REVEALER_TRANSITION_TYPE_SLIDE_UP    RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP
	REVEALER_TRANSITION_TYPE_SLIDE_DOWN  RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN
)

// StackTransitionType is a representation of GTK's GtkStackTransitionType.
type StackTransitionType int

const (
	STACK_TRANSITION_TYPE_NONE             StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_NONE
	STACK_TRANSITION_TYPE_CROSSFADE        StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_CROSSFADE
	STACK_TRANSITION_TYPE_SLIDE_RIGHT      StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT
	STACK_TRANSITION_TYPE_SLIDE_LEFT       StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT
	STACK_TRANSITION_TYPE_SLIDE_UP         StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP
	STACK_TRANSITION_TYPE_SLIDE_DOWN       StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN
	STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT
	STACK_TRANSITION_TYPE_SLIDE_UP_DOWN    StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN
)

/*
 * GtkHeaderBar
 */

type HeaderBar struct {
	Container
}

// Native returns a pointer to the underlying GtkHeaderBar.
func (v *HeaderBar) Native() *C.GtkHeaderBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkHeaderBar(p)
}

func wrapHeaderBar(obj *glib.Object) *HeaderBar {
	return &HeaderBar{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// HeaderBarNew is a wrapper around gtk_header_bar_new().
func HeaderBarNew() (*HeaderBar, error) {
	c := C.gtk_header_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	h := wrapHeaderBar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return h, nil
}

// SetTitle is a wrapper around gtk_header_bar_set_title().
func (v *HeaderBar) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_title(v.Native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_header_bar_get_title().
func (v *HeaderBar) GetTitle() string {
	cstr := C.gtk_header_bar_get_title(v.Native())
	return C.GoString((*C.char)(cstr))
}

// SetSubtitle is a wrapper around gtk_header_bar_set_subtitle().
func (v *HeaderBar) SetSubtitle(subtitle string) {
	cstr := C.CString(subtitle)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_subtitle(v.Native(), (*C.gchar)(cstr))
}

// GetSubtitle is a wrapper around gtk_header_bar_get_subtitle().
func (v *HeaderBar) GetSubtitle() string {
	cstr := C.gtk_header_bar_get_subtitle(v.Native())
	return C.GoString((*C.char)(cstr))
}

// SetCustomTitle is a wrapper around gtk_header_bar_set_custom_title().
func (v *HeaderBar) SetCustomTitle(titleWidget IWidget) {
	C.gtk_header_bar_set_custom_title(v.Native(), titleWidget.toWidget())
}

// GetCustomTitle is a wrapper around gtk_header_bar_get_custom_title().
func (v *HeaderBar) GetCustomTitle() (*Widget, error) {
	c := C.gtk_header_bar_get_custom_title(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// PackStart is a wrapper around gtk_header_bar_pack_start().
func (v *HeaderBar) PackStart(child IWidget) {
	C.gtk_header_bar_pack_start(v.Native(), child.toWidget())
}

// PackEnd is a wrapper around gtk_header_bar_pack_end().
func (v *HeaderBar) PackEnd(child IWidget) {
	C.gtk_header_bar_pack_end(v.Native(), child.toWidget())
}

// SetShowCloseButton is a wrapper around gtk_header_bar_set_show_close_button().
func (v *HeaderBar) SetShowCloseButton(setting bool) {
	C.gtk_header_bar_set_show_close_button(v.Native(), gbool(setting))
}

// GetShowCloseButton is a wrapper around gtk_header_bar_get_show_close_button().
func (v *HeaderBar) GetShowCloseButton() bool {
	c := C.gtk_header_bar_get_show_close_button(v.Native())
	return gobool(c)
}

/*
 * GtkLabel
 */

// GetLines() is a wrapper around gtk_label_get_lines().
func (v *Label) GetLines() int {
	c := C.gtk_label_get_lines(v.Native())
	return int(c)
}

// SetLines() is a wrapper around gtk_label_set_lines().
func (v *Label) SetLines(lines int) {
	C.gtk_label_set_lines(v.Native(), C.gint(lines))
}

/*
 * GtkListBox
 */

// ListBox is a representation of GTK's GtkListBox.
type ListBox struct {
	Container
}

// Native returns a pointer to the underlying GtkListBox.
func (v *ListBox) Native() *C.GtkListBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkListBox(p)
}

func wrapListBox(obj *glib.Object) *ListBox {
	return &ListBox{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// ListBoxNew is a wrapper around gtk_list_box_new().
func ListBoxNew() (*ListBox, error) {
	c := C.gtk_list_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapListBox(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l, nil
}

// Prepend is a wrapper around gtk_list_box_prepend().
func (v *ListBox) Prepend(child IWidget) {
	C.gtk_list_box_prepend(v.Native(), child.toWidget())
}

// Insert is a wrapper around gtk_list_box_insert().
func (v *ListBox) Insert(child IWidget, position int) {
	C.gtk_list_box_insert(v.Native(), child.toWidget(), C.gint(position))
}

// SelectRow is a wrapper around gtk_list_box_select_row().
func (v *ListBox) SelectRow(row *ListBoxRow) {
	C.gtk_list_box_select_row(v.Native(), row.Native())
}

// GetSelectedRow is a wrapper around gtk_list_box_get_selected_row().
func (v *ListBox) GetSelectedRow() *ListBoxRow {
	c := C.gtk_list_box_get_selected_row(v.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapListBoxRow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l
}

// SetSelectionMode is a wrapper around gtk_list_box_set_selection_mode().
func (v *ListBox) SetSelectionMode(mode SelectionMode) {
	C.gtk_list_box_set_selection_mode(v.Native(), C.GtkSelectionMode(mode))
}

// GetSelectionMode is a wrapper around gtk_list_box_get_selection_mode()
func (v *ListBox) GetSelectionMode() SelectionMode {
	c := C.gtk_list_box_get_selection_mode(v.Native())
	return SelectionMode(c)
}

// SetActivateOnSingleClick is a wrapper around gtk_list_box_set_activate_on_single_click().
func (v *ListBox) SetActivateOnSingleClick(single bool) {
	C.gtk_list_box_set_activate_on_single_click(v.Native(), gbool(single))
}

// GetActivateOnSingleClick is a wrapper around gtk_list_box_get_activate_on_single_click().
func (v *ListBox) GetActivateOnSingleClick() bool {
	c := C.gtk_list_box_get_activate_on_single_click(v.Native())
	return gobool(c)
}

// GetAdjustment is a wrapper around gtk_list_box_get_adjustment().
func (v *ListBox) GetAdjustment() *Adjustment {
	c := C.gtk_list_box_get_adjustment(v.Native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Adjustment{glib.InitiallyUnowned{obj}}
}

// SetAdjustment is a wrapper around gtk_list_box_set_adjustment().
func (v *ListBox) SetAdjuctment(adjustment *Adjustment) {
	C.gtk_list_box_set_adjustment(v.Native(), adjustment.Native())
}

// SetPlaceholder is a wrapper around gtk_list_box_set_placeholder().
func (v *ListBox) SetPlaceholder(placeholder IWidget) {
	C.gtk_list_box_set_placeholder(v.Native(), placeholder.toWidget())
}

// GetRowAtIndex is a wrapper around gtk_list_box_get_row_at_index().
func (v *ListBox) GetRowAtIndex(index int) *ListBoxRow {
	c := C.gtk_list_box_get_row_at_index(v.Native(), C.gint(index))
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapListBoxRow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l
}

// GetRowAtY is a wrapper around gtk_list_box_get_row_at_y().
func (v *ListBox) GetRowAtY(y int) *ListBoxRow {
	c := C.gtk_list_box_get_row_at_y(v.Native(), C.gint(y))
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapListBoxRow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l
}

// InvalidateFilter is a wrapper around gtk_list_box_invalidate_filter().
func (v *ListBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter(v.Native())
}

// InvalidateHeaders is a wrapper around gtk_list_box_invalidate_headers().
func (v *ListBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers(v.Native())
}

// InvalidateSort is a wrapper around gtk_list_box_invalidate_sort().
func (v *ListBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort(v.Native())
}

// TODO: SetFilterFunc
// TODO: SetHeaderFunc
// TODO: SetSortFunc

// DragHighlightRow is a wrapper around gtk_list_box_drag_highlight_row()
func (v *ListBox) DragHighlightRow(row *ListBoxRow) {
	C.gtk_list_box_drag_highlight_row(v.Native(), row.Native())
}

/*
 * GtkListBoxRow
 */

// ListBoxRow is a representation of GTK's GtkListBoxRow.
type ListBoxRow struct {
	Bin
}

// Native returns a pointer to the underlying GtkListBoxRow.
func (v *ListBoxRow) Native() *C.GtkListBoxRow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkListBoxRow(p)
}

func wrapListBoxRow(obj *glib.Object) *ListBoxRow {
	return &ListBoxRow{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

func ListBoxRowNew() (*ListBoxRow, error) {
	c := C.gtk_list_box_row_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapListBoxRow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// Changed is a wrapper around gtk_list_box_row_changed().
func (v *ListBoxRow) Changed() {
	C.gtk_list_box_row_changed(v.Native())
}

// GetHeader is a wrapper around gtk_list_box_row_get_header().
func (v *ListBoxRow) GetHeader() *Widget {
	c := C.gtk_list_box_row_get_header(v.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

// SetHeader is a wrapper around gtk_list_box_row_get_header().
func (v *ListBoxRow) SetHeader(header IWidget) {
	C.gtk_list_box_row_set_header(v.Native(), header.toWidget())
}

// GetIndex is a wrapper around gtk_list_box_row_get_index()
func (v *ListBoxRow) GetIndex() int {
	c := C.gtk_list_box_row_get_index(v.Native())
	return int(c)
}

/*
 * GtkRevealer
 */

// Revealer is a representation of GTK's GtkRevealer
type Revealer struct {
	Bin
}

// Native returns a pointer to the underlying GtkRevealer.
func (v *Revealer) Native() *C.GtkRevealer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRevealer(p)
}

func wrapRevealer(obj *glib.Object) *Revealer {
	return &Revealer{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// RevealerNew is a wrapper around gtk_revealer_new()
func RevealerNew() (*Revealer, error) {
	c := C.gtk_revealer_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRevealer(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
}

// GetRevealChild is a wrapper around gtk_revealer_get_reveal_child().
func (v *Revealer) GetRevealChild() bool {
	c := C.gtk_revealer_get_reveal_child(v.Native())
	return gobool(c)
}

// SetRevealChild is a wrapper around gtk_revealer_set_reveal_child().
func (v *Revealer) SetRevealChild(revealChild bool) {
	C.gtk_revealer_set_reveal_child(v.Native(), gbool(revealChild))
}

// GetChildRevealed is a wrapper around gtk_revealer_get_child_revealed().
func (v *Revealer) GetChildRevealed() bool {
	c := C.gtk_revealer_get_child_revealed(v.Native())
	return gobool(c)
}

// GetTransitionDuration is a wrapper around gtk_revealer_get_transition_duration()
func (v *Revealer) GetTransitionDuration() uint {
	c := C.gtk_revealer_get_transition_duration(v.Native())
	return uint(c)
}

// SetTransitionDuration is a wrapper around gtk_revealer_set_transition_duration().
func (v *Revealer) SetTransitionDuration(duration uint) {
	C.gtk_revealer_set_transition_duration(v.Native(), C.guint(duration))
}

// GetTransitionType is a wrapper around gtk_revealer_get_transition_type()
func (v *Revealer) GetTransitionType() RevealerTransitionType {
	c := C.gtk_revealer_get_transition_type(v.Native())
	return RevealerTransitionType(c)
}

// SetTransitionType is a wrapper around gtk_revealer_set_transition_type()
func (v *Revealer) SetTransitionType(transition RevealerTransitionType) {
	t := C.GtkRevealerTransitionType(transition)
	C.gtk_revealer_set_transition_type(v.Native(), t)
}

/*
 * GtkSearchBar
 */

// SearchBar is a representation of GTK's GtkSearchBar.
type SearchBar struct {
	Bin
}

// Native returns a pointer to the underlying GtkSearchBar.
func (v *SearchBar) Native() *C.GtkSearchBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSearchBar(p)
}

func wrapSearchBar(obj *glib.Object) *SearchBar {
	return &SearchBar{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// SearchBarNew is a wrapper around gtk_search_bar_new()
func SearchBarNew() (*SearchBar, error) {
	c := C.gtk_search_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSearchBar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// ConnectEntry is a wrapper around gtk_search_bar_connect_entry().
func (v *SearchBar) ConnectEntry(entry IEntry) {
	C.gtk_search_bar_connect_entry(v.Native(), entry.toEntry())
}

// GetSearchMode is a wrapper around gtk_search_bar_get_search_mode().
func (v *SearchBar) GetSearchMode() bool {
	c := C.gtk_search_bar_get_search_mode(v.Native())
	return gobool(c)
}

// SetSearchMode is a wrapper around gtk_search_bar_set_search_mode().
func (v *SearchBar) SetSearchMode(searchMode bool) {
	C.gtk_search_bar_set_search_mode(v.Native(), gbool(searchMode))
}

// GetShowCloseButton is a wrapper arounb gtk_search_bar_get_show_close_button().
func (v *SearchBar) GetShowCloseButton() bool {
	c := C.gtk_search_bar_get_show_close_button(v.Native())
	return gobool(c)
}

// SetShowCloseButton is a wrapper around gtk_search_bar_set_show_close_button()
func (v *SearchBar) SetShowCloseButton(visible bool) {
	C.gtk_search_bar_set_show_close_button(v.Native(), gbool(visible))
}

// HandleEvent is a wrapper around gtk_search_bar_handle_event()
func (v *SearchBar) HandleEvent(event *gdk.Event) {
	e := (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	C.gtk_search_bar_handle_event(v.Native(), e)
}

/*
 * GtkStack
 */

// Stack is a representation of GTK's GtkStack.
type Stack struct {
	Container
}

// Native returns a pointer to the underlying GtkStack.
func (v *Stack) Native() *C.GtkStack {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStack(p)
}

func wrapStack(obj *glib.Object) *Stack {
	return &Stack{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// StackNew is a wrapper around gtk_stack_new().
func StackNew() (*Stack, error) {
	c := C.gtk_stack_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapStack(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// AddNamed is a wrapper around gtk_stack_add_named().
func (v *Stack) AddNamed(child IWidget, name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_add_named(v.Native(), child.toWidget(), (*C.gchar)(cstr))
}

// AddTitled is a wrapper around gtk_stack_add_titled().
func (v *Stack) AddTitled(child IWidget, name, title string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_stack_add_titled(v.Native(), child.toWidget(), (*C.gchar)(cName),
		(*C.gchar)(cTitle))
}

// SetVisibleChild is a wrapper around gtk_stack_set_visible_child().
func (v *Stack) SetVisibleChild(child IWidget) {
	C.gtk_stack_set_visible_child(v.Native(), child.toWidget())
}

// GetVisibleChild is a wrapper around gtk_stack_get_visible_child().
func (v *Stack) GetVisibleChild() *Widget {
	c := C.gtk_stack_get_visible_child(v.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s
}

// SetVisibleChildName is a wrapper around gtk_stack_set_visible_child_name().
func (v *Stack) SetVisibleChildName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_name(v.Native(), (*C.gchar)(cstr))
}

// GetVisibleChildName is a wrapper around gtk_stack_get_visible_child_name().
func (v *Stack) GetVisibleChildName() string {
	c := C.gtk_stack_get_visible_child_name(v.Native())
	return C.GoString((*C.char)(c))
}

// SetVisibleChildFull is a wrapper around gtk_stack_set_visible_child_full().
func (v *Stack) SetVisibleChildFull(name string, transaction StackTransitionType) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_full(v.Native(), (*C.gchar)(cstr),
		C.GtkStackTransitionType(transaction))
}

// SetHomogeneous is a wrapper around gtk_stack_set_homogeneous().
func (v *Stack) SetHomogeneous(homogeneous bool) {
	C.gtk_stack_set_homogeneous(v.Native(), gbool(homogeneous))
}

// GetHomogeneous is a wrapper around gtk_stack_get_homogeneous().
func (v *Stack) GetHomogeneous() bool {
	c := C.gtk_stack_get_homogeneous(v.Native())
	return gobool(c)
}

// SetTransitionDuration is a wrapper around gtk_stack_set_transition_duration().
func (v *Stack) SetTransitionDuration(duration uint) {
	C.gtk_stack_set_transition_duration(v.Native(), C.guint(duration))
}

// GetTransitionDuration is a wrapper around gtk_stack_get_transition_duration().
func (v *Stack) GetTransitionDuration() uint {
	c := C.gtk_stack_get_transition_duration(v.Native())
	return uint(c)
}

// SetTransitionType is a wrapper around gtk_stack_set_transition_type().
func (v *Stack) SetTransitionType(transition StackTransitionType) {
	C.gtk_stack_set_transition_type(v.Native(), C.GtkStackTransitionType(transition))
}

// GetTransitionType is a wrapper around gtk_stack_get_transition_type().
func (v *Stack) GetTransitionType() StackTransitionType {
	c := C.gtk_stack_get_transition_type(v.Native())
	return StackTransitionType(c)
}

/*
 * GtkStackSwitcher
 */

// StackSwitcher is a representation of GTK's GtkStackSwitcher
type StackSwitcher struct {
	Box
}

// Native returns a pointer to the underlying GtkStackSwitcher.
func (v *StackSwitcher) Native() *C.GtkStackSwitcher {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStackSwitcher(p)
}

func wrapStackSwitcher(obj *glib.Object) *StackSwitcher {
	return &StackSwitcher{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// StackSwitcherNew is a wrapper around gtk_stack_switcher_new().
func StackSwitcherNew() (*StackSwitcher, error) {
	c := C.gtk_stack_switcher_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapStackSwitcher(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// SetStack is a wrapper around gtk_stack_switcher_set_stack().
func (v *StackSwitcher) SetStack(stack *Stack) {
	C.gtk_stack_switcher_set_stack(v.Native(), stack.Native())
}

// GetStack is a wrapper around gtk_stack_switcher_get_stack().
func (v *StackSwitcher) GetStack() *Stack {
	c := C.gtk_stack_switcher_get_stack(v.Native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapStack(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s
}

/*
 * GtkWindow
 */

// SetTitlebar is a wrapper around gtk_window_set_titlebar().
func (v *Window) SetTitlebar(titlebar IWidget) {
	C.gtk_window_set_titlebar(v.Native(), titlebar.toWidget())
}

// Close is a wrapper around gtk_window_close().
func (v *Window) Close() {
	C.gtk_window_close(v.Native())
}

func cast_3_10(class string, o *glib.Object) glib.IObject {
	var g glib.IObject
	switch class {
	case "GtkListBox":
		g = wrapListBox(o)
	case "GtkListBoxRow":
		g = wrapListBoxRow(o)
	case "GtkRevealer":
		g = wrapRevealer(o)
	case "GtkSearchBar":
		g = wrapSearchBar(o)
	case "GtkStack":
		g = wrapStack(o)
	}
	return g
}

func init() {
	cast_3_10_func = cast_3_10
}
