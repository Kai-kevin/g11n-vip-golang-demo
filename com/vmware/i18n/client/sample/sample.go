package sample

import "vipgoclient/com/vmware/i18n/client/util"

//TODO sample for call the format utils

type Translator struct {
	Locale string
	Component string
}


//Get the translation from VIP service without comment to the source, if there's no available translation will return the source as the default of value.
//
//@param tranlator contains locale and component information.
//@param args	the array of parameters for the place holder.
//e.g, if the source is "This is {0}, {1}", the first element of args represents {0}, the second represents [1].
//@return	a translation string.

func (translator *Translator)GetTranslation(key string,params ...interface{}) (string,error) {
	return util.GetTranslationCacheManagerInstance().
		GetTranslationByKeyWithParams(key,translator.Locale,translator.Component,params)
}
