/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package github.com/khulnasoft/proxy/proxylib */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h>

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */


#line 9 "proxylib.go"

#include "proxylib/types.h"

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef size_t GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
#ifdef _MSC_VER
#include <complex.h>
typedef _Fcomplex GoComplex64;
typedef _Dcomplex GoComplex128;
#else
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;
#endif

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


// OnNewConnection is used to register a new connection of protocol 'proto'.
// Note that the 'origBuf' and replyBuf' type '*[]byte' corresponds to 'InjectBuf' type, but due to
// cgo export restrictions we can't use the go type in the prototype.
//
extern FilterResult OnNewConnection(GoUint64 instanceId, GoString proto, GoUint64 connectionId, GoUint8 ingress, GoUint32 srcId, GoUint32 dstId, GoString srcAddr, GoString dstAddr, GoString policyName, GoSlice* origBuf, GoSlice* replyBuf);

// Each connection is assumed to be called from a single thread, so accessing connection metadata
// does not need protection.
//
// OnData gets all the unparsed data the datapath has received so far. The data is provided to the parser
// associated with the connection, and the parser is expected to find if the data frame contains enough data
// to make a PASS/DROP decision for the whole data frame. Note that the whole data frame need not be received,
// if the decision including the length of the data frame in bytes can be determined based on the beginning of
// the data frame only (e.g., headers including the length of the data frame). The parser returns a decision
// with the number of bytes on which the decision applies. If more data is available, then the parser will be
// called again with the remaining data. Parser needs to return MORE if a decision can't be made with
// the available data, including the minimum number of additional bytes that is needed before the parser is
// called again.
//
// The parser can also inject at arbitrary points in the data stream. This is indecated by an INJECT operation
// with the number of bytes to be injected. The actual bytes to be injected are provided via an Inject()
// callback prior to returning the INJECT operation. The Inject() callback operates on a limited size buffer
// provided by the datapath, and multiple INJECT operations may be needed to inject large amounts of data.
// Since we get the data on one direction at a time, any frames to be injected in the reverse direction
// are placed in the reverse direction buffer, from where the datapath injects the data before calling
// us again for the reverse direction input.
//
extern FilterResult OnData(GoUint64 connectionId, GoUint8 reply, GoUint8 endStream, GoSlice* data, GoSlice* filterOps);

// Make this more general connection event callback
//
extern void Close(GoUint64 connectionId);

// OpenModule is called before any other APIs.
// Called concurrently by different filter instances.
// Returns a library instance ID that must be passed to all other API calls.
// Calls with the same parameters will return the same instance.
// Zero return value indicates an error.
//
extern GoUint64 OpenModule(GoSlice params, GoUint8 debug);
extern void CloseModule(GoUint64 id);

#ifdef __cplusplus
}
#endif
