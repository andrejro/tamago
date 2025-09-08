TamaGo - bare metal Go - THSM EP1 support
=========================================

tamago | https://github.com/usbarmory/tamago  

Copyright (c) The TamaGo Authors. All Rights Reserved.  

![TamaGo gopher](https://github.com/usbarmory/tamago/wiki/images/tamago.svg?sanitize=true)

Authors
=======

Andrea Barisani  
andrea@inversepath.com  

Andrej Rosano  
andrej@inversepath.com  

Introduction
============

TamaGo is a framework that enables compilation and execution of unencumbered Go
applications on bare metal AMD64/ARM/RISC-V processors.

The [thsm_ep1](https://github.com/usbarmory/tamago/tree/master/board/tillitis/thsm_ep1)
package provides support for the THSM EP1 development board.

Documentation
=============

[![Go Reference](https://pkg.go.dev/badge/github.com/usbarmory/tamago.svg)](https://pkg.go.dev/github.com/usbarmory/tamago)

For more information about TamaGo see its
[repository](https://github.com/usbarmory/tamago) and
[project wiki](https://github.com/usbarmory/tamago/wiki).

For the underlying driver support for this board see package
[imx6ul](https://github.com/usbarmory/tamago/tree/master/soc/nxp/imx6ul).

The package API documentation can be found on
[pkg.go.dev](https://pkg.go.dev/github.com/usbarmory/tamago).

Supported hardware
==================

| SoC          | Board    | SoC package                                                              | Board package                                                                       |
|--------------|----------|--------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| NXP i.MX6ULL | THSM EP1 | [imx6ul](https://github.com/usbarmory/tamago/tree/master/soc/nxp/imx6ul) | [thsm_ep1](https://github.com/usbarmory/tamago/tree/master/board/tillitis/thsm_ep1) |

Compiling
=========

Go applications are simply required to import, the relevant board package to
ensure that hardware initialization and runtime support take place:

```golang
import (
	_ "github.com/usbarmory/tamago/board/tillitis/thsm_ep1"
)
```

Build the [TamaGo compiler](https://github.com/usbarmory/tamago-go)
(or use the [latest binary release](https://github.com/usbarmory/tamago-go/releases/latest)):

```
wget https://github.com/usbarmory/tamago-go/archive/refs/tags/latest.zip
unzip latest.zip
cd tamago-go-latest/src && ./all.bash
cd ../bin && export TAMAGO=`pwd`/go
```

Go applications can be compiled as usual, using the compiler built in the
previous step, but with the addition of the following flags/variables:

```
GOOS=tamago GOARM=7 GOARCH=arm ${TAMAGO} build -ldflags "-T 0x80010000 -R 0x1000" main.go
```

Build tags
==========

The following build tags allow application to override the package own definition of
[external functions required by the runtime](https://github.com/usbarmory/tamago/wiki/Internals#go-runtime-changes):

* `linkramsize`: exclude `ramSize` from `mem.go`
* `linkprintk`: exclude `printk` from `console.go`
