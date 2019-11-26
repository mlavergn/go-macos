package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

void cTestAlloc(int argc, void **args) {
  void **argv = (void**)args;
  NSLog(@"argc: %d", argc);
  for (int i=0; i<argc; i++) {
    if (i == 2) {
      NSLog(@"%i strp: %c", i, (char *)argv[i]);
    } else {
      NSLog(@"%i numpp: %d", i, (int **)argv[i]);
      NSLog(@"%i nump: %d", i, (int *)argv[i]);
      NSLog(@"%i num: %d", i, (int)argv[i]);
    }
  }
}

void cTestUnsafeArray(int argc, void *args) {
  void **argv = (void**)args;
  NSLog(@"argc: %d", argc);
  for (int i=0; i<argc; i++) {
    if (i == 2) {
      NSLog(@"%i strp: %c", i, (char *)argv[i]);
    } else {
      NSLog(@"%i numpp: %d", i, (int **)argv[i]);
      NSLog(@"%i nump: %d", i, (int *)argv[i]);
      NSLog(@"%i num: %d", i, (int)argv[i]);
    }
  }
}
*/
import "C"

import (
  "unsafe"
  "fmt"
)

func testAlloc() {
	fmt.Println("\testAlloc\n")
  
	arg1 := C.long(0)
	arg2 := C.long(7)
	arg3 := C.CString("Hello")
	defer C.free(unsafe.Pointer(arg3))
	arg4 := C.long(22)
	arg5 := C.long(3)
  
	argc := 5
	// argv := unsafe.Pointer(C.calloc(C.ulong(argc), unsafe.Sizeof(*numsPtr)))
	step := C.ulong(unsafe.Sizeof(unsafe.Pointer(&argc))) // 8 bytes (64 bit)
	fmt.Println(step)
	argv := C.calloc(C.ulong(argc), step)
  
	argp := C.ulong(uintptr(argv))
	fmt.Println(C.ulong(argp))
	*(*unsafe.Pointer)(unsafe.Pointer(uintptr(argp))) = unsafe.Pointer(&arg1)
	
	argp = argp + step
	fmt.Println(C.ulong(argp))
	*(*unsafe.Pointer)(unsafe.Pointer(uintptr(argp))) = unsafe.Pointer(&arg2)
	
	argp = argp + step
	fmt.Println(C.ulong(argp))
	*(*unsafe.Pointer)(unsafe.Pointer(uintptr(argp))) = unsafe.Pointer(&arg3)
  
	argp = argp + step
	fmt.Println(C.ulong(argp))
	*(*unsafe.Pointer)(unsafe.Pointer(uintptr(argp))) = unsafe.Pointer(&arg4)
  
	argp = argp + step
	fmt.Println(C.ulong(argp))
	*(*unsafe.Pointer)(unsafe.Pointer(uintptr(argp))) = unsafe.Pointer(&arg5)
  
	C.cTestAlloc(C.int(argc), &argv)
  }
  
  func testUnsafeArray() {
	fmt.Println("\ntestUnsafeArray\n")
  
	arg1 := C.long(0)
	arg2 := C.long(7)
  
	// the string breaks from contiguous memory, using an int here passes
	// arg3 := C.long(8)
	arg3 := C.CString("hello")
	defer C.free(unsafe.Pointer(arg3))
  
	arg4 := C.long(22)
	arg5 := C.long(3)

	argv := [5]unsafe.Pointer { unsafe.Pointer(&arg1), unsafe.Pointer(&arg2), unsafe.Pointer(&arg3), unsafe.Pointer(&arg4), unsafe.Pointer(&arg5) }
  
	// works fine for primitive values, breaks on pointer types
	C.cTestUnsafeArray(C.int(len(argv)), argv[0])
  }
  
  func main() {
	// testAlloc()
	testUnsafeArray()
  }