/***
 * Motto Module
 *
 * This file add motto module to let user can use specify Plurk API
 */

package robot

import (
	"github.com/ddliu/motto"
	"github.com/elct9620/go-plurk-robot/logger"
	"github.com/robertkrimen/otto"
)

func plurkModuleLoader(vm *motto.Motto) (otto.Value, error) {
	module, _ := vm.Object(`({})`)

	// Set module functions
	module.Set("addPlurk", plurk_AddPlurk)

	return vm.ToValue(module)
}

// Add Plurk API
func plurk_AddPlurk(call otto.FunctionCall) otto.Value {
	message, _ := call.Argument(0).ToString()
	qualifier, _ := call.Argument(1).ToString()
	lang, _ := call.Argument(2).ToString()

	// No message specify, return error
	if len(message) <= 0 {
		return otto.FalseValue()
	}

	// Default qualifier
	if qualifier == "undefined" {
		qualifier = ":"
	}

	if lang == "undefined" {
		lang = "en"
	}

	timeline := client.GetTimeline()
	res, err := timeline.PlurkAdd(message, qualifier, make([]int, 0), false, lang, true)

	if err != nil {
		logger.Error("Add Plurk failed, because %s", err.Error())
		return otto.FalseValue()
	}

	logger.Info("New plurk added, content is %s", res.RawContent)

	return otto.TrueValue()
}
