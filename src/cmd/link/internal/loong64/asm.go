// Inferno utils/5l/asm.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/5l/asm.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//	Portions Copyright © 2021 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package loong64

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/loader"
	"cmd/link/internal/sym"
	"debug/elf"
	"fmt"
	"log"
	"sync"
)

func gentext2(ctxt *ld.Link, ldr *loader.Loader) {}

func adddynrel(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s *sym.Symbol, r *sym.Reloc) bool {
	log.Fatalf("adddynrel not implemented")
	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	// loong64 ELF relocation (endian neutral)
	//		offset	uint64
	//		sym		  uint64
	//		addend	int64

	elfsym := ld.ElfSymForReloc(ctxt, r.Xsym)
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR, objabi.R_DWARFSECREF:
		switch r.Siz {
		case 4:
			ctxt.Out.Write64(uint64(sectoff))
			ctxt.Out.Write64(uint64(elf.R_LARCH_32) | uint64(elfsym)<<32)
			ctxt.Out.Write64(uint64(r.Xadd))
		case 8:
			ctxt.Out.Write64(uint64(sectoff))
			ctxt.Out.Write64(uint64(elf.R_LARCH_64) | uint64(elfsym)<<32)
			ctxt.Out.Write64(uint64(r.Xadd))
		default:
			return false
		}
		//TODO: XXX: tls not implement now
	case objabi.R_ADDRLOONG64TLS:
		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_TLS_TPREL) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_ABSOLUTE))
		ctxt.Out.Write64(uint64(0xfff))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_AND))
		ctxt.Out.Write64(uint64(0x0))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_POP_32_U_10_12))
		ctxt.Out.Write64(uint64(0x0))

	case objabi.R_ADDRLOONG64TLSU:
		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_TLS_TPREL) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_ABSOLUTE))
		ctxt.Out.Write64(uint64(0xc))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_SR))
		ctxt.Out.Write64(uint64(0x0))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_POP_32_S_5_20) | uint64(0)<<32)
		ctxt.Out.Write64(uint64(0x0))

	case objabi.R_CALLLOONG64:
		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_PLT_PCREL) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_POP_32_S_0_10_10_16_S2))
		ctxt.Out.Write64(uint64(0x0))

	case objabi.R_ADDRLOONG64:
		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_PCREL) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd + 0x4))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_PCREL) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd + 0x804))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_ABSOLUTE))
		ctxt.Out.Write64(uint64(0xc))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_SR))
		ctxt.Out.Write64(uint64(0x0))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_ABSOLUTE))
		ctxt.Out.Write64(uint64(0xc))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_SL))
		ctxt.Out.Write64(uint64(0x0))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_SUB))
		ctxt.Out.Write64(uint64(0x0))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_POP_32_S_10_12))
		ctxt.Out.Write64(uint64(0x0))

	case objabi.R_ADDRLOONG64U:
		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_PCREL) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd + 0x800))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_PUSH_ABSOLUTE))
		ctxt.Out.Write64(uint64(0xc))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_SR))
		ctxt.Out.Write64(uint64(0x0))

		ctxt.Out.Write64(uint64(sectoff))
		ctxt.Out.Write64(uint64(elf.R_LARCH_SOP_POP_32_S_5_20) | uint64(0)<<32)
		ctxt.Out.Write64(uint64(0x0))
	}

	return true
}
func elfsetupplt(ctxt *ld.Link, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
	return
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	return false
}

func archreloc(target *ld.Target, syms *ld.ArchSyms, r *sym.Reloc, s *sym.Symbol, val int64) (int64, bool) {
	if target.IsExternal() {
		switch r.Type {
		default:
			return val, false
		case objabi.R_ADDRLOONG64,
			objabi.R_ADDRLOONG64U:
			r.Done = false

			// set up addend for eventual relocation via outer symbol.
			rs := r.Sym
			r.Xadd = r.Add
			for rs.Outer != nil {
				r.Xadd += ld.Symaddr(rs) - ld.Symaddr(rs.Outer)
				rs = rs.Outer
			}

			if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Sect == nil {
				ld.Errorf(s, "missing section for %s", rs.Name)
			}
			r.Xsym = rs

			return val, true
		case objabi.R_ADDRLOONG64TLS,
			objabi.R_ADDRLOONG64TLSU,
			objabi.R_CALLLOONG64,
			objabi.R_JMPLOONG64:
			r.Done = false
			r.Xsym = r.Sym
			r.Xadd = r.Add
			return val, true
		}
	}

	switch r.Type {
	case objabi.R_CONST:
		return r.Add, true
	case objabi.R_GOTOFF:
		return ld.Symaddr(r.Sym) + r.Add - ld.Symaddr(syms.GOT), true
	case objabi.R_ADDRLOONG64,
		objabi.R_ADDRLOONG64U:
		pc := s.Value + int64(r.Off)
		t := ld.Symaddr(r.Sym) + r.Add - pc
		o1 := target.Arch.ByteOrder.Uint32(s.P[r.Off:])
		if r.Type == objabi.R_ADDRLOONG64 {
			return int64(o1&0xffc003ff | (uint32((t+4-((t+4+1<<11)>>12<<12))<<10) & 0x3ffc00)), true
		}
		return int64(o1&0xfe00001f | (uint32((t+1<<11)>>12<<5) & 0x1ffffe0)), true
	case objabi.R_ADDRLOONG64TLS,
		objabi.R_ADDRLOONG64TLSU:
		// thread pointer is at 0x7000 offset from the start of TLS data area
		t := ld.Symaddr(r.Sym) + r.Add
		o1 := target.Arch.ByteOrder.Uint32(s.P[r.Off:])
		if r.Type == objabi.R_ADDRLOONG64TLS {
			return int64(o1&0xffc003ff | (uint32(t&0xfff) << 10)), true
		}
		return int64(o1&0xfe00001f | (uint32((t)>>12<<5) & 0x1ffffe0)), true
	case objabi.R_CALLLOONG64,
		objabi.R_JMPLOONG64:
		pc := s.Value + int64(r.Off)
		// Low 26 bits = (S + A - PC) >> 2
		t := ld.Symaddr(r.Sym) + r.Add - pc
		o1 := target.Arch.ByteOrder.Uint32(s.P[r.Off:])
		return int64(o1&0xfc000000 | (uint32((t>>2)&^0xffff0000) << 10) | (uint32((t>>2)&^0xfc00ffff) >> 16)), true
	}

	return val, false
}

