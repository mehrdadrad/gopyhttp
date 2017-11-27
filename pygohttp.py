from ctypes import *

lib = cdll.LoadLibrary("./gopyhttp.so")

class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

class GoSlice(Structure):
    _fields_ = [("data", POINTER(c_void_p)),
                ("len", c_longlong), ("cap", c_longlong)]

url = GoString(b"https://www.google.com", 22)

lib.HTTPRequest.argtypes = [GoString]
lib.HTTPRequest.restype = None
lib.HTTPRequest(url)
