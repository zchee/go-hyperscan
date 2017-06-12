// Copyright 2017 The go-hyperscan Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// WARNING: This file has automatically been generated on Mon, 12 Jun 2017 17:47:08 JST.
// By https://git.io/c-for-go. DO NOT EDIT.

#include "hs/hs.h"
#include "hs/hs_common.h"
#include "hs/hs_compile.h"
#include "hs/hs_runtime.h"
#include <stdlib.h>
#pragma once

#define __CGOGEN 1

// hs_alloc_t_29384626 is a proxy for callback hs_alloc_t.
void* hs_alloc_t_29384626(unsigned long int size);

// hs_free_t_83e09d93 is a proxy for callback hs_free_t.
void hs_free_t_83e09d93(void* ptr);

// match_event_handler_6c9f61fb is a proxy for callback match_event_handler.
int match_event_handler_6c9f61fb(unsigned int id, unsigned long long from, unsigned long long to, unsigned int flags, void* context);

