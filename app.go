package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

void cInvoke(void *cobject, char *cmethod, int argc, void *args) {
  id object = cobject;
  NSString *method = [NSString stringWithUTF8String: cmethod];

  SEL sel = NSSelectorFromString(method);
  NSInvocation *inv = [NSInvocation invocationWithMethodSignature:[object methodSignatureForSelector:sel]];
  [inv setSelector:sel];
  [inv setTarget:object];
  // arguments 0 and 1 are self and _cmd respectively
  // automatically set by NSInvocation
  for (int i=0; i<argc; i++) {
    [inv setArgument:&args[i] atIndex:i+2];
  }
  [inv invoke];
}

void* cPerformSelector(void *cobject, char *cmethod) {
  id object = cobject;
  NSString *method = [NSString stringWithUTF8String: cmethod];

  return [object performSelector:NSSelectorFromString(method)];
}

void* cNSAutoreleasePool(void) {
  return [NSAutoreleasePool new];
}

void* cNSApplication(void) {
  return [NSApplication sharedApplication];
}

void Demo(void) {
  // [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
  // menu items
  id menubar = [[NSMenu new] autorelease];
  id appMenuItem = [[NSMenuItem new] autorelease];
  [menubar addItem:appMenuItem];
  [NSApp setMainMenu:menubar];
  id appMenu = [[NSMenu new] autorelease];
  id appName = [[NSProcessInfo processInfo] processName];
  id quitTitle = [@"Quit " stringByAppendingString:appName];
  id quitMenuItem = [[[NSMenuItem alloc] initWithTitle:quitTitle
    action:@selector(terminate:) keyEquivalent:@"q"]
      autorelease];
  [appMenu addItem:quitMenuItem];
  [appMenuItem setSubmenu:appMenu];
  // window
  id window = [[[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, 200, 200)
    styleMask:NSTitledWindowMask backing:NSBackingStoreBuffered defer:NO]
      autorelease];
  [window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
  [window setTitle:appName];
  [window makeKeyAndOrderFront:nil];
}
*/
import "C"

import (
  "unsafe"
)

// ObjC data type helpers

type nsObject unsafe.Pointer
type nsInteger C.long
type nsEnum nsInteger
func nsBool(val bool) C.schar { 
  if val { 
    return C.schar(1) 
  }
  return C.schar(0) 
}
func nsString(val string) *C.char {
  return C.CString(val)
}

// Go data type helpers

func goArrayCString(val unsafe.Pointer, len int) []*C.char {
  return (*[1 << 30]*C.char)(unsafe.Pointer(val))[:len:len]
}

// NSApp text
type NSApp struct {
  instance unsafe.Pointer
}

// NewNSApp text
func NewNSApp() NSApp {
  nsapp := NSApp{}
  nsapp.instance = C.cNSApplication()
  return nsapp
}

// NSApplicationActivationPolicy text
const (
  NSApplicationActivationPolicyRegular int = iota
  NSApplicationActivationPolicyAccessory int = iota
  NSApplicationActivationPolicyProhibited int = iota
)

func (nsapp *NSApp) setActivationPolicy(val int) {
  arg1 := nsEnum(val)
  argv := []unsafe.Pointer { unsafe.Pointer(&arg1) }
  C.cInvoke(nsapp.instance, C.CString("setActivationPolicy:"), C.int(len(argv)), argv[0])
}

func (nsapp *NSApp) activateIgnoringOtherApps(val bool) {
  arg1 := nsBool(val)
  argv := []unsafe.Pointer { unsafe.Pointer(&arg1) }
  C.cInvoke(nsapp.instance, C.CString("activateIgnoringOtherApps:"), C.int(len(argv)), argv[0])
}

func (nsapp *NSApp) run() {
  C.cPerformSelector(nsapp.instance, C.CString("run"))
}

// NSAutoreleasePool text
func NSAutoreleasePool() unsafe.Pointer {
	return C.cNSAutoreleasePool()
}

// demo
func main() {
  NSAutoreleasePool()
  nsapp := NewNSApp()
  nsapp.setActivationPolicy(NSApplicationActivationPolicyRegular)
  C.Demo()
  // activate
  nsapp.activateIgnoringOtherApps(true)
  nsapp.run()
}