func archrelocvariant(target *ld.Target, syms *ld.ArchSyms, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	return -1
}

func asmb(ctxt *ld.Link, _ *loader.Loader) {
	if ctxt.IsELF {
		ld.Asmbelfsetup()
	}

	var wg sync.WaitGroup
	sect := ld.Segtext.Sections[0]
	offset := sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff
	ld.WriteParallel(&wg, ld.Codeblk, ctxt, offset, sect.Vaddr, sect.Length)

	for _, sect := range ld.Segtext.Sections[1:] {
		offset := sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff
		ld.WriteParallel(&wg, ld.Datblk, ctxt, offset, sect.Vaddr, sect.Length)
	}

	if ld.Segrodata.Filelen > 0 {
		ld.WriteParallel(&wg, ld.Datblk, ctxt, ld.Segrodata.Fileoff, ld.Segrodata.Vaddr, ld.Segrodata.Filelen)
	}

	if ld.Segrelrodata.Filelen > 0 {
		ld.WriteParallel(&wg, ld.Datblk, ctxt, ld.Segrelrodata.Fileoff, ld.Segrelrodata.Vaddr, ld.Segrelrodata.Filelen)
	}

	ld.WriteParallel(&wg, ld.Datblk, ctxt, ld.Segdata.Fileoff, ld.Segdata.Vaddr, ld.Segdata.Filelen)

	ld.WriteParallel(&wg, ld.Dwarfblk, ctxt, ld.Segdwarf.Fileoff, ld.Segdwarf.Vaddr, ld.Segdwarf.Filelen)
	wg.Wait()
}

func asmb2(ctxt *ld.Link) {
	/* output symbol table */
	ld.Symsize = 0

	ld.Lcsize = 0
	symo := uint32(0)
	if !*ld.FlagS {
		// TODO: rationalize
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				symo = uint32(ld.Segdwarf.Fileoff + ld.Segdwarf.Filelen)
				symo = uint32(ld.Rnd(int64(symo), int64(*ld.FlagRound)))
			}

		case objabi.Hplan9:
			symo = uint32(ld.Segdata.Fileoff + ld.Segdata.Filelen)
		}

		ctxt.Out.SeekSet(int64(symo))
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				ld.Asmelfsym(ctxt)
				ctxt.Out.Write(ld.Elfstrdat)

				if ctxt.LinkMode == ld.LinkExternal {
					ld.Elfemitreloc(ctxt)
				}
			}

		case objabi.Hplan9:
			ld.Asmplan9sym(ctxt)

			sym := ctxt.Syms.Lookup("pclntab", 0)
			if sym != nil {
				ld.Lcsize = int32(len(sym.P))
				ctxt.Out.Write(sym.P)
			}
		}
	}

	ctxt.Out.SeekSet(0)
	switch ctxt.HeadType {
	default:
	case objabi.Hplan9: /* plan 9 */
		magic := uint32(4*18*18 + 7)
		if ctxt.Arch == sys.ArchLOONG64 {
			magic = uint32(4*26*26 + 7)
		}
		ctxt.Out.Write32(magic)                      /* magic */
		ctxt.Out.Write32(uint32(ld.Segtext.Filelen)) /* sizes */
		ctxt.Out.Write32(uint32(ld.Segdata.Filelen))
		ctxt.Out.Write32(uint32(ld.Segdata.Length - ld.Segdata.Filelen))
		ctxt.Out.Write32(uint32(ld.Symsize))          /* nsyms */
		ctxt.Out.Write32(uint32(ld.Entryvalue(ctxt))) /* va of entry */
		ctxt.Out.Write32(0)
		ctxt.Out.Write32(uint32(ld.Lcsize))

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd:
		ld.Asmbelf(ctxt, int64(symo))
	}

	if *ld.FlagC {
		fmt.Printf("textsize=%d\n", ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", ld.Segdata.Length-ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", ld.Symsize)
		fmt.Printf("lcsize=%d\n", ld.Lcsize)
		fmt.Printf("total=%d\n", ld.Segtext.Filelen+ld.Segdata.Length+uint64(ld.Symsize)+uint64(ld.Lcsize))
	}
}
