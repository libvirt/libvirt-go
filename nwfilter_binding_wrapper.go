/*
 * This file is part of the libvirt-go project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (C) 2018 Red Hat, Inc.
 *
 */

package libvirt

/*
#cgo pkg-config: libvirt
#include <assert.h>
#include "nwfilter_binding_wrapper.h"

const char *virNWFilterBindingGetPortDevWrapper(virNWFilterBindingPtr binding)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNWFilterBindingGetPortDev(binding);
#endif
}


const char *virNWFilterBindingGetFilterNameWrapper(virNWFilterBindingPtr binding)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNWFilterBindingGetFilterName(binding);
#endif
}


char *virNWFilterBindingGetXMLDescWrapper(virNWFilterBindingPtr binding,
					 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNWFilterBindingGetXMLDesc(binding, flags);
#endif
}


int virNWFilterBindingDeleteWrapper(virNWFilterBindingPtr binding)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNWFilterBindingDelete(binding);
#endif
}


int virNWFilterBindingRefWrapper(virNWFilterBindingPtr binding)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNWFilterBindingRef(binding);
#endif
}


int virNWFilterBindingFreeWrapper(virNWFilterBindingPtr binding)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNWFilterBindingFree(binding);
#endif
}


*/
import "C"
