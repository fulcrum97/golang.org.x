// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package windriver

import "unsafe"
import "syscall"
import "golang.org/x/exp/shiny/driver/internal/win32"

var _ unsafe.Pointer

var (
	modgdi32   = syscall.NewLazyDLL("gdi32.dll")
	modmsimg32 = syscall.NewLazyDLL("msimg32.dll")
	moduser32  = syscall.NewLazyDLL("user32.dll")

	procCreateDIBSection   = modgdi32.NewProc("CreateDIBSection")
	procCreateCompatibleDC = modgdi32.NewProc("CreateCompatibleDC")
	procDeleteDC           = modgdi32.NewProc("DeleteDC")
	procSelectObject       = modgdi32.NewProc("SelectObject")
	procAlphaBlend         = modmsimg32.NewProc("AlphaBlend")
	procBitBlt             = modgdi32.NewProc("BitBlt")
	procCreateSolidBrush   = modgdi32.NewProc("CreateSolidBrush")
	procFillRect           = moduser32.NewProc("FillRect")
	procDeleteObject       = modgdi32.NewProc("DeleteObject")
)

func _CreateDIBSection(dc syscall.Handle, bmi *_BITMAPINFO, usage uint32, bits **byte, section syscall.Handle, offset uint32) (bitmap syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateDIBSection.Addr(), 6, uintptr(dc), uintptr(unsafe.Pointer(bmi)), uintptr(usage), uintptr(unsafe.Pointer(bits)), uintptr(section), uintptr(offset))
	bitmap = syscall.Handle(r0)
	if bitmap == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _CreateCompatibleDC(dc win32.HDC) (newdc win32.HDC, err error) {
	r0, _, e1 := syscall.Syscall(procCreateCompatibleDC.Addr(), 1, uintptr(dc), 0, 0)
	newdc = win32.HDC(r0)
	if newdc == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _DeleteDC(dc win32.HDC) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteDC.Addr(), 1, uintptr(dc), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _SelectObject(dc win32.HDC, gdiobj syscall.Handle) (newobj syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procSelectObject.Addr(), 2, uintptr(dc), uintptr(gdiobj), 0)
	newobj = syscall.Handle(r0)
	if newobj == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _AlphaBlend(dcdest win32.HDC, xoriginDest int32, yoriginDest int32, wDest int32, hDest int32, dcsrc win32.HDC, xoriginSrc int32, yoriginSrc int32, wsrc int32, hsrc int32, ftn uintptr) (err error) {
	r1, _, e1 := syscall.Syscall12(procAlphaBlend.Addr(), 11, uintptr(dcdest), uintptr(xoriginDest), uintptr(yoriginDest), uintptr(wDest), uintptr(hDest), uintptr(dcsrc), uintptr(xoriginSrc), uintptr(yoriginSrc), uintptr(wsrc), uintptr(hsrc), uintptr(ftn), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _BitBlt(dcdest win32.HDC, xdest int32, ydest int32, width int32, height int32, dcsrc win32.HDC, xsrc int32, ysrc int32, rop int32) (ok int32, err error) {
	r0, _, e1 := syscall.Syscall9(procBitBlt.Addr(), 9, uintptr(dcdest), uintptr(xdest), uintptr(ydest), uintptr(width), uintptr(height), uintptr(dcsrc), uintptr(xsrc), uintptr(ysrc), uintptr(rop))
	ok = int32(r0)
	if ok == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _CreateSolidBrush(color _COLORREF) (brush syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateSolidBrush.Addr(), 1, uintptr(color), 0, 0)
	brush = syscall.Handle(r0)
	if brush == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _FillRect(dc win32.HDC, rc *_RECT, brush syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFillRect.Addr(), 3, uintptr(dc), uintptr(unsafe.Pointer(rc)), uintptr(brush))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _DeleteObject(object syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteObject.Addr(), 1, uintptr(object), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}