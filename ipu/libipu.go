// Copyright 2019 NXP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libipu

import (
        "android/soong/android"
        "android/soong/cc"
        "github.com/google/blueprint/proptools"
)

func init() {
    android.RegisterModuleType("libipu_defaults", libipuDefaultsFactory)
}

func libipuDefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, libipuDefaults)
    return module
}

func libipuDefaults(ctx android.LoadHookContext) {
    type props struct {
        srcs []string
        Target struct {
                Android struct {
                        Enabled *bool
                }
        }
    }

    p := &props{}
    if ctx.Config().VendorConfig("IMXPLUGIN").Bool("HAVE_FSL_IMX_IPU") {
        p.Target.Android.Enabled = proptools.BoolPtr(true)
    } else {
        p.Target.Android.Enabled = proptools.BoolPtr(false)
    }
    if ctx.Config().VendorConfig("IMXPLUGIN").String("BOARD_SOC_CLASS") == "IMX5X" {
        p.srcs = append(p.srcs, "mxc_ipu_hl_lib.c")
        p.srcs = append(p.srcs, "mxc_ipu_lib.c")
    } else {
        p.srcs = append(p.srcs, "mxc_ipu_hl_lib_dummy.c")
    }

    ctx.AppendProperties(p)
}

