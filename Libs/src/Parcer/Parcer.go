//   Copyright (c) 2016 Ivan A Kostko (github.com/ivan-kostko)

//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

//       http://www.apache.org/licenses/LICENSE-2.0

//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package Parcer

//go:generate stringer -type=SupportedCodec

// Represents ENUM of supported codecs
type SupportedCodec int

const (
	DefaultXML SupportedCodec = iota
	DefaultJSON
	DefaultYAML
)

// Represents the list of registered serializers
var registeredSerializers map[SupportedCodec]Serializer

// Represents the list of registered deserializers
var registeredDeserializers map[SupportedCodec]Deserializer

// Defines functionality of parser as combination of two functions: Serialize + Deserialize
type Parser struct {
	Serializer
	Deserializer
}
