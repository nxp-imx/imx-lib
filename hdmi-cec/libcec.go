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

package libcec

import (
        "android/soong/android"
        "android/soong/cc"
        "github.com/google/blueprint/proptools"
)

func init() {
    android.RegisterModuleType("libcec_defaults", libcecDefaultsFactory)
}

func libcecDefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, libcecDefaults)
    return module
}

func libcecDefaults(ctx android.LoadHookContext) {
    type props struct {
        Cflags []string
        Target struct {
                Android struct {
                        Enabled *bool
                }
        }
    }

    p := &props{}
    if ctx.Config().VendorConfig("IMXPLUGIN").Bool("BOARD_HAVE_HDMI") {
        p.Target.Android.Enabled = proptools.BoolPtr(true)
    } else {
        p.Target.Android.Enabled = proptools.BoolPtr(false)
    }
    if ctx.Config().VendorConfig("IMXPLUGIN").String("BOARD_SOC_CLASS") == "IMX6" {
        p.Cflags = append(p.Cflags, "-DBUILD_FOR_ANDROID")
        p.Cflags = append(p.Cflags, "-DIMX6Q")
    } else {
        p.Cflags = append(p.Cflags, "-DBUILD_FOR_ANDROID")
        p.Cflags = append(p.Cflags, "-D"+ ctx.Config().VendorConfig("IMXPLUGIN").String("BOARD_SOC_TYPE"))
    }

    ctx.AppendProperties(p)
}
