package config

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/alexsasharegan/dotenv"
	"github.com/fairyhunter13/reflecthelper/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// List of all errors
var (
	ErrInputMustBePointer = errors.New("the input interface must be a pointer")
)

var (
	typeString          = reflect.TypeOf("hello")
	typeMapStringString = reflect.TypeOf(map[string]string{})
	typeMapStringBool   = reflect.TypeOf(map[string]bool{})
	typeMap             = reflect.TypeOf(map[string]interface{}{})
)

func isUnsupportedTypeMap(to reflect.Type) bool {
	return to != typeMapStringString && to != typeMapStringBool && to != typeMap
}

// StringToVariousMapsHookFunc generate decode hook func to decode JSON string to map[string]string.
func StringToVariousMapsHookFunc(mapType reflect.Type) mapstructure.DecodeHookFunc {
	return func(from reflect.Type, to reflect.Type, data interface{}) (res interface{}, err error) {
		res = data
		if from != typeString || isUnsupportedTypeMap(to) {
			return
		}

		var mapRes interface{}
		switch mapType {
		case typeMapStringBool:
			mapRes = map[string]bool{}
		case typeMapStringString:
			mapRes = map[string]string{}
		default:
			mapRes = map[string]interface{}{}
		}
		dataStr, _ := data.(string)
		err = json.Unmarshal([]byte(dataStr), &mapRes)
		if err != nil {
			return
		}

		res = mapRes
		return
	}
}

// LoadConfig decode the env config to the iface using the specified viperConf.
func LoadConfig(iface interface{}, paths ...string) (err error) {
	if reflecthelper.GetKind(reflect.ValueOf(iface)) != reflect.Ptr {
		err = ErrInputMustBePointer
		return
	}

	err = dotenv.Load(paths...)
	if err != nil {
		return
	}

	viperConf := viper.GetViper()
	viperConf.SetConfigType("env")
	viperConf.AutomaticEnv()
	viperConf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	bindEnvs(viperConf, iface)
	err = viperConf.Unmarshal(iface, func(dc *mapstructure.DecoderConfig) {
		dc.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			StringToVariousMapsHookFunc(typeMap),
			StringToVariousMapsHookFunc(typeMapStringBool),
			StringToVariousMapsHookFunc(typeMapStringString),
			mapstructure.StringToIPHookFunc(),
			mapstructure.StringToIPNetHookFunc(),
			dc.DecodeHook,
			mapstructure.StringToTimeHookFunc(time.RFC3339),
		)
	})
	return
}

func bindEnvs(viperConf *viper.Viper, iface interface{}, parts ...string) {
	val := reflecthelper.GetChildElem(reflect.ValueOf(iface))
	if reflecthelper.GetKind(val) != reflect.Struct {
		return
	}

	if viperConf == nil {
		viperConf = viper.GetViper()
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		v := val.Field(i)
		t := typ.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(viperConf, v.Interface(), append(parts, tv)...)
		default:
			viperConf.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
